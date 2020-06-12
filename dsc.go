package main

import (
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Exit(2)
	}
	for _, f := range files {
		println(f.Name())
	}

}
