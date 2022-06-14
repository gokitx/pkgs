package slicex

import (
	"strings"
	"testing"
)

func TestRevemoRepByMap(t *testing.T) {
	v1 := []int{1, 1, 2, 2, 2, 3, 4}
	t.Log(RevemoRepByMap(v1))

	v2 := []float32{1.1, 1.1, 2, 2.2, 2, 3, 4}
	t.Log(RevemoRepByMap(v2))

	v3 := []string{"a", "a", "c", "d", "b", "e"}
	t.Log(RevemoRepByMap(v3))
}

func TestRevemoRepWithSort(t *testing.T) {
	v1 := []int{1, 1, 2, 2, 2, 3, 4}
	t.Log(RevemoRepWithSort(v1))

	v2 := []float32{1.1, 1.1, 2, 2.2, 2, 3, 4}
	t.Log(RevemoRepWithSort(v2))

	v3 := []string{"a", "a", "c", "d", "b", "e"}
	t.Log(RevemoRepWithSort(v3))

	v4 := []string{"a", "b", "c", "d", "b", "a", "a"}
	t.Log(RevemoRepWithSort(v4))
}

func TestContainsIn(t *testing.T) {
	v1 := []int{1, 1, 2, 2, 2, 3, 4}
	t.Log(ContainsIn(v1, 2))

	v2 := []float32{1.1, 1.1, 2, 2.2, 2, 3, 4}
	t.Log(ContainsIn(v2, 2.3))

	v3 := []string{"aac", "abc", "c", "d", "b", "e"}
	t.Log(ContainsIn(v3, "e"))

	t.Log(ContainsIn(v3, "ab", func(v, k string) bool { return strings.Contains(v, k) }))
}
