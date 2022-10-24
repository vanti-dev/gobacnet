package objecttype

import (
	"regexp"
	"strings"
)

// stringToObjectType holds a mapping from the string representation of an ObjectType back to that object type.
// The key is the value returned by ObjectType.String lower-case with punctuation removed.
//
//	Analog Value, analog-value, ANALOG_VALUE => analogvalue
var stringToObjectType map[string]ObjectType

func init() {
	stringToObjectType = make(map[string]ObjectType, nextObjectType)
	for i := 0; i < nextObjectType; i++ {
		ot := ObjectType(i)
		s := ot.String()
		if !strings.HasPrefix(s, "ObjectType(") {
			stringToObjectType[encodeKey(s)] = ot
		}
	}
}

// FromString returns an ObjectType whose ObjectType.String matches s ignoring case and non-alphanumeric characters.
// If none do then returns ObjectType(0) and false.
func FromString(s string) (ObjectType, bool) {
	o, ok := stringToObjectType[encodeKey(s)]
	return o, ok
}

var nonWordRegex = regexp.MustCompile(`[^0-9a-z]+`) // applied against a lowercase string
func encodeKey(val string) string {
	val = strings.ToLower(val)
	return nonWordRegex.ReplaceAllString(val, "")
}
