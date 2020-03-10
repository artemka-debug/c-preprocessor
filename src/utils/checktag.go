package utils

import (
	"fmt"
	"strings"
	t "allcpptags"
)

func CheckTag(v string) (string, bool) {
	if strings.HasPrefix(v, t.DEFINE) {
		fmt.Println(t.DEFINE)
		return t.DEFINE, true
	}

	return "", false
}