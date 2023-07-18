package encoding

import (
	"github.com/vanti-dev/gobacnet/enum/pdutype"
	bactype "github.com/vanti-dev/gobacnet/types"
)

func (e *Encoder) ReadMultipleProperty(invokeID uint8, data bactype.ReadMultipleProperty) error {
	a := bactype.APDU{
		DataType:         pdutype.ConfirmedServiceRequest,
		Service:          bactype.ServiceConfirmedReadPropMultiple,
		MaxSegs:          0,
		MaxApdu:          MaxAPDU,
		InvokeId:         invokeID,
		SegmentedMessage: false,
	}
	e.APDU(a)
	err := e.objects(data.Objects)
	if err != nil {
		return err
	}

	return e.Error()
}

func (e *Encoder) objects(objects []bactype.Object) error {
	var tag uint8
	for _, obj := range objects {
		tag = 0
		e.contextObjectID(tag, obj.ID.Type, obj.ID.Instance)
		// Tag 1 - Opening Tag
		tag = 1
		e.openingTag(tag)

		e.properties(obj.Properties)

		// Tag 1 - Closing Tag
		e.closingTag(tag)
	}
	return nil
}

func (e *Encoder) properties(properties []bactype.Property) error {
	// for each property
	var tag uint8
	for _, prop := range properties {
		// Tag 0 - Property ID
		tag = 0
		e.contextPropertyID(tag, prop.ID)

		// Tag 1 (OPTIONAL) - Array Length
		if prop.ArrayIndex != ArrayAll {
			tag = 1
			e.contextUnsigned(tag, prop.ArrayIndex)
		}
	}
	return nil
}
