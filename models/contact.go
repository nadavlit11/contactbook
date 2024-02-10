package models

type Contact struct {
	ID        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Phone     string `json:"phone" db:"phone"`
	Address   string `json:"address" db:"address"`
}
