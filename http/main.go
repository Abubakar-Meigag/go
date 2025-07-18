package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err :=  http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}

	// Ensure bs is large enough to hold the response body
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)
} 

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote: ", len(bs), "bytes")

	return len(bs), nil
}
