package define

func Handle(data *[]string, index int) *[]string {
	expVar, expVal, expParams := parseDefine((*data)[index])
	//fmt.Printf("RESPONSE [%s] [%s] [%s]\n", expVar, expVal, expParams)
	swapCallsWithDefines(data, index, expVar, expVal, expParams)

	return data
}
