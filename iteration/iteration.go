package iteration

// const repeatCount = 5

// Repeat repeats a character 5 times
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
