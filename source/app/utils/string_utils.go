package utils

import "strings"

func SnakeToCamelCase(s string) string {
	words := strings.Split(s, "_")
	return toCamelCase(words)
}

func KebabToCamelCase(s string) string {
	words := strings.Split(s, "-")
	return toCamelCase(words)
}

func toCamelCase(words []string) string {
	if len(words) == 0 {
		return ""
	}

	for i := 1; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
	}
	return strings.Join(words, "")
}
