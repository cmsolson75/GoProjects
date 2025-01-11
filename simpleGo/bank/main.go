package main

import (
	"log"
	"net/http"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/repository"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/handler"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/service"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/workflows"
)

// var templates = template.Must(template.ParseGlob(filepath.Join("templates", "*html")))

// func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
// 	err := templates.ExecuteTemplate(w, tmpl+".html", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

type App struct {
	Router *http.ServeMux
}

func NewApp(router *http.ServeMux) *App {
	return &App{
		Router: router,
	}
}

var DB = repository.NewDBContext()

func main() {
	customerRepo := repository.NewInMemoryCustomer(DB)
	accountRepo := repository.NewInMemoryAccount(DB)
	customerService := service.NewCustomerService(customerRepo)
	accountService := service.NewAccountService(accountRepo)
	userWorkflow := workflows.NewUserWorkflow(*customerService, *accountService)

	customerHandler := handler.NewCustomerHandler(*customerService, *userWorkflow)
	accountHandler := handler.NewAccountHandler(*accountService)

	app := NewApp(http.NewServeMux())

	customerService.AddCustomer("cam@email.com", "Cameron")
	customer, _ := customerService.ViewCustomerByEmail("cam@email.com")

	accountService.CreateNewAccount(customer)
	accountService.Deposit(customer, 3000)

	fs := http.FileServer(http.Dir("./static/"))
	app.Router.Handle("GET /static/", http.StripPrefix("/static", fs))

	app.Router.HandleFunc("POST /new-customer", customerHandler.HandlePostNewCustomer)
	app.Router.HandleFunc("POST /login", customerHandler.HandlePostLogin)
	app.Router.HandleFunc("GET /account", accountHandler.HandleGetAccount)
	app.Router.HandleFunc("POST /deposit", accountHandler.HandlePostDeposit)
	app.Router.HandleFunc("POST /withdraw", accountHandler.HandlePostWithdraw)
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
