package models

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" validate:"required,min=3"` // Title is required, min length 3
	Author string `json:"author" validate:"required"`      // Author is required
	ISBN   string `json:"isbn" validate:"required,len=13"` // ISBN must be exactly 13 characters
}
