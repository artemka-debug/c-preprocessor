package utils

import (
	t "allcpptags"
	"strings"
)

func CheckTag(v string) (string, bool) {
	tag := strings.Split(v, " ")

	if tag[0] == t.DEFINE {
		return t.DEFINE, true
	} else if tag[0] == t.IF {
		return t.IF, true
	} else if tag[0] == t.ENDIF {
		return t.ENDIF, true
	} else if tag[0] == t.IFDEF {
		return t.IFDEF, true
	}

	return "", false
}