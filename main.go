package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates  *template.Template = template.Must(template.ParseFiles("templates/list.html"))
var validPath = regexp.MustCompile("^/(list)/([a-zA-Z0-9]+)$")

type Todo struct {
	ID			string  `json:"id"`
	IsFinished 	bool	`json:"is_finished"`
	Title		string	`json:"title"`
	DueDate		uint64	`json:"due_date"`
}

var todos []Todo = []Todo{
	{ID: "1", IsFinished: false, Title: "Todo 1", DueDate: 1629927916662},
	{ID: "2", IsFinished: false, Title: "Todo 2", DueDate: 1629927916962},
	{ID: "3", IsFinished: true, Title: "Todo 3", DueDate: 1629927916692},
	{ID: "4", IsFinished: false, Title: "Todo 4", DueDate: 1629927916940},
}

func main() {
	http.HandleFunc("/todos", handleListing)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleListing(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "list.html", todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


