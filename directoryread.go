package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("module_app",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			s := ""
			words := strings.Split(path, "\\")
			length := len(words)
			for i := 0; i < length-1; i++ {
				s += "----"
			}
			s += ">" + words[length-1]
			fmt.Println(s)

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
