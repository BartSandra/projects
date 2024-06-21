package todo

type Todo struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    DueDate     string `json:"due_date"`
    Completed   bool   `json:"completed"`
}