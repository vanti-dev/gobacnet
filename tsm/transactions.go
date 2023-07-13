/*Copyright (C) 2017 Alex Beltran

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to:
The Free Software Foundation, Inc.
59 Temple Place - Suite 330
Boston, MA  02111-1307, USA.

As a special exception, if other files instantiate templates or
use macros or inline functions from this file, or you compile
this file and link it with other works to produce a work based
on this file, this file does not by itself cause the resulting
work to be covered by the GNU General Public License. However
the source code for this file must still be made available in
accordance with section (3) of the GNU General Public License.

This exception does not invalidate any other reasons why a work
based on this file might be covered by the GNU General Public
License.
*/

package tsm

import (
	"context"
	"fmt"
	"sync"
)

// MaxTransaction is the default max number of transactions that can occur
// concurrently
const MaxTransaction = 255
const invalidID = 0

type state struct {
	data      chan interface{}
	reclaimed chan struct{} // closed when Put is called for the state id
	sends     int32         // must hold TSM.mutex when accessing this variable
}

// TSM is the transaction state manager. It handles passing data to other
// processes and keeping track of what transactions are currently processed
type TSM struct {
	mutex  sync.Mutex
	states map[int]*state
	pool   sync.Pool
	free   struct {
		id    chan int
		space chan struct{}
	}
}

// New creates a new transaction manager that can handle at most size concurrent transactions.
func New(size int) *TSM {
	t := &TSM{
		states: make(map[int]*state), pool: sync.Pool{
			// Operation doesn't include a new channel. We want that done when a get is
			// done since we close all channels when putting into the pool.
			New: func() interface{} {
				s := new(state)
				return s
			},
		},
	}

	// Generate free ids.
	t.free.id = make(chan int, MaxTransaction)
	for i := invalidID + 1; i < MaxTransaction; i++ {
		t.free.id <- i
	}

	// Generate free space
	t.free.space = make(chan struct{}, size)
	for i := 0; i < size; i++ {
		t.free.space <- struct{}{}
	}

	return t
}

// Send data to invoked id
func (t *TSM) Send(id int, b interface{}) error {
	t.mutex.Lock()
	s, ok := t.states[id]
	if ok {
		s.sends++
	}
	t.mutex.Unlock()

	if !ok {
		return fmt.Errorf("id %d is not receiving", id)
	}
	select {
	case s.data <- b:
		// There's a chance that between sending and entering the below lock someone Put the id.
		// In that case when Put sees s.sends it will be >0 so it won't close it, so we have to check.
		t.mutex.Lock()
		s.sends--
		if s.sends == 0 {
			select {
			case <-s.reclaimed:
				close(s.data)
			default:
			}
		}
		t.mutex.Unlock()
	case <-s.reclaimed:
		t.mutex.Lock()
		s.sends--
		if s.sends == 0 {
			close(s.data)
		}
		t.mutex.Unlock()
		return fmt.Errorf("id %d is not receiving", id)
	}
	return nil
}

// Receive attempts to receive a byte array from the invoked id
func (t *TSM) Receive(ctx context.Context, id int) (interface{}, error) {
	t.mutex.Lock()
	s, ok := t.states[id]
	t.mutex.Unlock()

	if !ok {
		return nil, fmt.Errorf("id %d is not sending", id)
	}

	// Wait for data
	select {
	case b, ok := <-s.data:
		if !ok {
			return nil, fmt.Errorf("id %d is not sending", id)
		}
		return b, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ID returns the invoke id that was used to save the state of this connection.
func (t *TSM) ID(ctx context.Context) (int, error) {
	var id int
	select {
	case <-t.free.space:
		// got a free spot, lets try and get a free id
		select {
		case id = <-t.free.id:
		case <-ctx.Done():
			t.free.space <- struct{}{}
			return 0, fmt.Errorf("unable to get a free id: %w", ctx.Err())
		}
	case <-ctx.Done():
		return 0, fmt.Errorf("no free space: %w", ctx.Err())
	}

	// skip error checking, since we control new generation and what is put in the pool.
	s := t.pool.Get().(*state)
	s.data = make(chan interface{})
	s.reclaimed = make(chan struct{})
	s.sends = 0 // just to be safe
	t.mutex.Lock()
	t.states[id] = s
	t.mutex.Unlock()
	return id, nil
}

// Put allows the id to be reused in the transaction manager.
// Put only returns an error if id is not known.
func (t *TSM) Put(id int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	s, ok := t.states[id]
	if !ok {
		return fmt.Errorf("id %d does not exist in the transactions", id)
	}

	close(s.reclaimed)
	if s.sends == 0 {
		close(s.data)
	} // else the send will close it for us
	t.pool.Put(s)
	t.free.id <- id
	t.free.space <- struct{}{}
	delete(t.states, id)
	return nil
}
