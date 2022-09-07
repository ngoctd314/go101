package datastruct

import "strings"

// Set datastruct solving for unique problem
type Set interface {
	// Exist checks a key is in Set
	Exist(string) bool
	// Has checks multi key is in Set
	Has(...string) bool
	// Add inserts multi key into set
	Add(...string)
	// String convert set key => string. Use , as separator
	String() string
}

type set map[string]struct{}

var _voidstruct = struct{}{}

// NewSet implement Set interface
func NewSet() Set {
	return make(set)
}

// Add implements Set
func (s set) Add(keys ...string) {
	for i := 0; i < len(keys); i++ {
		s[keys[i]] = _voidstruct
	}
}

// Get implements Set
func (set) Exist(key string) bool {
	panic("unimplemented")
}

// Has implements Set
func (set) Has(...string) bool {
	panic("unimplemented")
}

// String implements Set
func (s set) String() string {
	str := []string{}
	for k := range s {
		str = append(str, k)
	}
	return strings.Join(str, ",")
}
