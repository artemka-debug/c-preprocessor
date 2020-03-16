package process

import (
	tags "allcpptags"
	"handletag"
	"strings"
	"utils"
)

func Data(data []byte) string {
	lineByLineData := strings.Split(string(data), "\n")
	AllDefines := make(map[string]map[string][]string)
	var indToDelete []int
	test := 0
	var finalProg *[]string

	for i, v := range lineByLineData {
		if tag, ok := utils.CheckTag(v); ok {
			if tag == tags.DEFINE || tag == tags.IFDEF {
				finalProg = handletag.Tag(AllDefines, lineByLineData, tag, i)
			} else {
				finalProg = handletag.Tag(nil, lineByLineData, tag, i)
			}
			indToDelete = append(indToDelete, i)
		}
	}


	//fmt.Printf("DEFIENS [%+v]\n", finalProg)

	for i := range *finalProg {
		if utils.ItemExists(indToDelete, i) {
			utils.Remove(finalProg, i - test)
			test++
		}
	}

	return strings.Join(*finalProg, "\n")
}