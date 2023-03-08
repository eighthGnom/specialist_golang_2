package models

import "github.com/eighthGnom/gin_gorm/storage"

func GetAllArticles(article *[]Article) error {
	return storage.DB.Find(article).Error
}

func AddNewArticle(article *Article) error {
	return storage.DB.Create(article).Error
}

func GetArticleByID(article *Article, id string) error {
	return storage.DB.Where("id = ?", id).First(article).Error
}

func DeleteArticleByID(article *Article, id string) error {
	return storage.DB.Where("id = ?", id).Delete(article).Error
}

func UpdateArticleByID(article *Article, id string) error {
	return storage.DB.Update(article).Error
}
