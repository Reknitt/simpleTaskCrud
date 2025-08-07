package entities

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}
