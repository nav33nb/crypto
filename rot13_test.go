// Testing code for the ROT13 module
package rot13

import (
	"fmt"
	"testing"
)

// Uses a very common string (with whitespace) with each alphabet to check for encoding/decoding
func TestRot13(t *testing.T) {
	str := "the quick brown fox jumps over the lazy dog"
	ans := "gurdhvpxoebjasbkwhzcfbiregurynmlqbt"
	enc := Rot13(str)
	dec := Rot13(enc)
	fmt.Println("enc == ans && dec == str ?", enc == ans && dec == str)

	if enc != ans || dec != str {
		t.Errorf("\nEncrypt(%q) = %q; want %q\nDecrypt(%q) = %q; want %q", str, enc, ans, enc, str, dec)
	}
}
