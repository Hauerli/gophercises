package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// import file
	file, err := os.Open("problems.csv")
	// if file not found - error
	if err != nil {
		log.Fatal(err)
	}

	// create new reader buffer to import file recods
	r := csv.NewReader(bufio.NewReader(file))
	// go over records
	for {
		record, err := r.Read()
		// break if EndOfFile is reached
		if err == io.EOF {
			break
		}
		// error is something is wrong
		if err != nil {
			log.Fatal(err)
		}

		question := record[0]
		answer := record[1]

		fmt.Println(question, " ", answer)

	}
}
