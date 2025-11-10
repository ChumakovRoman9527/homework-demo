package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// **Описание**: Создайте систему управления товарами интернет-магазина с возможностью добавления новых товаров через HTTP POST запрос
//
// **Входные данные**:
// - HTTP POST запрос на маршрут `/products`
// - JSON в теле запроса с полями: name (string), price (float64), category (string)
// - Структура Product с полями: ID uint, Name string, Price float64, Category string, CreatedAt time.Time
// - Репозиторий с методом Create(product *Product) error
//
// **Выходные данные**:
// - HTTP статус 201 (Created) с JSON созданного товара при успехе
// - HTTP статус 400 (Bad Request) при некорректных данных запроса
// - HTTP статус 500 (Internal Server Error) при ошибках базы данных
//
// **Ограничения**:
// - Название товара должно быть непустой строкой (минимум 1 символ)
// - Цена должна быть положительным числом больше 0
// - Категория должна быть непустой строкой
// - Используйте только стандартную библиотеку Go для HTTP операций
// - Обязательная валидация всех входных данных
//
// **Примеры**:
// Input: POST /products {"name": "Laptop Gaming", "price": 1299.99, "category": "Electronics"}
// Output: HTTP 201 {"id": 1, "name": "Laptop Gaming", "price": 1299.99, "category": "Electronics", "created_at": "2024-01-15T10:30:00Z"}
//
// Input: POST /products {"name": "", "price": -100, "category": "Electronics"}
// Output: HTTP 400 Bad Request

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type TaskRepository struct {
	// Ваша реализация репозитория
	items  map[uint]*Task // Хранение объектов по ключу ID
	mu     sync.Mutex     // Для блокировки одновременных изменений
	nextID uint           // Следующий уникальный ID
}

/*
	func NewProductRepository() *ProductRepository {
		return &ProductRepository{}
	}
*/
var ErrNotFound = errors.New("not found")

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		items:  make(map[uint]*Task),
		nextID: 1,
	}
}
func (r *TaskRepository) Create(task *Task) error {
	// Ваша реализация
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.items[task.ID] = task
	r.nextID++

	return nil
}

func (r *TaskRepository) GetTask(id uint) (*Task, error) {
	// Ваша реализация

	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}

	return task, nil
}

func (r *TaskRepository) DeleteTask(id uint) error {
	// Ваша реализация

	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.items[id]
	if !ok {
		return ErrNotFound
	}
	now := time.Now().UTC()
	task.DeletedAt = now
	return nil
}

func (r *TaskRepository) Update(task *Task) error {
	// Ваша реализация
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

type TaskHandler struct {
	repo *TaskRepository
}

func idFromPath(r *http.Request) (uint64, error) {

	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewTaskHandler(repo *TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	task_id, err := idFromPath(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid ID"})
		return
	}

	task, err := h.repo.GetTask(uint(task_id))
	if err == ErrNotFound {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Error"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	task_id, err := idFromPath(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid ID"})
		return
	}
	err = h.repo.DeleteTask(uint(task_id))
	if err == ErrNotFound {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Error"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"OK": "task deleted"})
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req TaskCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON"})
		return
	}
	if len(req.Title) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Название Задачи должно быть непустой строкой (минимум 1 символ)"})
		return
	}

	if !(req.Priority == "high" || req.Priority == "medium" || req.Priority == "low") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "приоритет должна быть high medium low"})
		return
	}

	if len(req.Description) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Описание Задачи должна быть непустой строкой"})
		return
	}

	task_id, err := idFromPath(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid ID"})
		return
	}
	task, err := h.repo.GetTask(uint(task_id))

	if err == ErrNotFound {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Error"})
		return
	}

	task = &Task{
		ID:          uint(task_id),
		Title:       req.Title,
		Priority:    req.Priority,
		Description: req.Description,
		UpdatedAt:   time.Now().UTC(),
	}

	err = h.repo.Update(task)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Error"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Ваш код здесь

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req TaskCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON"})
		return
	}
	if len(req.Title) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Название Задачи должно быть непустой строкой (минимум 1 символ)"})
		return
	}

	if !(req.Priority == "high" || req.Priority == "medium" || req.Priority == "low") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "приоритет должна быть high medium low"})
		return
	}

	if len(req.Description) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Описание Задачи должна быть непустой строкой"})
		return
	}

	task := &Task{
		ID:          0,
		Title:       req.Title,
		Priority:    req.Priority,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = h.repo.Create(task)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Error"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func main() {
	repo := NewTaskRepository()
	handler := NewTaskHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", handler.GetTask)
	mux.HandleFunc("PATCH /tasks/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Ваш код для запуска сервера
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}
}
