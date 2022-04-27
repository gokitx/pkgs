package slicex

func RevemoRepByMap[T comparable](v []T) []T {
	state := make(map[T]struct{})
	for _, item := range v {
		_, ok := state[item]
		if !ok {
			state[item] = struct{}{}
		}
	}

	r := make([]T, 0)
	for item := range state {
		r = append(r, item)
	}
	return r
}

func ContainsIn[T comparable](v []T, sub T,
	f ...func(T, T) bool) bool {
	if len(f) != 0 {
		for _, vv := range v {
			if f[0](vv, sub) {
				return true
			}
		}
		return false
	}

	for _, vv := range v {
		if vv == sub {
			return true
		}
	}
	return false
}
