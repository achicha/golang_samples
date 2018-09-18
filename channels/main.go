/*learn channels. wep page availability checker*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://facebook123123.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// create a channel
	c := make(chan string)

	// create Go Routines for each fetch call
	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		get link from Child GO Routine using Channel
		`range c` = we are waiting for result from the channel
		`func()` is a function literal. == anonymous function or lambda in python.
	*/
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	var msg string

	// make request
	resp, err := http.Get(link)

	// handle response/errors
	if err != nil {
		msg = link + " is DOWN"

	} else {
		msg = link + " is UP. status=" + string(resp.Status)
	}

	// output
	fmt.Println(msg)
	c <- link // push our link to the channel (to the MAIN go routine)
}
