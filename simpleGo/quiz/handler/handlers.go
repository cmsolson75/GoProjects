package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/cmsolson75/GoProjects/simpleGo/quiz/quiz"
	"github.com/cmsolson75/GoProjects/simpleGo/quiz/service"
)

var templates = template.Must(template.ParseGlob(filepath.Join("templates", "*html")))

type ApplicationState struct {
	Questions []quiz.QuizEntry
}

func NewApplicationState(questions []quiz.QuizEntry) *ApplicationState {
	return &ApplicationState{Questions: questions}
}

// the Any feels weird on data, but it was the only way to get this to work.
func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type QuizDataHandler struct {
	data *service.QuizService
}

func NewQuizDataHandler(quizService *service.QuizService) *QuizDataHandler {
	return &QuizDataHandler{data: quizService}
}

func (q *QuizDataHandler) QuizHandler(w http.ResponseWriter, r *http.Request) {
	// abstract this DRY
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}
	numberQuestions, err := strconv.Atoi(r.FormValue("qnum"))
	if err != nil {
		http.Error(w, "error converting form value to int", http.StatusInternalServerError)
		return
	}

	questionSubset, err := q.data.RandomSubset(numberQuestions)
	if err != nil {
		log.Println(err)
		return
	}
	data := NewApplicationState(questionSubset)
	renderTemplate(w, "quiz", data)

}

func scoreFormat(correct int, totalQuestions int) string {
	percent := (float64(correct) / float64(totalQuestions)) * 100.0
	return fmt.Sprintf("(%d/%d) %.0f%%\n", correct, totalQuestions, percent)
}

func GradeHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}
	correct := 0
	totalQuestions := len(r.Form) / 2

	for i := 0; i < totalQuestions; i++ {
		userAnswer := r.FormValue(fmt.Sprintf("%d", i))
		correctAnswer := r.FormValue(fmt.Sprintf("%d_answer", i))
		if userAnswer == correctAnswer {
			correct++
		}
	}
	scoreDisplay := scoreFormat(correct, totalQuestions)
	data := struct {
		Score string
	}{
		Score: scoreDisplay,
	}

	renderTemplate(w, "score", data)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "root", nil)
}
