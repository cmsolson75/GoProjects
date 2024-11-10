package quiz

import (
	"encoding/json"
	"log"
	"os"
)

type QuizEntry struct {
	Question string
	Answers  []string
	Correct  string
}

type QuizRepository interface {
	ViewQuestions() []QuizEntry
}

type JSONQuiz struct {
	Questions *[]QuizEntry
}

func NewJSONQuiz(jsonFile string) (*JSONQuiz, error) {
	questions := []QuizEntry{}

	data, err := os.Open(jsonFile)
	if err != nil {
		log.Println("Error opening json file:", err)
		return &JSONQuiz{}, err
	}

	defer data.Close()

	err = json.NewDecoder(data).Decode(&questions)
	if err != nil {
		log.Println("Error decoding json file:", err)
		return &JSONQuiz{}, err
	}
	return &JSONQuiz{Questions: &questions}, nil
}

func (j *JSONQuiz) ViewQuestions() []QuizEntry {
	return *j.Questions
}
