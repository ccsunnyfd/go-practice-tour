package word

import "strings"

//
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

//
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

//
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}
