package base64

import (
	"crypto/rand"
	b64 "encoding/base64"
	"testing"
)

const n = 65536

func Benchmark_1(b *testing.B) {
	x := make([]byte, n)
	y := make([]byte, b64.StdEncoding.EncodedLen(n))
	rand.Read(x)
	for b.Loop() {
		Encode(y, x)
		Decode(x, y)
	}
}
func Benchmark_2(b *testing.B) {
	x := make([]byte, n)
	y := make([]byte, b64.StdEncoding.EncodedLen(n))
	rand.Read(x)
	for b.Loop() {
		b64.StdEncoding.Encode(y, x)
		b64.StdEncoding.Decode(x, y)
	}
}
