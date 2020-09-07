package main

import "testing"

// Run benchmark on getQuizeQuestionsFromCsv function
func BenchmarkGetQuizQuestionsFromCsv(b *testing.B) {

	csvFilePath := "./Resources/problems.csv"
	for n := 0; n < b.N; n++ {
		getQuizQuestionsFromCsv(csvFilePath)
	}
}
