package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Task represents a to-do task.
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // e.g., "pending" or "completed"
}

var tasks = []Task{} // In-memory storage for tasks
const Dport = ":8012"

var guide ="Welcome to task server"

func main() {
	http.HandleFunc("/", returnInfo)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/task/", taskHandler)
	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, nil)
}


func returnInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "This is from jevica ")
	fmt.Fprintln(w, "My Student Id is :500218849")
	fmt.Fprintln(w, "My github is : https://github.com/jevdsouza")
	fmt.Fprintln(w,"This is a tutorial on the API. It performs basic CRUD functions.")
	fmt.Fprintln(w,"GET /tasks shows a list of all tasks.")
	fmt.Fprintln(w,"POST /tasks creates a new task.")
	fmt.Fprintln(w,"PUT /task/{id} updates an existing task.")
	fmt.Fprintln(w,"DELETE /task/{id}: Deletes the task identified by {id}.")
	}
	
// Handle requests to the /tasks endpoint
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task.ID = uuid.New().String() // Generate a unique ID for the task
		tasks = append(tasks, task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handle requests to the /task/{id} endpoint
func taskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL path
	taskID := strings.TrimPrefix(r.URL.Path, "/task/")

	switch r.Method {
	case "PUT":
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		found := false
		for i, task := range tasks {
			if task.ID == taskID {
				updatedTask.ID = task.ID // Ensuring the ID remains unchanged
				tasks[i] = updatedTask
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedTask)
	case "DELETE":
		index := -1
		for i, task := range tasks {
			if task.ID == taskID {
				index = i
				break
			}
		}
		if index != -1 {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Task not found", http.StatusNotFound)
		}
	}
}
