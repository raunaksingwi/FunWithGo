package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type question struct {
	question, answer string
}

func main() {
	var csvFilePath string
	var timeLimit int
	var score int

	flag.StringVar(&csvFilePath, "csvFile", "Resources/problems.csv", "Path to the input csv files that contains the questions for the quiz. (Default: Resources/problems.csv)")
	flag.IntVar(&timeLimit, "timeLimit", 30, "Time bound for answering each question (in seconds)")
	flag.Parse()

	responseCh := make(chan string)
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalln("Cannot open the csv, error: ", err)
	}

	// Read all problems and answers from csv file and store in the quizQuestions array.
	reader := csv.NewReader(csvFile)
	allQuestions, err := reader.ReadAll()
	totalQuestions := len(allQuestions)
	quizQuestions := make([]question, totalQuestions)

	for i, question := range allQuestions {
		quizQuestions[i].question = question[0]
		quizQuestions[i].answer = strings.ToLower(strings.TrimSpace(question[1]))
	}

	//TODO: Traverse the quizQuestions array in random order.
	for _, question := range quizQuestions {

		fmt.Printf("%s? ", question.question)
		expire := time.NewTimer(time.Duration(timeLimit) * time.Second)

		go func() {
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			responseCh <- response
		}()

		select {
		case <-expire.C:
			fmt.Printf("\nTIMEOUT!\n\n")
			continue

		case response := <-responseCh:
			expire.Stop()
			if strings.ToLower(question.answer) == response {
				score++
				fmt.Printf("CORRECT\n\n")
			} else {
				fmt.Printf("INCORRECT\n\n")
			}
			continue
		}

	}

	fmt.Printf("Your score is: %d/%d\n", score, totalQuestions)
	percScore := float32(score) / float32(totalQuestions) * 100
	fmt.Printf("Percentage(%%) Score: %d%%\n\n", int(percScore))

	if percScore > 60 {
		fmt.Println("Congratulations, you passed!!!")
	} else {
		fmt.Println("Better luck next time")
	}
}