package main

import (
	"fmt"
	"github.com/hunttraitor/linecounter/src"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Error: must pass in one path")
	}

	path, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal(err)
	}

	files := &filelist.FileList{}

	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files.Add(path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}
