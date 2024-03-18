package models

type Users struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(50)" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"desktipsi"`
}