package iteration

import "strings"

func Repeat(char string, number int) string {
	var repeated strings.Builder
	for i := 0; i < number; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
