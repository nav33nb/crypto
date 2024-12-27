// A simple letter substitution cipher, commonly called ROT13. It works by shifting any of the 26 alphabets by 13 characters, such that ABC -> NOP and vice versa. Same algorithm provides encoding and decoding.
package rot13

// Encode wraps the process function to encode a string.
func Encode(input string) string {
	return process(input)
}

// Decode wraps the process function to decode a string.
func Decode(input string) string {
	return process(input)
}

// Accepts an input string to encode/decode, Returns processed string
func process(input string) string {
	var floor rune
	var processed string
	for _, ch := range input {
		switch {
		case 'a' <= ch && ch <= 'z':
			floor = 'a'
		case 'A' <= ch && ch <= 'Z':
			floor = 'A'
		default:
			processed += string(ch)
			continue
		}
		ch = (((ch - floor) + 13) % 26) + floor
		processed += string(ch)
	}
	return processed
}
