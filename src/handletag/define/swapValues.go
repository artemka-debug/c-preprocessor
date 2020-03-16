package define

import (
	"strings"
)

func swapCallsWithDefines(data []string, index int, expVar string, expVal string, expParams []string) {
	for i, str := range (data)[index + 1:] {
		divided := strings.Split(str, " ")

		for ind, v := range divided {
			if len(v) != 0 && (v[len(v) - 1] == ';') {
				val := v[:len(v) - 1]

				if ValEqExpVal(val, expVar) {

					swappedExpVal := ""
					if expParams != nil {
						swappedExpVal = swapValues(expVal, val, expParams)
					} else {
						swappedExpVal = expVal
					}

					divided[ind] =  swappedExpVal + string(divided[ind][len(divided[ind]) - 1])
					(data)[index + i + 1] = strings.Join(divided, " ")
				}
			}
		}
	}
}

func swapValues(expVal string, val string, expParams []string) string {
	expToParse := val[strings.Index(val, "(") + 1:strings.Index(val, ")")]

	paramsWithVals := GetParamVal(expParams, expToParse)
	//fmt.Printf("[%+v]\n", paramsWithVals)
	expValCopy := strings.Split(expVal, " ")

	for i, v := range expValCopy {
		if val, ok := paramsWithVals[v]; ok {
			expValCopy[i] = val
		}
	}

	return strings.Join(expValCopy, " ")
}

