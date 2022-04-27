package random

import (
	"math/rand"
	"time"
)

var (
	DigitSeeds       = []byte("0123456789")
	UpperLetterSeeds = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LowerLetterSeeds = []byte("abcdefghijklmnopqrstuvwxyz")
	PunctuationSeeds = []byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
	Hexadecimal      = []byte("0123456789abcfefABCDEF")
)

func RandomString(length int) string {
	return randomString(append(append(DigitSeeds, UpperLetterSeeds...), LowerLetterSeeds...), length)
}

func RandHexadecimal(length int) string {
	return randomString(Hexadecimal, length)
}

func RandomDigitString(length int) string {
	return randomString(DigitSeeds, length)
}

func RandomStringWithSeeds(length int, seeds []byte) string {
	return randomString(seeds, length)
}

func randomString(seeds []byte, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]byte, length, length)
	for i := 0; i < length; i++ {
		slice[i] = seeds[r.Intn(len(seeds))]
	}
	return string(slice)
}
