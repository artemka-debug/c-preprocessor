package process

import (
	"fmt"
	"handletag"
	"strings"
	"utils"
)

func Data(data []byte) string {
	lineByLineData := strings.Split(string(data), "\n")
	AllDefines := make(map[string]map[string][]string)

	for i, v := range lineByLineData {
		if tag, ok := utils.CheckTag(v); ok {
			if tag == "#define" {
				handletag.Tag(AllDefines, &lineByLineData, tag, i)
			} else {
				handletag.Tag(nil, &lineByLineData, tag, i)
			}
		}
	}
	fmt.Printf("DEFIENS [%+v]\n", AllDefines)

	return strings.Join(lineByLineData, "\n")
}