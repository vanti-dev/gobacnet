package property

import (
	"regexp"
	"strings"
)

// stringToID holds a mapping from the string representation of an ID back to that object type.
// The key is the value returned by ID.String lower-case with punctuation removed.
//
//	PresentValue, present-value, PRESENT VALUE => presentvalue
var stringToID map[string]ID

func init() {
	stringToID = make(map[string]ID, nextID)
	for i := 0; i < nextID; i++ {
		ot := ID(i)
		s := ot.String()
		if !strings.HasPrefix(s, "ID(") {
			stringToID[encodeKey(s)] = ot
		}
	}
}

// FromString returns an ID whose ID.String matches s ignoring case and non-alphanumeric characters.
// If none do then returns ID(0) and false.
func FromString(s string) (ID, bool) {
	o, ok := stringToID[encodeKey(s)]
	return o, ok
}

var nonWordRegex = regexp.MustCompile(`[^0-9a-z]+`) // applied against a lowercase string
func encodeKey(val string) string {
	val = strings.ToLower(val)
	return nonWordRegex.ReplaceAllString(val, "")
}
