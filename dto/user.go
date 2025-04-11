package dto

type User struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserWithTasks struct {
	User  *User  `json:"user"`
	Tasks []Task `json:"tasks"`
}
