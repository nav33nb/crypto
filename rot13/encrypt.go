package rot13

func Encrypt(text string) string {
	var floor rune
	var enc string
	for _, ch := range text {
		switch {
		case 'a' <= ch && ch <= 'z':
			floor = 'a'
		case 'A' <= ch && ch <= 'Z':
			floor = 'A'
		default:
			enc += string(ch)
			continue
		}
		ch = (((ch - floor) + 13) % 26) + floor
		enc += string(ch)
	}
	return enc
}
