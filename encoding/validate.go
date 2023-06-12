package encoding

import (
	"fmt"

	"github.com/vanti-dev/gobacnet/property"

	"github.com/vanti-dev/gobacnet/types"
)

func isValidObjectType(idType types.ObjectType) error {
	if idType > MaxObject {
		return fmt.Errorf("object types is %d which must be less then %d", idType, MaxObject)
	}
	return nil
}

func isValidPropertyType(propType property.ID) error {
	if propType > MaxPropertyID {
		return fmt.Errorf("object types is %d which must be less then %d", propType, MaxPropertyID)
	}
	return nil
}
