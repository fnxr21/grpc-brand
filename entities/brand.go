package entities

type Brand struct {
	ID   uint   `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Year uint   `json:"year" gorm:"not null"` // Correct type mapping
}
