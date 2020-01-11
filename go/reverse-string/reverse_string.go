package reverse

// Reverse a string
func Reverse(input string) string {
	output := ""
	characters := []rune(input)

	for _, character := range characters {
		output = string(character) + output
	}

	return output
}
