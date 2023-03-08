package main

import (
	"log"

	"github.com/eighthGnom/gin_gorm/models"
	"github.com/eighthGnom/gin_gorm/routers"
	"github.com/eighthGnom/gin_gorm/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	err error
)

func main() {
	storage.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=restapi sslmode=disable")
	if err != nil {
		log.Fatal("error while assessing database", err)
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{})
	router := routers.SetupRouter()
	router.Run()
}
