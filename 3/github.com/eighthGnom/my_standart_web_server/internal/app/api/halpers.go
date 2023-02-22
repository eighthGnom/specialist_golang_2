package api

import (
	"my_standart_web_server/storage"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (api API) configureRouterField() {
	api.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello user"))
	})
}

func (api API) configureLoggerField() error {
	level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(level)
	return nil
}

func (api API) configureStorageField() error {
	stor := storage.New(api.config.StorageConfig)
	err := stor.Open()
	if err != nil {
		return err
	}
	api.storage = stor
	api.logger.Info("Connected to database successfully")
	return nil
}
