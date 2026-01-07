package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string  `valid:"length(3|100)~Title must long 3-100"`
	Price float64 `valid:"rage(50|5000)~Price must rage 50-5000"`
	Code  string  `valid:"matches([B][k]\\d{6}$)~Code is start with BK and affter with six number eg.BK123456"`
}
