package models

//TODO Search info about FK on this typo of structs

//BorrowedGame is the struc of the table called "borrowedGames"
type BorrowedGame struct {
	IDGame     uint   `json:"id_game" gorm:"primaryKey"`
	IDBorrower uint   `json:"id_borrower"`
	BorrowDate string `json:"borrow_date"`
}
