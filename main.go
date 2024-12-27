package main

import (
	"fmt"

	"github.com/nav33nb/crypto/rot13"
)

func main() {
	str := "Hello"
	enc := rot13.Rot13(str)
	dec := rot13.Rot13(enc)
	fmt.Println("str == dec ?", str == dec)
}
