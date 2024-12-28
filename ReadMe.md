# Rather simple rot13 implementation in Go

A simple module i created in go to play with publishing, exports and versioning in Go
 
> In the highly unlikely (not recommended) case you want to use this for your production grade applications
```sh
go mod init # if not done already
go get github.com/nav33nb/crypto
```

Link - [https://pkg.go.dev/github.com/nav33nb/ncrypt](https://pkg.go.dev/github.com/nav33nb/ncrypt)

## Usage
```go
package main

import (
	"fmt"

	"github.com/nav33nb/ncrypt/rot13"
)

func main() {
	str := "Hello"
	enc := rot13.Encode(str)
	dec := rot13.Decode(enc)

	fmt.Println("str == dec ?", str == dec)
}

// OUTPUT
// str == dec ? true
```
