package entity

type Item struct {
	Id       string
	Price    float64
	Quantity int16
	Product  Product
}

// type Article struct {
// 	ID        uint       `json:"id" gorm:"primary_key"`
// 	Title     string     `json:"title"`
// 	Content   string     `json:"content"`
// 	Author    string     `json:"author"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
// }
