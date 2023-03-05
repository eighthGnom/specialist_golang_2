package storage

import (
	"fmt"
	"log"

	"github.com/eighthGnom/standard_web_server/internal/app/models"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	articleTable = "articles"
)

func (a *ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s  (title, author, content) VALUES ($1, $2, $3) RETURNING article_id", articleTable)
	err := a.storage.db.QueryRow(query, article.Title, article.Author, article.Content).Scan(&article.ID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", articleTable)
	rows, err := a.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	articles := make([]*models.Article, 0)
	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ID, &article.Title, &article.Author, &article.Content); err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *ArticleRepository) FindByID(requiredID int) (*models.Article, bool, error) {
	var isFound bool
	articles, err := a.SelectAll()
	if err != nil {
		return nil, isFound, err
	}
	var foundArticle *models.Article
	for _, article := range articles {
		if article.ID == requiredID {
			isFound = true
			foundArticle = article
			break
		}
	}
	return foundArticle, isFound, nil
}

func (a *ArticleRepository) DeleteByID(requiredID int) (*models.Article, error) {
	article, ok, err := a.FindByID(requiredID)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE article_id = $1", articleTable)
		_, err := a.storage.db.Exec(query, requiredID)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

func (a *ArticleRepository) UpdateByID(id int, newArticle *models.Article) (*models.Article, error) {
	oldArticle, ok, err := a.FindByID(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("UPDATE articles SET title = $1, author = $2, content = $3 WHERE article_id = $4")
		_, err := a.storage.db.Exec(query, newArticle.Title, newArticle.Author, newArticle.Content, id)
		if err != nil {
			return nil, err
		}
	}

	return oldArticle, nil
}
