package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	
	dice := flag.String("d", "d4", "The type of dice to roll. Format is DX where X is an int. Default is D6")
	numRoll := flag.Int("n",1,"The number of die rolls. Default is 1")
	flag.Parse() //After defining all flags call the function "flag.Parse()" to parse the command line into the defined flags.
	matched, _ := regexp.Match("d\\d+$", []byte(*dice))
	if matched {
		stringDiceSize := (*dice)[1:]
		diceSize, _ := strconv.Atoi(stringDiceSize)
		for i := 0; i < *numRoll; i++ {
			roll := rand.Intn(diceSize)+1
			fmt.Printf("You chose %d \n", roll)	
		}
		
	} else {
		fmt.Printf("Something is wrong dice format is dx  %s \n", *dice)
	}
}
