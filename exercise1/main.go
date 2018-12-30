package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// TrueAnswer - Answers during Quiz
var TrueAnswer int

// FalseAnswer - Answers during Quiz
var FalseAnswer int

// FalseAnswer - Answers not answered in Time
var questLeft int

// Print the Question
func printquest(qu string) {
	fmt.Println("What is ", qu, " ?")
}

// Create the Scanner for Terminal input and return input
func importAnswer() string {
	// create new scanner for terminal input
	scanner := bufio.NewScanner(os.Stdin)
	// catch error for scanner
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	// check the scan
	scanner.Scan()
	panswer := scanner.Text()
	return panswer

}

// Compare answers
func compareAnswer(qanswer string, panswer string) {
	if panswer == qanswer {
		TrueAnswer++
	} else {
		FalseAnswer++
	}
}

// Return count of remaining questions
func leftQuestions(r [][]string, err error) int {

	// count left entrys
	return len(r)

}

// Print the results
func finalstats() {
	fmt.Println("-------------------------------------")
	fmt.Println("You did answer ", TrueAnswer+FalseAnswer, " of a total of", TrueAnswer+FalseAnswer+questLeft, "Questions")
	fmt.Println("Correct answers: ", TrueAnswer)
	fmt.Println("Wrong answers: ", FalseAnswer)
}

func main() {

	// creating flag for filepath
	fpath := flag.String("path", "problems.csv", "Path of File")
	// create timerflag
	timer := flag.Int("timer", 30, "Timer")
	// Parse entered flags
	flag.Parse()

	// create timer
	timeout := time.After(time.Duration(*timer) * time.Second)

	// import file
	file, err := os.Open(*fpath)
	// if file not found - error
	if err != nil {
		log.Fatal(err)
	}

	// create new reader buffer to import file records
	reader := csv.NewReader(bufio.NewReader(file))

	// Wait for Enter to start the game
	fmt.Println("Press Enter to start the game you got", *timer, "Seconds.")
	fmt.Scanln()

	// go over records
loop:
	for {
		// select for timeout channel
		select {
		// if timer is 0
		case <-timeout:
			// call func give slice to count remaining questions
			questLeft = leftQuestions(reader.ReadAll())
			fmt.Println("Sorry you did not finish in time !")
			finalstats()
			break loop
		// do as long as timer is not 0
		default:
			record, err := reader.Read()
			// break if EndOfFile is reached
			if err == io.EOF {
				// call func give slice to count remaining questions
				questLeft = leftQuestions(reader.ReadAll())

				// call finalstats to tell correct and wrong answers
				finalstats()
				break loop
			}
			// error is something is wrong
			if err != nil {
				log.Fatal(err)
			}

			// assign question and answer from csv into variables
			question := record[0]
			qanswer := record[1]

			// call question printer
			printquest(question)
			// import person answer
			panswer := importAnswer()

			// compare answers
			compareAnswer(qanswer, panswer)
		}
	}
}
