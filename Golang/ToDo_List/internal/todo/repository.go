package todo

import (
    "database/sql"
    "log"
)

// Repository структура для репозитория
type Repository struct {
    db *sql.DB
}

// NewRepository создает новый экземпляр Repository
func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

// SelectTodos получает задачи с возможностью фильтрации и пагинации
func (r *Repository) SelectTodos(page, limit int, status, dueDate string) ([]*Todo, error) {
    query := "SELECT id, title, description, due_date, completed FROM todos WHERE 1=1"
    var params []interface{}

    if status != "" {
        query += " AND completed = $1"
        params = append(params, status == "true")
    }

    if dueDate != "" {
        query += " AND due_date = $2"
        params = append(params, dueDate)
    }

    query += " ORDER BY id LIMIT $3 OFFSET $4"
    params = append(params, limit, (page-1)*limit)

    rows, err := r.db.Query(query, params...)
    if err != nil {
        log.Printf("Error selecting todos: %v", err)
        return nil, err
    }
    defer rows.Close()

    var todos []*Todo
    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed); err != nil {
            log.Printf("Error scanning todo: %v", err)
            return nil, err
        }
        todos = append(todos, &t)
    }
    return todos, nil
}

// SelectTodo получает задачу по ID
func (r *Repository) SelectTodo(id int) (*Todo, error) {
    query := "SELECT id, title, description, due_date, completed FROM todos WHERE id = $1"
    row := r.db.QueryRow(query, id)

    var t Todo
    if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed); err != nil {
        return nil, err
    }
    return &t, nil
}

// InsertTodo добавляет новую задачу
func (r *Repository) InsertTodo(title, description, dueDate string, completed bool) error {
    query := "INSERT INTO todos (title, description, due_date, completed) VALUES ($1, $2, $3, $4)"
    _, err := r.db.Exec(query, title, description, dueDate, completed)
    if err != nil {
        log.Printf("Error inserting todo: %v", err)
        return err
    }
    return nil
}

// UpdateTodo обновляет существующую задачу
func (r *Repository) UpdateTodo(id int, title, description, dueDate string, completed bool) error {
    query := "UPDATE todos SET title = $1, description = $2, due_date = $3, completed = $4 WHERE id = $5"
    _, err := r.db.Exec(query, title, description, dueDate, completed, id)
    if err != nil {
        log.Printf("Error updating todo: %v", err)
        return err
    }
    return nil
}

// DeleteTodo удаляет задачу по ID
func (r *Repository) DeleteTodo(id int) error {
    query := "DELETE FROM todos WHERE id = $1"
    _, err := r.db.Exec(query, id)
    if err != nil {
        log.Printf("Error deleting todo: %v", err)
        return err
    }
    return nil
}

