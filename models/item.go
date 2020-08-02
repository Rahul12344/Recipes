package models

//Item crowdsourced item
type Item struct {
	ItemID   string  `json:"item_id" gorm:"not null;unique"`
	ItemName string  `json:"itemname" gorm:"not null"`
	Source   string  `json:"source" gorm:"not null"`
	Store    string  `json:"store" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Lat      float64 `json:"lat" gorm:"not null"`
	Lon      float64 `json:"lon" gorm:"not null"`
}

//TableName name of table
func (Item) TableName() string {
	return "items"
}
