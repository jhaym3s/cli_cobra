package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problems struct {
	q string
	a string
}

func main() {
	quizFileName := flag.String("quiz", "quiz.cvs", "")
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
			correctAnswers := 0
			for i, p := range problems {
				fmt.Printf("Question %d %s = \n", i+1, p.q)
				var answer string
				fmt.Scanf("%s\n", &answer)
				if answer == p.a{
					correctAnswers++ 
					fmt.Println("You are correct")
				}else{
					fmt.Println("Mumu")
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
