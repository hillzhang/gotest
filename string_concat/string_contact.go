package string_concat

func ConcatOperator(original *string, concat string) {
	// This could be written as 'return *original + concat'
	// but I wanted to confirm no special compiler optimizations
	// existed for concatenating a string to itself.
	*original = *original + concat
}

func SelfConcatOperator(input string, n int) string {
	output := ""
	for i := 0; i < n; i++ {
		ConcatOperator(&output, input)
	}
	return output
}