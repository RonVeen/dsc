package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var indentSize = 0

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	dir := os.Args[1]
	processDir(dir, "")
	fmt.Println("Done!")
}

func processDir(dir string, path string) {

	temppath := dir
	if  path != "" {
		temppath = path + "/" + temppath
	}
	fmt.Println(temppath)
	indentSize++
	var directorySize int64 = 0
	files, _ := readDirectory(temppath)
	for _, f := range files {
		if f.IsDir() {
			processDir(f.Name(), temppath)
		} else {
			processFile(f)
			directorySize += f.Size()
		}
	}
	fmt.Printf("Total size of %v: %v\n", temppath, directorySize)
	indentSize--
}

func readDirectory(dirname string ) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func processFile(f os.FileInfo) {
	tabs := ""
	for i:= 0; i < indentSize; i++ {
		tabs += "\t"
	}
	// fmt.Println(tabs, f.Name())
}