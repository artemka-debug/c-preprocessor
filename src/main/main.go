package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"process"
)

func main() {
	if len(os.Args) != 1 {
		data, err := ioutil.ReadFile((os.Args[1:])[0])
	
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			CreateFile(process.Data(data))
		}
	} else {
		fmt.Println("Path to file was not provided")
	}
}