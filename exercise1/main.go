package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// TrueAnswer - Answers during Quiz
var TrueAnswer int

// FalseAnswer - Answers during Quiz
var FalseAnswer int

// Print the results
func finalstats() {
	fmt.Println("-------------------------------------")
	fmt.Println("You finished the Quiz:")
	fmt.Println("Correct answers: ", TrueAnswer)
	fmt.Println("Wrong answers: ", FalseAnswer)
}

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

// compare answers
func compareAnswer(qanswer string, panswer string) {
	if panswer == qanswer {
		TrueAnswer++
	} else {
		FalseAnswer++
	}
}

func main() {

	// creating flag for filepath
	var fpath = flag.String("path", "problems.csv", "Path of File")
	flag.Parse()

	// import file
	file, err := os.Open(*fpath)
	// if file not found - error
	if err != nil {
		log.Fatal(err)
	}

	// create new reader buffer to import file recods
	reader := csv.NewReader(bufio.NewReader(file))
	// go over records

	for {
		record, err := reader.Read()
		// break if EndOfFile is reached
		if err == io.EOF {
			// call finalstats to tell correct and wrong answers
			finalstats()
			break
		}
		// error is something is wrong
		if err != nil {
			log.Fatal(err)
		}

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
