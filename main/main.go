package main

import (
	"github.com/birowo/base64"
)

func main() {
	x := []byte("QWERTYUIOPASDFGHJKLZXCVB")
	y := make([]byte, 32)
	base64.Encode(y, x)
	println("y:", string(y))
	z := make([]byte, 24)
	base64.Decode(z, y)
	println("z:", string(z))
	if string(x) != string(z) {
		println("FAIL")
	}
}
