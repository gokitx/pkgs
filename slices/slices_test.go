package slices

import "testing"

func TestRevemoRepByMap(t *testing.T) {
	v1 := []int{1, 1, 2, 2, 2, 3, 4}
	t.Log(RevemoRepByMap(v1))

	v2 := []float32{1.1, 1.1, 2, 2.2, 2, 3, 4}
	t.Log(RevemoRepByMap(v2))

	v3 := []string{"a", "a", "c", "d", "b", "e"}
	t.Log(RevemoRepByMap(v3))
}
