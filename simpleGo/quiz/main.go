package main

import (
	"log"
	"net/http"

	"github.com/cmsolson75/GoProjects/simpleGo/quiz/handler"
	"github.com/cmsolson75/GoProjects/simpleGo/quiz/quiz"
	"github.com/cmsolson75/GoProjects/simpleGo/quiz/service"
)

func main() {
	jsonQuiz, err := quiz.NewJSONQuiz("questions.json")
	if err != nil {
		log.Println(err)
		return
	}
	quizService := service.NewQuizService(jsonQuiz)
	quizDataHandler := handler.NewQuizDataHandler(quizService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/quiz", quizDataHandler.QuizHandler)
	mux.HandleFunc("/grade", handler.GradeHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
