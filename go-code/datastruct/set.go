package datastruct

// Set datastruct solving for unique problem
type Set interface {
	// Exist checks a key is in Set
	Exist(...string) bool
	// Evict deletes matching key in Set
	Evict(...string)
	// Add inserts multi key into Set
	Add(...string)
	// String convert set key => string. Use , as separator
	String() []string
	// Intersection intersect other set into current set
	Intersection(Set)
	// Union union other set into current set
	Union(Set)
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
func (s set) Exist(keys ...string) bool {
	for i := 0; i < len(keys); i++ {
		if _, ok := s[keys[i]]; !ok {
			return false
		}
	}
	return true
}

// Has implements Set
func (set) Has(...string) bool {
	panic("unimplemented")
}

// String implements Set
func (s set) String() []string {
	str := []string{}
	for k := range s {
		str = append(str, k)
	}
	return str
}

// Intersection implements Set
func (s set) Intersection(otherSet Set) {
	for k := range s {
		if !otherSet.Exist(k) {
			s.Evict(k)
		}
	}
}

// Union implements Set
func (s set) Union(otherSet Set) {
	s.Add(otherSet.String()...)
}

// Evict implements Set
func (s set) Evict(keys ...string) {
	for i := 0; i < len(keys); i++ {
		delete(s, keys[i])
	}
}
