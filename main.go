package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// set flag csv and set its default value and description and assign it to csvFilename
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")
	// parse flag
	flag.Parse()

	//try opening the file
	file, err := os.Open(*csvFilename)

	//if theres an error while opening  the file display it along with the file name
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// read from csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	// if error the show this error message
	if err != nil {
		exit("Failed to print the specified csv file.")
	}

	//if no error print contents of file
	problems := parseLines(lines)
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string

		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

// function to print error message by providing the message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
