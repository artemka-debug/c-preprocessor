package define

import (
	"strings"
	"utils"
)

func ValEqExpVal(val string, expVal string) bool {
	if val == expVal {
		return true
	} else if strings.Index(val, "(") != -1 {
		for i := 0; i < utils.Min(len(val), len(expVal)) && expVal[i] == val[i]; i++ {
			if (val[i + 1] == '(' || val[i + 1] == ';') && len(expVal) == i + 1 {
				return true
			}
		}
	}

	return false
}


