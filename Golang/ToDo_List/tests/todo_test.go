package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "todo-app/internal/todo"
    "database/sql"
    "regexp"

    "github.com/stretchr/testify/assert"
    "github.com/DATA-DOG/go-sqlmock"
)

func setupTestDB() (*sql.DB, sqlmock.Sqlmock, error) {
    db, mock, err := sqlmock.New()
    if err != nil {
        return nil, nil, err
    }
    return db, mock, nil
}

func setupTestHandler(db *sql.DB) *todo.Handler {
    repo := todo.NewRepository(db)
    service := todo.NewService(repo)
    return todo.NewHandler(service)
}

func TestTodoAPI(t *testing.T) {
    db, mock, err := setupTestDB()
    assert.NoError(t, err)
    defer db.Close()

    handler := setupTestHandler(db)

    t.Run("GET /todos", func(t *testing.T) {
        rows := sqlmock.NewRows([]string{"id", "title", "description", "due_date", "completed"}).
            AddRow(1, "Test Todo", "Test Description", "2024-01-01", false)
        mock.ExpectQuery("SELECT id, title, description, due_date, completed FROM todos").
            WillReturnRows(rows)

        req, err := http.NewRequest("GET", "/todos", nil)
        assert.NoError(t, err)

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusOK, rr.Code)
        var todos []todo.Todo
        err = json.NewDecoder(rr.Body).Decode(&todos)
        assert.NoError(t, err)
        assert.Len(t, todos, 1)
    })

    t.Run("POST /todos", func(t *testing.T) {
        newTodo := todo.Todo{
            Title:       "New Todo",
            Description: "New Description",
            DueDate:     "2024-01-01",
            Completed:   false,
        }
        body, err := json.Marshal(newTodo)
        assert.NoError(t, err)

        mock.ExpectExec("INSERT INTO todos").
            WithArgs(newTodo.Title, newTodo.Description, newTodo.DueDate, newTodo.Completed).
            WillReturnResult(sqlmock.NewResult(1, 1))

        req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
        assert.NoError(t, err)
        req.Header.Set("Content-Type", "application/json")

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusCreated, rr.Code)
    })

    t.Run("PUT /todos", func(t *testing.T) {
        updatedTodo := todo.Todo{
            Title:       "Updated Todo",
            Description: "Updated Description",
            DueDate:     "2024-01-02",
            Completed:   true,
        }
        body, err := json.Marshal(updatedTodo)
        assert.NoError(t, err)

        mock.ExpectExec("UPDATE todos").
            WithArgs(updatedTodo.Title, updatedTodo.Description, updatedTodo.DueDate, updatedTodo.Completed, 1).
            WillReturnResult(sqlmock.NewResult(1, 1))

        req, err := http.NewRequest("PUT", "/todos?id=1", bytes.NewBuffer(body))
        assert.NoError(t, err)
        req.Header.Set("Content-Type", "application/json")

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusOK, rr.Code)
    })

    t.Run("DELETE /todos", func(t *testing.T) {
        mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todos WHERE id = $1")).
            WithArgs(1).
            WillReturnResult(sqlmock.NewResult(1, 1))

        req, err := http.NewRequest("DELETE", "/todos?id=1", nil)
        assert.NoError(t, err)

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusOK, rr.Code)
    })

    t.Run("GET /todos/{id}", func(t *testing.T) {
        row := sqlmock.NewRows([]string{"id", "title", "description", "due_date", "completed"}).
            AddRow(1, "Test Todo", "Test Description", "2024-01-01", false)
        mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, due_date, completed FROM todos WHERE id = $1")).
            WithArgs(1).
            WillReturnRows(row)
    
        req, err := http.NewRequest("GET", "/todos?id=1", nil)
        assert.NoError(t, err)
    
        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)
    
        assert.Equal(t, http.StatusOK, rr.Code)
        var todo todo.Todo
        err = json.NewDecoder(rr.Body).Decode(&todo)
        assert.NoError(t, err)
        assert.Equal(t, "Test Todo", todo.Title)
    })    

    t.Run("POST /todos with invalid data", func(t *testing.T) {
        invalidTodo := map[string]interface{}{
            "title":       "New Todo",
            "description": "New Description",
            "due_date":    12345,
            "completed":   false,
        }
        body, err := json.Marshal(invalidTodo)
        assert.NoError(t, err)

        req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
        assert.NoError(t, err)
        req.Header.Set("Content-Type", "application/json")

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusBadRequest, rr.Code)
    })

    t.Run("PUT /todos with invalid data", func(t *testing.T) {
        invalidTodo := map[string]interface{}{
            "title":       "Updated Todo",
            "description": "Updated Description",
            "due_date":    12345,
            "completed":   true,
        }
        body, err := json.Marshal(invalidTodo)
        assert.NoError(t, err)

        req, err := http.NewRequest("PUT", "/todos?id=1", bytes.NewBuffer(body))
        assert.NoError(t, err)
        req.Header.Set("Content-Type", "application/json")

        rr := httptest.NewRecorder()
        http.HandlerFunc(handler.HandleTodos).ServeHTTP(rr, req)

        assert.Equal(t, http.StatusBadRequest, rr.Code)
    })
}
