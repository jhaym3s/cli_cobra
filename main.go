package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://www.google.com/",
		"https://github.com",
		"https://stackoverflow.com/",
		"https://www.instagram.com/",
	}
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	c := make(chan string)
	for _, v := range sites {
		go checkWebsiteStats(v, c)
	}
	for l := range c { 
		go func (link string)  {
			time.Sleep(time.Second)
			checkWebsiteStats(link,c)
		}(l)
		
	}
		
	

}
func checkWebsiteStats(s string, c chan string) {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println(s, "is inactive")
		c <- s
		panic(err)
	}
	fmt.Println(s, "is active")
	c <- s
	defer resp.Body.Close()
}
