package main

import (
	"fmt"
	//"io"
	"os"
	"path/filepath"
	"strings"
)

var ignoredFiles = []string{".git", "dockerfile",".","hw1.md"}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out interface{}, filePath string, printFiles bool) (error) {

	var prinPath string
	var pathList []string


	err := filepath.Walk(filePath, func(path string, file os.FileInfo, err error) error {

		if isFileIgnored(path) || (!file.IsDir() && !printFiles) {
			return nil
		}

		pathList = append(pathList, path)


		for _,line := range pathList{

			//тут нужно както отформатировать вывод
			prinPath = prinPath + buildDirStruct(line)
		}

		fmt.Println(prinPath)


		return nil
	})
	return err

}


func isFileIgnored(path string) bool{
	currentPathList :=strings.Split(path, string(os.PathSeparator))



	for _,ignored := range  ignoredFiles {

			if currentPathList[0] == ignored {

				return true
			}

	}
	return false
}

func buildDirStruct(path string) string{

	var DirStruct string

	pathListAll := strings.Split(path, string(os.PathSeparator))

	pathList := pathListAll[1:]
	if len(pathList) == 0 {
		return DirStruct
	}
	for index, item := range pathList {
		if index == (len(pathList) - 1) {
			DirStruct = DirStruct + `├───` + item
		} else {
			DirStruct = DirStruct +  "│\t"
		}
	}
	return DirStruct + "\n"

}
