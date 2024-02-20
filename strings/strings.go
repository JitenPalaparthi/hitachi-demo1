package strings

// Reverse function reverse the string that is passed as an input argument
func Reverse(str string) string {
	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr, err
}
