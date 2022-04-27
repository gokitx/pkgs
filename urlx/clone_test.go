package urlx

import (
	"net/url"
	"testing"
)

func TestCloneUrlValues(t *testing.T) {
	v := make(url.Values)
	v["xx"] = []string{"1", "2"}
	v["xz"] = []string{"3"}

	v2 := CloneUrlValues(v)

	v["xx"] = append(v["xx"], "33")

	t.Log(v)
	t.Log(v2)
}
