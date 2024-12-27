package rot13

func Decrypt(text string) string {
	var factor rune
	var enc string
	for _, ch := range text {
		switch {
		case 97 <= ch && ch <= 122:
			factor = 97
		case 65 <= ch && ch <= 90:
			factor = 65
		default:
			enc += string(ch)
			continue
		}
		ch = (((ch - factor) + 13) % 26) + factor
		enc += string(ch)
	}
	return enc
}
