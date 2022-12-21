package utils

import "regexp"

func CheckRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}
