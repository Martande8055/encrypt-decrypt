package encription

// an identifier starts with uppercase letter is exported.
// an identifier starts with lowercase letter or underscore can
// only accessed from within package where it's declared.
func Encrypt(str string) string {

	var enc_str string = ""

	for _, i := range str {
		ass := int(i)
		encpt := string(ass + 3)
		enc_str += encpt
	}
	return enc_str
}
