package dto

type Task struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	Description string `json:"description"`
}
