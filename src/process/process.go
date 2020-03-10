package process

import (
	"strings"
	"utils"
	handle "handletag"
)

func Data(data []byte) (string) {
	lineByLineData := strings.Split(string(data), "\n")

	for i, v := range lineByLineData {
		if tag, ok := utils.CheckTag(v); ok {
			handle.Tag(&data, tag, i)
		}
	}

	return string(data)
}