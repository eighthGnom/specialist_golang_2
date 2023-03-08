package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (a *Article) TableName() string {
	return "article"
}
