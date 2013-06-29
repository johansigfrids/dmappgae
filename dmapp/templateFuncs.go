package dmapp

import (
	"strconv"
)

func modInt(i int) string {
	s := strconv.Itoa(i)
	if i >= 0 {
		return "+" + s
	}
	return s
}

func CommaList(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	var result string
	for _, s := range ss {
		result = result + s + ", "
	}
	return result[:len(result)-2]
}
