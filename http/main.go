package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// body := make([]byte, 9000)
	// resp.Body.Read(body)

	// fmt.Println(string(body))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %d bytes \n", len(bs))
	return len(bs), nil
}
