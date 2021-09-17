package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("ourFile.txt")
	checkError(err)
	fmt.Println("Content of the file is : ", string(content))

	f, er := os.Open("ourFile.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 10)
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("Content of the file(limited) : ", string(bucket[:bytesRead]))
	fmt.Println("Bytes read : ", bytesRead)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
