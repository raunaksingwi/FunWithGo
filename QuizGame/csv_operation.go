package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type question struct {
	question, answer string
}

func getQuizQuestionsFromCsv(csvFilePath string) (quizQuestions []question, totalQuestions int) {

	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalln("Cannot open the csv file, error: ", err)
	}

	// Read all problems and answers from csv file and store in the quizQuestions array.
	reader := csv.NewReader(csvFile)
	allQuestions, err := reader.ReadAll()
	totalQuestions = len(allQuestions)
	quizQuestions = make([]question, totalQuestions)

	for i, question := range allQuestions {
		quizQuestions[i].question = question[0]
		quizQuestions[i].answer = strings.ToLower(strings.TrimSpace(question[1]))
	}

	return
}
