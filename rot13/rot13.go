package rot13

func Rot13(input string) string {
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
