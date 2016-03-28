package models

type Anek struct {
	ID      int    `sql:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Сontent string `sql:"content" gorm:"not null;type:text" json:"content"`
	Rating  int    `sql:"rating" gorm:"not null;default:0" json:"rating"`
}
