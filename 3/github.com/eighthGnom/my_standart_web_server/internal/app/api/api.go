package api

import (
	"my_standart_web_server/storage"
	"net/http"

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
		logger: logrus.New()}
}

func (api API) Start() error {
	api.configureRouterField()
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	err := api.configureStorageField()
	if err != nil {
		return err
	}

	api.logger.Info("Starting server at port ", api.config.BindAddr)
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
