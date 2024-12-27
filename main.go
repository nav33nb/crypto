package main

import (
	"fmt"

	rot13 "github.com/nav33nb/crypto/rot13"
)

func main() {
	enc := rot13.Encrypt("Hello World !!!")
	fmt.Println(enc)
	dec := rot13.Decrypt(enc)
	fmt.Println(dec)
}
