package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var templates  *template.Template = template.Must(template.ParseFiles("templates/list.html"))
var validPath = regexp.MustCompile("^/(list)/([a-zA-Z0-9]+)$")

type Todo struct {
	ID			string  `json:"id"`
	IsFinished 	bool	`json:"is_finished"`
	Title		string	`json:"title"`
	Body		string 	`json:"body"`
	DueDate		time.Time	`json:"due_date"`
}

var layout string = "2006-01-02T15:04:05.000Z"

var todos []Todo = []Todo{
	{ID: "0", IsFinished: false, Title: "Todo 1", Body: "Todo body todo body", DueDate: time.Now()},
	{ID: "1", IsFinished: false, Title: "Todo 2", Body: "Todo body todo body", DueDate: time.Now()},
	{ID: "2", IsFinished: true, Title: "Todo 3", Body: "Todo body todo body", DueDate: time.Now()},
	{ID: "3", IsFinished: false, Title: "Todo 4", Body: "Todo body todo body", DueDate: time.Now()},
}

func main() {
	http.HandleFunc("/todos", handleListing)
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("/delete", handleDelete)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleListing(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "list.html", todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	newTodoTitle := r.FormValue("todo-title")
	newTodoBody := r.FormValue("todo-body")

	if newTodoTitle == "" {
		http.Error(w, "Todo Title cannot be empty", http.StatusBadRequest)
		return
	}

	newTodo := Todo{ID: strconv.Itoa(len(todos) + 1), IsFinished: false, Title: newTodoTitle, Body: newTodoBody, DueDate: time.Now()}
	todos = append(todos, newTodo)
	http.Redirect(w, r, "/todos", http.StatusFound)
}

func removeTodo(todoId string) []Todo {
	for i, todo := range todos {
		if todo.ID == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			return todos
		}
	}

	return todos
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == ""{
		http.Error(w, "Id should be given and should be a number", http.StatusBadRequest)
		return
	}

	todos = removeTodo(id)
	http.Redirect(w, r, "/todos", http.StatusFound)
}

