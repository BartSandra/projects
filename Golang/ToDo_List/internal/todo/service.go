package todo

type Service struct {
    repo *Repository
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

// GetTodos получает задачи с фильтрацией и пагинацией
func (s *Service) GetTodos(page, limit int, status, dueDate string) ([]*Todo, error) {
    if page == 0 {
        page = 1
    }
    if limit == 0 {
        limit = 10
    }
    return s.repo.SelectTodos(page, limit, status, dueDate)
}

// GetTodo получает задачу по ID
func (s *Service) GetTodo(id int) (*Todo, error) {
    return s.repo.SelectTodo(id)
}

// CreateTodo создает новую задачу
func (s *Service) CreateTodo(title, description, dueDate string, completed bool) error {
    return s.repo.InsertTodo(title, description, dueDate, completed)
}

// UpdateTodo обновляет существующую задачу
func (s *Service) UpdateTodo(id int, title, description, dueDate string, completed bool) error {
    return s.repo.UpdateTodo(id, title, description, dueDate, completed)
}

// DeleteTodo удаляет задачу по ID
func (s *Service) DeleteTodo(id int) error {
    return s.repo.DeleteTodo(id)
}
