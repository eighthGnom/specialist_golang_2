package models

type Article struct {
	ID      int    `json:"article_id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
