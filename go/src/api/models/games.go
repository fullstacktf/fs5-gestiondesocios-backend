package models

//TODO search info about using dates on structs

type date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//Game is the struc of the table called "games"
type Game struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	IDOwner       uint   `json:"idOwner"`
	EntryDate     string `json:"entryDate" gorm:"-"`
	Disponibility bool   `json:"disponibility"`
	Comments      string `json:"comments" gorm:"size:200"`
}
