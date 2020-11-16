package models

import "gorm.io/gorm"

type entryDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//Games is the struc of the table called "games"
type Games struct {
	gorm.Model
	ID            uint      `json:"id"`
	IDOwner       uint      `json:"idOwner"`
	EntryDate     entryDate `json:"entryDate"`
	Disponibility bool      `json:"disponibility"`
	Comments      string    `json:"comments" gorm:"size:200"`
}
