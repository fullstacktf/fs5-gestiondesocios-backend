package models

//TODO Search info about FK on this typo of structs

//BorrowedGame is the struc of the table called "borrowedGames"
type BorrowedGame struct {
	IDGame     uint   `json:"idGame" gorm:"primaryKey"`
	IDBorrower uint   `json:"idBorrower"`
	BorrowDate string `json:"borrowDate" gorm:"-"`
}
