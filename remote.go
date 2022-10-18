package gobacnet

import (
	"github.com/vanti-dev/gobacnet/property"
	"github.com/vanti-dev/gobacnet/types"
)

// RemoteDevices is like WhoIs but does not use broadcast.
// Use RemoteDevices if you already know the address and IDs of the devices you intend to communicate with.
//
// The ids should each refer to an object of type types.Device.
// This will attempt to collect network comm settings like MaxApdu and act equivalently to the data you'd get in
// an IAm response.
func (c *Client) RemoteDevices(addr types.Address, ids ...types.ObjectInstance) ([]types.Device, error) {
	defaultDevice := types.Device{
		Addr:    addr,
		MaxApdu: 1000,
	}
	req := types.ReadMultipleProperty{}
	for _, id := range ids {
		oid := types.ObjectID{Type: types.DeviceType, Instance: id}
		req.Objects = append(req.Objects,
			types.Object{ID: oid, Properties: []types.Property{
				{Type: property.MaxApduLengthAccepted, ArrayIndex: ArrayAll},
				{Type: property.SegmentationSupported, ArrayIndex: ArrayAll},
				{Type: property.VendorIdentifier, ArrayIndex: ArrayAll},
			}},
		)
	}
	res, err := c.ReadProperties(defaultDevice, req)
	if err != nil {
		return nil, err
	}
	devices := make([]types.Device, len(res.Objects))
	for i, object := range res.Objects {
		device := types.Device{Addr: addr, ID: object.ID}
		maxApduProp, segProp, vendoProp := object.Properties[0], object.Properties[1], object.Properties[2]
		device.MaxApdu = maxApduProp.Data.(uint32)
		device.Segmentation = types.Enumerated(segProp.Data.(uint32))
		device.Vendor = vendoProp.Data.(uint32)
		devices[i] = device
	}
	return devices, nil
}
