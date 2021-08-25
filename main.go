package main

import (
	"log"
	"net/http"
)


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
	//renderTemplate(w, "list", todos)
}
