package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	data := []byte("hello to go")
	err := ioutil.WriteFile("ourFile1.txt", data, 1)
	checkError(err)
	fmt.Println("Done Writing")

	f, er := os.Create("ourFile2.txt")
	defer f.Close()
	checkError(er)
	bytesWritter, e := f.WriteString("Hello")
	checkError(e)
	fmt.Println("Bytes writtens are : ", bytesWritter)
	fmt.Println("Done Writing")

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
