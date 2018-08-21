package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of argument. Please provide 1 filename as argument.")
		os.Exit(1)
	}

	filename := os.Args[1]

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot read file: ", err)
		os.Exit(2)
	}

	fmt.Print(string(b))

	file, err := os.Open(filename)
	io.Copy(os.Stdout, file)
}
