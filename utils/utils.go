package utils

import "strings"

//QlSeparator ...
func QlSeparator(query string) map[string]string {
	urlMap := make(map[string]string)
	var s = strings.Split(query, "?")
	var s2 = strings.Split(s[1], "&")
	for i := 0; i < len(s2); i++ {
		s3 := strings.Split(s2[i], "=")

		urlMap[s3[0]] = s3[1]
	}
	return urlMap
}
