package handletag

import (
	tags "allcpptags"
	_ "fmt"
	"handletag/define"
	"handletag/ifdef"
)

func Tag(allDefines map[string]map[string][]string, data []string, tag string, index int) *[]string {
	if tag == tags.IF {

	} else if tag == tags.ELSE {

	} else if tag == tags.ENDIF {

	} else if tag == tags.DEFINE {
		return define.Handle(allDefines, data, index)
	} else if tag == tags.IFDEF {
		ifdef.Handle(allDefines, data, index)
	}

	return &data
}
 