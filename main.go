package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks = []Task{{ID: "1", Title: "Estudar Docker", Done: false}}
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mu.Lock()
		defer mu.Unlock()

		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(tasks)
		case http.MethodPost:
			var t Task
			if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
				tasks = append(tasks, t)
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(t)
			}
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
