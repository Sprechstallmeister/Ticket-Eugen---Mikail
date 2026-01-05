package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID    int
	Name  string
	Email string
}

// DataStorage - интерфейс для хранения данных
type DataStorage interface {
	Save(user User) string
	Load(data string) User
}

// JSONStorage - хранилище в формате JSON
type JSONStorage struct{}

func (j *JSONStorage) Save(user User) string {
	data, _ := json.Marshal(user)
	result := string(data)
	fmt.Println("[JSON SAVE]", result)
	return result
}

func (j *JSONStorage) Load(data string) User {
	var user User
	json.Unmarshal([]byte(data), &user)
	fmt.Printf("[JSON LOAD] Loaded: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
	return user
}

// XMLStorage - хранилище в формате XML
type XMLStorage struct{}

func (x *XMLStorage) Save(user User) string {
	result := fmt.Sprintf("<User><ID>%d</ID><Name>%s</Name><Email>%s</Email></User>", 
		user.ID, user.Name, user.Email)
	fmt.Println("[XML SAVE]", result)
	return result
}

func (x *XMLStorage) Load(data string) User {
	// Упрощённый парсинг XML (для демонстрации)
	var user User
	fmt.Sscanf(data, "<User><ID>%d</ID><Name>%s</Name><Email>%s</Email></User>", 
		&user.ID, &user.Name, &user.Email)
	fmt.Printf("[XML LOAD] Loaded: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
	return user
}

// CSVAdapter - адаптер для CSV формата
type CSVAdapter struct {
	storage DataStorage
}

func (c *CSVAdapter) Save(user User) string {
	result := fmt.Sprintf("%d,%s,%s", user.ID, user.Name, user.Email)
	fmt.Println("[CSV SAVE]", result)
	return result
}

func (c *CSVAdapter) Load(data string) User {
	parts := strings.Split(data, ",")
	var user User
	fmt.Sscanf(parts[0], "%d", &user.ID)
	user.Name = parts[1]
	user.Email = parts[2]
	fmt.Printf("[CSV LOAD] Loaded: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
	return user
}

func main() {
	fmt.Println("=== Data Storage Adapter Pattern Demo ===\n")

	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	fmt.Printf("Original User: ID=%d, Name=%s, Email=%s\n\n", user.ID, user.Name, user.Email)

	// JSON формат
	fmt.Println("--- JSON Format ---")
	jsonStorage := &JSONStorage{}
	jsonData := jsonStorage.Save(user)
	jsonStorage.Load(jsonData)
	fmt.Println()

	// XML формат
	fmt.Println("--- XML Format ---")
	xmlStorage := &XMLStorage{}
	xmlData := xmlStorage.Save(user)
	xmlStorage.Load(xmlData)
	fmt.Println()

	// CSV формат через адаптер
	fmt.Println("--- CSV Format (via Adapter) ---")
	csvAdapter := &CSVAdapter{storage: jsonStorage}
	csvData := csvAdapter.Save(user)
	csvAdapter.Load(csvData)
}
