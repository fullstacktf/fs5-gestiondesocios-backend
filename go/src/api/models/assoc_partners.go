package models

//AssocPartner is the struc of the table called "games"
type AssocPartner struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	PartnerName string `json:"partner_name" gorm:"size:30"`
}
