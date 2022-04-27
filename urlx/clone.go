package urlx

import "net/url"

func CloneUrlValues(v url.Values) url.Values {
	if v == nil {
		return nil
	}

	// Find total number of values.
	nv := 0
	for _, vv := range v {
		nv += len(vv)
	}
	sv := make([]string, nv) // shared backing array for headers' values
	v2 := make(url.Values, len(v))
	for k, vv := range v {
		n := copy(sv, vv)
		v2[k] = sv[:n:n]
		sv = sv[n:]
	}
	return v2
}
