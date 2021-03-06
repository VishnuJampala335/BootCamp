package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func parsequiz(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func getinput(input chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	input <- answer
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	quizFilename := flag.String("csv", "problems.csv", "a csv file containing quiz")
	timeLimit := flag.Int("limit", 15, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*quizFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s\n", *quizFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("falied to parse the csv file")
	}

	problems := parsequiz(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemLoop:
	for i, p := range problems {

		fmt.Printf("problem #%d: %s = \n", i+1, p.q)

		answerCh := make(chan string)

		go getinput(answerCh)

		select {

		case <-timer.C:
			fmt.Println()
			break problemLoop

		case answer := <-answerCh:

			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("you got %d questions correct out of %d \n", correct, len(problems))

}
