package encription

func Encrypt(str string) string {

	var enc_str string = ""

	for _, i := range str {
		ass := int(i)
		encpt := string(ass + 3)
		enc_str += encpt
	}
	return enc_str
}
