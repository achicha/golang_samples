/*
learn interfaces. part2.
http request
NB! this is not work with `-i` in `Go tool instruments`
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	// http request
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// reading resp.Body text

	// 1-way. use raw `Reader` interface to read INTO the byte  slice
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//resp.Body.Close()

	// 2-way. use io.Copy() func, that implements `Writer` interface. input: []byte, output: terminal
	//io.Copy(os.Stdout, resp.Body)

	// 3-way. using built-in function ioutil.ReadAll()
	//fmt.Println(resp.StatusCode)
	//bodyText, err := ioutil.ReadAll(resp.Body)
	//resp.Body.Close()
	//
	//if err != nil {
	//   //log.Fatal(err)
	//   fmt.Println("Error: ", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Printf("%s", bodyText)

	// 4-way. custom implementation of the `Writer`
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
