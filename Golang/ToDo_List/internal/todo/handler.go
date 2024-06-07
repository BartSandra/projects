package todo

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"
    "database/sql"
)

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
    idParam, ok := r.URL.Query()["id"]
    if !ok || len(idParam[0]) < 1 {
        switch r.Method {
        case http.MethodGet:
            h.getTodos(w, r)
        case http.MethodPost:
            h.createTodo(w, r)
        case http.MethodPut:
            h.updateTodo(w, r)
        case http.MethodDelete:
            h.deleteTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    } else {
        id, err := strconv.Atoi(idParam[0])
        if err != nil {
            http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
            return
        }
        switch r.Method {
        case http.MethodGet:
            h.getTodo(w, r, id)
        case http.MethodPut:
            h.updateTodo(w, r)
        case http.MethodDelete:
            h.deleteTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

// getTodo обрабатывает GET
func (h *Handler) getTodos(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    status := r.URL.Query().Get("status")
    dueDate := r.URL.Query().Get("due_date")

    todos, err := h.service.GetTodos(page, limit, status, dueDate)
    if err != nil {
        http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(todos)
}

// getTodo обрабатывает GET запросы для конкретной задачи
func (h *Handler) getTodo(w http.ResponseWriter, _ *http.Request, id int) {
    t, err := h.service.GetTodo(id)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }
    json.NewEncoder(w).Encode(t)
}

// createTodo обрабатывает POST запросы
func (h *Handler) createTodo(w http.ResponseWriter, r *http.Request) {
    var t Todo
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
        return
    }
    if err := validateTodo(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.service.CreateTodo(t.Title, t.Description, t.DueDate, t.Completed); err != nil {
        http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}

// updateTodo обрабатывает PUT запросы
func (h *Handler) updateTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
        return
    }
    var t Todo
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
        return
    }
    if err := validateTodo(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.service.UpdateTodo(id, t.Title, t.Description, t.DueDate, t.Completed); err != nil {
        http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// deleteTodo обрабатывает DELETE запросы
func (h *Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
        return
    }
    err = h.service.DeleteTodo(id)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusOK)
}

func validateTodo(t *Todo) error {
    if t.Title == "" {
        return errors.New("title is required")
    }
    if t.DueDate == "" {
        return errors.New("due date is required")
    }
    return nil
}
