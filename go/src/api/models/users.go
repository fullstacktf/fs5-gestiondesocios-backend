package models

//AssocUser is the struc of the table called "users"
type AssocUser struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username"`
	UserPassword string `json:"user_password"`
}
