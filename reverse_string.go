package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	bytes := []rune(input)
	for i := 0; i < len(bytes)-1; i++ {
		for j := 0; j < len(bytes)-i-1; j++ {
			bytes[j], bytes[j+1] = bytes[j+1], bytes[j]
		}
	}
	output = string(bytes)
	return output
}
