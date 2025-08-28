package main

import (
	b64 "encoding/base64"

	"github.com/birowo/base64"
)

func main() {
	x := []byte("QWERTYUIOPASDFGHJKLZXCVBNMqwertyui")
	xLen := len(x)
	y := make([]byte, b64.StdEncoding.EncodedLen(xLen))
	base64.Encode(y, x)
	z := make([]byte, xLen)
	n, _ := base64.Decode(z, y)
	println("x:", string(x), "y:", string(y), "z:", string(z[:n]))
	if string(x) != string(z[:n]) {
		println("FAIL")
	}
	b64.StdEncoding.Encode(y, x)
	n, _ = b64.StdEncoding.Decode(z, y)
	println("x:", string(x), "y:", string(y), "z:", string(z[:n]))
	if string(x) != string(z[:n]) {
		println("FAIL")
	}
	println("yLen:", len(y))
}
