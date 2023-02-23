package api

import (
	"net/http"

	"github.com/eighthGnom/standard_web_server/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	router  *mux.Router
	logger  *logrus.Logger
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

}

func (api *API) Start() error {
	api.configureRouterField()
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	if err := api.configureStorageField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
