# Rather simple rot13 implementation in Go

A simple module i created in go to play with publishing, exports and versioning in Go
 
> In the highly unlikely (not recommended) case you want to use this for your production grade applications
```sh
go mod init # if not done already
go get github.com/nav33nb/crypto
```

## Usage
```
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

// OUTPUT
// str == dec ? true
```