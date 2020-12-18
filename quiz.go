package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var correctAnswersCounter int
var totalCounter int

func main() {
	//create a reader which holds CSV file

	quizReader := readCsv("problems.csv")
	processQuiz(quizReader)
	fmt.Printf("You've guessed %s out of %s correct", fmt.Sprint(correctAnswersCounter), fmt.Sprint(totalCounter))

}

func readCsv(fileName string) *csv.Reader {
	csvfile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}

	//parse the file
	r := csv.NewReader(csvfile)
	//fmt.Printf("%T", r)
	return r

}

func processQuiz(quiz *csv.Reader) {
	//iterate and read the file

	for {
		question, err := quiz.Read()
		if err == io.EOF {
			break

		}

		if err != nil {
			log.Fatal(err)

		}
		totalCounter++
		fmt.Printf("Question %s: %s ", fmt.Sprint(totalCounter), question[0])

		processUserInput(question[1])
	}

}

func processUserInput(correctAnswer string) {

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(userInput)
	userInput = strings.TrimSpace(userInput)

	if userInput == correctAnswer {
		correctAnswersCounter++
	}

}
