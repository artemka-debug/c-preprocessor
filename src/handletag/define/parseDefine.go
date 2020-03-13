package define

func parseDefine(expression string) (string, string, []string) {
	exp := expression[8:]
	expVar := ""
	expVal := ""
	var expParams []string

	for i := 0; i < len(exp); i++ {
		if exp[i] == ' '{
			expVar = exp[:i]
			expVal = exp[i + 1:]

			return expVar, expVal, expParams
		} else if  exp[i] == '(' {
			expVar = exp[:i]
			lastComma := i + 1
			for exp[i] != ')' {
				if exp[i] == ',' {
					expParams = append(expParams, exp[lastComma:i])
					lastComma = i + 2
				}
				i++
			}
			expParams = append(expParams, exp[lastComma:i])

			i += 2

			for i != len(exp) {
				expVal += string(exp[i])
				i++
			}

			return expVar, expVal, expParams
		}
	}

	return expVar, expVal, expParams
}
