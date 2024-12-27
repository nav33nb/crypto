package rot13

func Decrypt(text string) string {
	var floor rune
	var dec string
	for _, ch := range text {
		switch {
		case 'a' <= ch && ch <= 'z':
			floor = 'a'
		case 'A' <= ch && ch <= 'Z':
			floor = 'A'
		default:
			dec += string(ch)
			continue
		}
		ch = (((ch - floor) + 13) % 26) + floor
		dec += string(ch)
	}
	return dec
}
