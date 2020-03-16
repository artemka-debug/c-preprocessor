package ifdef

import (
	tags "allcpptags"
	"fmt"
	"strings"
	"utils"
)

func Handle(allDefines map[string]map[string][]string, data []string, index int) *[]string {
	define := getDefine(data[index])
	var dataToReturn []string

	if _, ok := allDefines[define[0]]; ok {
		dataToReturn = data
	} else {
		dataToReturn = removeToEndIf(data, index)
	}

	return &dataToReturn
}

func removeToEndIf(data []string, index int) []string {
	dataCopy := data
	// A variable to handle nested ifdefs
	ifs := 1
	i := index + 1;
	dataCopy = utils.Remove(&dataCopy, index)

	for ifs != 0 {
		divided := strings.Split((dataCopy)[i], " ")

		if divided[0] == tags.IFDEF {
			ifs++
		} else if divided[0] == tags.ENDIF {
			ifs--
		}

		dataCopy = utils.Remove(&dataCopy, i)
	}
	//finalData = append(finalData, finalData[i:])

	fmt.Print("FINAL PROG", strings.Join(dataCopy, "\n"))

	return data
}

func getDefine(str string) []string {
	return (strings.Split(str, " "))[1:]
}
