package random

import "testing"

func TestRandomString(t *testing.T) {
	t.Log(RandomString(4))
	t.Log(RandHexadecimal(4))
	t.Log(RandomDigitString(4))
	t.Log(RandomStringWithSeeds(5, PunctuationSeeds))
}
