package entities

type Todo struct {
	ID          uint    `json:"id"`
	Title        string  `json:"title"`
	CreatedAt       string `json:"created_at"`
	Description string  `json:"description"`
	Completed bool `json:"completed"`
}