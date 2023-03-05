package api

import (
	"net/http"

	"github.com/eighthGnom/standard_web_server/internal/app/middelware"
	"github.com/eighthGnom/standard_web_server/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v1"
)

func (api *API) configureRouterField() {
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	//api.router.HandleFunc(prefix+"/articles/{id}", api.GetArticleByID).Methods("GET")
	api.router.Handle(prefix+"/articles/{id}", middelware.JwtMiddleware.Handler(
		http.HandlerFunc(api.GetArticleByID)),
	).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleByID).Methods("DELETE")
	api.router.HandleFunc(prefix+"/articles", api.CreateArticles).Methods("POST")
	api.router.HandleFunc(prefix+"/user/register", api.PostUserRegister).Methods("POST")
	api.router.HandleFunc(prefix+"/articles/{id}", api.UpdateArticleByID).Methods("PUT")
	api.router.HandleFunc(prefix+"/user/auth", api.PostToAuth).Methods("POST")
}

func (api *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(log_level)
	return nil
}

func (api *API) configureStorageField() error {
	api.storage = storage.New(api.config.StorageConfig)
	if err := api.storage.Open(); err != nil {
		return err
	}
	api.logger.Info("Connected to database successfully")
	return nil
}
