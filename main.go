package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dice := flag.String("d", "d4", "The type of dice to roll. Format is DX where X is an int. Default is D6")
	numRoll := flag.Int("n", 1, "The number of die rolls. Default is 1")
	sum := flag.Bool("s",false,"This gives you the sum of all the  rolled dice")
	adv := flag.Bool("adv", false,"This gives you the highest number of all the  rolled dice")
	disAd:= flag.Bool("disAd", false,"This gives you the lowest number of all the  rolled dice")
	flag.Parse() //After defining all flags call the function "flag.Parse()" to parse the command line into the defined flags.
	matched, _ := regexp.Match("d\\d+$", []byte(*dice))
	if matched {
		value := rollDice(dice, numRoll)
		if *sum {
			fmt.Printf("Sum of dice is %d\n ", sumDice(value))
		} 
		if *adv{
			advantage(value)
		}
		if *disAd{
			disAdvantage(value)
		}
		printDice(value)
	} else {
		fmt.Printf("Something is wrong dice format is dx  %s \n", *dice)
	}
}

func rollDice(dice *string, numRoll *int) []int {
	var rolls []int
	stringDiceSize := (*dice)[1:]
	diceSize, _ := strconv.Atoi(stringDiceSize)
	for i := 0; i < *numRoll; i++ {
		roll := rand.Intn(diceSize) + 1
		fmt.Printf("You chose %d \n", roll)
		rolls = append(rolls, roll)
	}
	return rolls
}

func printDice(roll []int) {
	for i, v := range roll {
		fmt.Printf("Roll %d was %d\n", i+1, v)
	}

}
func sumDice(dice []int ) int {
	sum := 0
	for _, v := range dice {
		sum += v
	}
	return sum
}

func advantage(dice []int)  {
	sort.Ints(dice)
	fmt.Printf("advantage is %d", dice[len(dice)-1])
}
func disAdvantage(dice []int)  {
	sort.Ints(dice)
	fmt.Printf("advantage is %d", dice[0])
}
