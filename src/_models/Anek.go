package models

type Anek struct {
	ID      int    `sql:"id;AUTO_INCREMENT" gorm:"primary_key" json:"id"`
	Сontent string `sql:"content;size:255" json:"content"`
	Rating  int    `sql:"rating;size:255" json:"rating"`
}
