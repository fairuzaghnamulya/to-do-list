package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// Todo represents a todo item
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Time      string `json:"time"`
}

var (
	todos    []Todo
	mu       sync.RWMutex
	dataFile = "./database/todos.json"
)

// InitDB initializes the database
func InitDB() error {
	// Create database directory if not exists
	if err := os.MkdirAll("./database", 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	// Load existing data
	if err := loadFromFile(); err != nil {
		log.Printf("No existing data found, starting fresh: %v", err)
		todos = []Todo{}
	}

	log.Printf("Database initialized with %d todos", len(todos))
	return nil
}

// loadFromFile loads todos from JSON file
func loadFromFile() error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &todos)
}

// saveToFile saves todos to JSON file
func saveToFile() error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// InsertTodo adds a new todo
func InsertTodo(id, title, time string) error {
	mu.Lock()
	defer mu.Unlock()

	todo := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
		Time:      time,
	}
	todos = append([]Todo{todo}, todos...) // Add to beginning
	return saveToFile()
}

// GetAllTodos returns all todos
func GetAllTodos() ([]Todo, error) {
	mu.RLock()
	defer mu.RUnlock()
	return todos, nil
}

// ToggleTodo toggles the completed status
func ToggleTodo(id string) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			return saveToFile()
		}
	}
	return fmt.Errorf("todo not found: %s", id)
}

// DeleteTodo removes a todo by ID
func DeleteTodo(id string) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return saveToFile()
		}
	}
	return fmt.Errorf("todo not found: %s", id)
}
