package models

type Food struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	MakerID    int     `json:"maker_id"`
	CategoryID int     `json:"category_id"`
	Price      float64 `json:"price"`
}
