package structs

type User struct {
	ID      uint   `json:"_id" `
	IsActive bool `json:"isActive" `
	Name    string   `json:"name" `
	Gender  string   `json:"gender"`
}

type Users struct{
	Usrs []User
}