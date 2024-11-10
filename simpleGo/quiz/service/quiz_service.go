package service

import (
	"errors"
	"math/rand"

	"github.com/cmsolson75/GoProjects/simpleGo/quiz/quiz"
)

type QuizService struct {
	Quiz quiz.QuizRepository
}

func NewQuizService(quizRepo quiz.QuizRepository) *QuizService {
	return &QuizService{Quiz: quizRepo}
}

func (q *QuizService) RandomSubset(n int) ([]quiz.QuizEntry, error) {
	questions := q.Quiz.ViewQuestions()
	numberOfQuestions := len(questions)

	if n > numberOfQuestions {
		return []quiz.QuizEntry{}, errors.New("n greater than number of questions")
	}

	var subset []quiz.QuizEntry

	for i := 0; i < n; i++ {
		idx := rand.Intn(numberOfQuestions)
		subset = append(subset, questions[idx])

		questions = append(questions[:idx], questions[idx+1:]...)

		numberOfQuestions--
	}

	return subset, nil
}
