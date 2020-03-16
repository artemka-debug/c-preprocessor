package define

func Handle(allDefines map[string]map[string][]string, data []string, index int) *[]string {
	expVar, expVal, expParams := parseDefine((data)[index])
	//fmt.Printf("RESPONSE [%s] [%s] [%s]\n", expVar, expVal, expParams)

	val, ok := allDefines[expVar]
	if !ok {
		val = make(map[string][]string)
		allDefines[expVar] = val
	}
	val[expVal] = expParams

	swapCallsWithDefines(data, index, expVar, expVal, expParams)

	return &data
}
