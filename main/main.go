package main

import (
	b64 "encoding/base64"

	"github.com/birowo/base64"
)

func main() {
	x := []byte("QWERTYUIOPASDFGHJKLZXCVBNMqwerty")
	xLen := len(x)
	y := make([]byte, b64.StdEncoding.EncodedLen(xLen))
	z := make([]byte, xLen)

	for i := 24; i < 33; i++ {
		j := b64.StdEncoding.EncodedLen(i)
		base64.Encode(y[:j], x[:i])
		n, _ := b64.StdEncoding.Decode(z[:i], y[:j])
		println("x:", string(x[:i]), "y:", string(y[:j]), "z:", string(z[:n]))
		if string(x[:i]) != string(z[:n]) {
			println("FAIL")
		}
	}
	println()
	for i := 24; i < 33; i++ {
		j := b64.StdEncoding.EncodedLen(i)
		b64.StdEncoding.Encode(y[:j], x[:i])
		n, _ := base64.Decode(z[:i], y[:j])
		println("x:", string(x[:i]), "y:", string(y[:j]), "z:", string(z[:n]))
		if string(x[:i]) != string(z[:n]) {
			println("FAIL")
		}
	}
}
