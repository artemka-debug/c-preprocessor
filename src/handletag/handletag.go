package handletag

import (
	tags "allcpptags"
	_ "fmt"
	"handletag/define"
	iftag "handletag/if"
)

func Tag(data *[]string, tag string, index int) *[]string {
	if tag == tags.IF {
		iftag.Handle(data, index)
	} else if tag == tags.ELSE {

	} else if tag == tags.ENDIF {

	} else if tag == tags.DEFINE {
		return define.Handle(data, index)
	}

	return data
}
 