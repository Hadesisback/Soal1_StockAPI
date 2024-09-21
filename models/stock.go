package models

type Stock struct {
    ID        uint    `json:"id" gorm:"primaryKey"`
    Name      string  `json:"name"`
    Code      string  `json:"code"`
    Price     float64 `json:"price"`
    Frequency int     `json:"frequency"`
    Volume    int     `json:"volume"`
}