package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"process"
)

func main() {
	data, err := ioutil.ReadFile((os.Args[1:])[0])

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		CreateFile(string(process.Data(data)))
	}
}