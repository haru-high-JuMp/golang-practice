package iteration

// const repeatedCount = 5
func Repeat(char string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated = repeated + char
	}
	return repeated
}
