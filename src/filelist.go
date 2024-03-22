package filelist

import (
	"bufio"
	"log"
	"os"
)

type fileObj struct {
	Name         string
	LineCount    int
	CommentCount int
}

type FileList []fileObj

func (f *FileList) Add(path string) {
	file := fileObj{
		Name:         path,
		LineCount:    CountLines(path),
		CommentCount: 0,
	}
	*f = append(*f, file)
}

func CountLines(path string) int {
	var count int
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		count++
		// fmt.Println(fileScanner.Text())
	}

	return count
}
