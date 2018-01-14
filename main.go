package main

import (
	"net/http"
	"fmt"
)

func main(){
	links := []string{
		"https://google.com",
		"https://flipkart.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://netflix.com",
	}
	c := make(chan string)
	for _, link := range links{
		go checkLink(link, c)
	}

	// Reading a  message from channel is blocking operation
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	for {
		go checkLink(<- c, c)
	}
}

func checkLink(link string, c chan string){
	if _, err := http.Get(link); err != nil{
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}