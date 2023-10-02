package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {
	dice := flag.String("d", "d4", "The type of dice to roll. Format is DX where X is an int. Default is D6")
	flag.Parse() //After defining all flags call the function "flag.Parse()" to parse the command line into the defined flags.
	matched, _ := regexp.Match("d\\d+$", []byte(*dice))
	if matched {
		fmt.Printf("You rolled %s\n", *dice)
	}else{
		fmt.Printf("Something is wrong with  %s value \n", *dice)
	}
}
