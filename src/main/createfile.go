package main

import (
	"os"
	"fmt"
)

func CreateFile(data string) {
	f, err := os.Create("myprog.c")

	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(data)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	ferr := f.Close()

	if ferr != nil {
		fmt.Println(ferr, l)
		return
	}		
}