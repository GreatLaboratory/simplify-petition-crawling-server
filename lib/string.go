package lib

import "strings"

// CleanStr -> str을 깨끗한 문자열로 변환
func CleanStr(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
