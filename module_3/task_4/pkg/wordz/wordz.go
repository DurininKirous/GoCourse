package wordz

import (
	"crypto/rand"
	"math/big"
)

var Words = []string{"one", "two", "three", "four", "five", "six", "seven"}

func Random() string {
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Words[index]
}
