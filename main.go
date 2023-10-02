package main

import (
	"flag"
	"fmt"
)


func main()  {
	dice := flag.String("d","d6","The type of dice to roll. Format is DX where X is an int. Default is D6")
	flag.Parse()
	fmt.Println(*dice)
}