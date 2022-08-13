package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel:=make(chan string)

	for _, link := range links {
		go checkLink(link,channel)
	}
	var receive string = <-channel
	fmt.Println(receive)
	fmt.Println(" all done! ")
}
func checkLink(link string,chn chan string) {
	fmt.Println(" check link function is called ")
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		chn <- link + " might be down "
	} else {
		fmt.Println(link, " is up ")
		chn <-link + " is working "
	}
}
