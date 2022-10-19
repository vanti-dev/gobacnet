package objecttype

import "strings"

var stringToObjectType map[string]ObjectType

func init() {
	stringToObjectType = make(map[string]ObjectType, nextObjectType)
	for i := 0; i < nextObjectType; i++ {
		ot := ObjectType(i)
		s := ot.String()
		if !strings.HasPrefix(s, "ObjectType(") {
			stringToObjectType[s] = ot
		}
	}
}

// FromString returns an ObjectType whose ObjectType.String matches s exactly.
// If none do then
func FromString(s string) ObjectType {
	if ot, ok := stringToObjectType[s]; ok {
		return ot
	}
	return ObjectType(0)
}
