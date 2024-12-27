package main

import (
	"fmt"

	"github.com/nav33nb/crypto/rot13"
)

func main() {
	str := "Hello"
	enc := rot13.Encrypt(str)
	dec := rot13.Decrypt(enc)
	fmt.Println("str == dec ?", str == dec)
}
