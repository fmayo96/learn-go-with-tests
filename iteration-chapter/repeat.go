package iteration

import "strings"

func Repeat(c string, repeatTimes int) string {
	var result strings.Builder
	for range repeatTimes {
		result.WriteString(c)
	}
	return result.String()
}
