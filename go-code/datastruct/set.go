package datastruct

// Set datastruct solving for unique problem
type Set map[string]struct{}

var _voidstruct = struct{}{}

// NewSet implement Set interface
func NewSet() Set {
	return make(Set)
}

// Add implements Set
func (s Set) Add(keys ...string) {
	for i := 0; i < len(keys); i++ {
		s[keys[i]] = _voidstruct
	}
}

// Exist implements Set
func (s Set) Exist(keys ...string) bool {
	for i := 0; i < len(keys); i++ {
		if _, ok := s[keys[i]]; !ok {
			return false
		}
	}
	return true
}

// Has implements Set
func (Set) Has(...string) bool {
	panic("unimplemented")
}

// String implements Set
func (s Set) String() []string {
	str := []string{}
	for k := range s {
		str = append(str, k)
	}
	return str
}

// Intersection implements Set
func (s Set) Intersection(otherSet Set) {
	for k := range s {
		if !otherSet.Exist(k) {
			s.Evict(k)
		}
	}
}

// Union implements Set
func (s Set) Union(otherSet Set) {
	s.Add(otherSet.String()...)
}

// Evict implements Set
func (s Set) Evict(keys ...string) {
	for i := 0; i < len(keys); i++ {
		delete(s, keys[i])
	}
}
