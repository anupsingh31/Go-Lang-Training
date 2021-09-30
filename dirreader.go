package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	root := "module_app"
	fmt.Println(">", root)
	IOReadDir(0, root)
}
func IOReadDir(count int, root string) {
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			count++
			for i := 0; i < count; i++ {
				fmt.Print("----")
			}
			fmt.Println(">", file.Name())
			root += "\\" + file.Name()
			IOReadDir(count, root)
			s := strings.LastIndex(root, "\\")
			root = root[:s]
			count--
		} else {
			for i := 0; i < count; i++ {
				fmt.Print("----")
			}
			fmt.Println("---->", file.Name())
		}
	}
}
