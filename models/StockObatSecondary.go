package models

type StockObatSecondaryModel struct {
	ID           uint    `gorm:"primary_key"`
	Nama     string  `gorm:"column:username"`
	Deskripsi        string  `gorm:"column:email;unique_index"`
	Keterangan          string  `gorm:"column:bio;size:1024"`
	Status        *string `gorm:"column:image"`
	Created_at    string  `gorm:"column:password;not null"`
}
