package decription

func Decript(str string) string {
	dec_str := ""

	for _, i := range str {
		ass := int(i)
		chr := ass - 3
		dec_str += string(chr)
	}
	return dec_str
}
