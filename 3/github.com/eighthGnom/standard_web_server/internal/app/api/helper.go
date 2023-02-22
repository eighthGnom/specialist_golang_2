package api

import (
	"net/http"

	"github.com/eighthGnom/standard_web_server/storage"
	"github.com/sirupsen/logrus"
)

func (api *API) configureRouterField() {
	api.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, username"))
	})
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
