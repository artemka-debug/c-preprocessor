package define

func GetParamVal(paramName []string, paramVals string) map[string]string {
	paramVal := make(map[string]string)
	lastComma := 0
	index := 0

	for i, c := range paramVals {
		if c == ',' {
			paramVal[paramName[index]] = paramVals[lastComma:i]
			lastComma = i + 1
			index += 1
		}
	}
	paramVal[paramName[index]] = paramVals[lastComma:]

	return paramVal
}
