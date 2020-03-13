package utils

import (
	"strings"
	t "allcpptags"
)

func CheckTag(v string) (string, bool) {
	if strings.HasPrefix(v, t.DEFINE) {
		return t.DEFINE, true
	} else if strings.HasPrefix(v, t.IF) {
		return t.IF, true
	} else if strings.HasPrefix(v, t.ENDIF) {
		return t.ENDIF, true
	}

	return "", false
}