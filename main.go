package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problems struct {
	q string
	a string
}

func main() {
	quizFileName := flag.String("quiz", "quiz.cvs", "")
	timeLimit := flag.Int("limit",30,"Seconds for the quiz")
	flag.Parse()
	cvsFile, err := os.Open(*quizFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open your CVS file %s", *quizFileName))
	} else {
		cvsReader := csv.NewReader(cvsFile)
		cont, err := cvsReader.ReadAll()
		
		if err != nil {
			exit("Something is wrong with your cvs file")
		} else {
			problems := parseLines(cont)
			timer := time.NewTimer(time.Duration(*timeLimit)* time.Second)
			correctAnswers := 0
			for i, p := range problems {
				fmt.Printf("Question %d %s = \n", i+1, p.q)
				ansChan := make(chan string)
				go func ()  {
					var answer string
				fmt.Scanf("%s\n", &answer)
				ansChan <- answer
				}()
				select{
				case <-timer.C:
					fmt.Printf("You had %d correct out of %d\n",correctAnswers,len(problems))
					return // if we use a break here
				case answer := <-ansChan:
				if answer == p.a{
					correctAnswers++ 
					fmt.Println("You are correct")
				}
		}
			}
		fmt.Printf("You had %d correct out of %d\n",correctAnswers,len(problems))
		}

	}
}

func parseLines(lines [][]string) []problems {
	ret := make([]problems, len(lines))
	for i, v := range lines {
		ret[i] = problems{
			q: v[0],
			a: strings.TrimSpace(v[1]),
		}
	}

	return ret
}
func exit(errMsg string) {
	fmt.Println(errMsg)
	os.Exit(1)
}


