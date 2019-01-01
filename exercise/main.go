package main

// need to add timer, timer should stop game even if not asking any question

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func printResult(cAnswer int, wAnswer int, sum int) {
	fmt.Println("You did answer", cAnswer+wAnswer, " out of ", sum, " questions.")
	fmt.Println("Correct answers: ", cAnswer)
	fmt.Println("Wrong answers: ", wAnswer)
}

func main() {

	cAnswer := 0
	wAnswer := 0

	flagpath := flag.String("path", "problem.csv", "path defines the path for the problems.csv")
	flagtimer := flag.Int("timer", 30, "timer defines time to finish the questions")
	flag.Parse()

	csvfile, err := os.Open(*flagpath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvfile))
	csvlines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	questsum := len(csvlines)

	for _, line := range csvlines {

		question := line[0]
		answer := line[1]

		fmt.Println("What is ", question, " ?")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == answer {
			cAnswer++
		} else {
			wAnswer++
		}
	}
	printResult(cAnswer, wAnswer, questsum)

}
