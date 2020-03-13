package process

import (
	"handletag"
	"strings"
	"utils"
)

func Data(data []byte) string {
	lineByLineData := strings.Split(string(data), "\n")

	for i, v := range lineByLineData {
		if tag, ok := utils.CheckTag(v); ok {
			handletag.Tag(&lineByLineData, tag, i)
		}
	}

	return strings.Join(lineByLineData, "\n")
}