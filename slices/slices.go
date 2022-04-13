package slices

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
