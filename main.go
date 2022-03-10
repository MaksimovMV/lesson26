package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func concatenate(fileNames []string) {
	fileContents := make([]string, 0)

	for i, fileName := range fileNames {
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 444)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if i < 2 {
			content, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			fileContents = append(fileContents, string(content))
		} else {
			resultContent := strings.Join(fileContents, " ")
			_, err := file.WriteString(resultContent)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	resultContent := strings.Join(fileContents, " ")
	fmt.Println(resultContent)
}

func main() {
	fileNames := os.Args[1:]
	concatenate(fileNames)
}
