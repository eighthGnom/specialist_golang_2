package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/eighthGnom/standard_web_server/internal/app/middelware"
	"github.com/eighthGnom/standard_web_server/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func (api *API) GetAllArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get All Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Article().SelectAll(): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	writer.WriteHeader(200)
	err = json.NewEncoder(writer).Encode(articles)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}
}

func (api *API) GetArticleByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Error while converting id to integer", err)
		message := Message{
			StatusCode: 400,
			Message:    "Invalid id value, use a value that can be converted to an integer",
			IsError:    true,
		}
		writer.WriteHeader(400)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	article, ok, err := api.storage.Article().FindByID(id)
	if err != nil {
		api.logger.Info("Error  while Article().FindByID(id): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	if !ok {
		api.logger.Info("Cant find article in database with id: ", id, " ", err)
		message := Message{
			StatusCode: 404,
			Message:    "Cant find article in database with this id",
			IsError:    true,
		}
		writer.WriteHeader(404)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	writer.WriteHeader(200)
	err = json.NewEncoder(writer).Encode(article)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}

}

func (api *API) DeleteArticleByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by ID DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Error while converting id to integer: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Invalid id value, use a value that can be converted to an integer",
			IsError:    true,
		}
		writer.WriteHeader(400)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	article, err := api.storage.Article().DeleteByID(id)
	if err != nil {
		api.logger.Info("Error while Article().DeleteByID(id): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	if article == nil {
		api.logger.Info("Cant find article in database with id: ", id, " ", err)
		message := Message{
			StatusCode: 404,
			Message:    "Cant find article in database with this id",
			IsError:    true,
		}
		writer.WriteHeader(404)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	message := Message{
		StatusCode: 200,
		Message:    "Resource deleted successfully",
		IsError:    false,
	}
	writer.WriteHeader(200)
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}
}

func (api *API) CreateArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Create Article POST /api/v1/articles")
	newArticle := &models.Article{}
	err := json.NewDecoder(request.Body).Decode(newArticle)
	if err != nil {
		api.logger.Info("Error while decoding article: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Cant decode this article, invalid format",
			IsError:    true,
		}
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	newArticle, err = api.storage.Article().Create(newArticle)
	if err != nil {
		api.logger.Info("Error while Article().Create(newArticle): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}

	writer.WriteHeader(201)
	err = json.NewEncoder(writer).Encode(newArticle)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}
}

func (api *API) UpdateArticleByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	api.logger.Info("Update Article by ID PUT /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Error while converting id to integer: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Invalid id value, use a value that can be converted to an integer",
			IsError:    true,
		}
		writer.WriteHeader(400)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}

	newArticle := &models.Article{}
	err = json.NewDecoder(request.Body).Decode(newArticle)
	if err != nil {
		api.logger.Info("Error while decoding article: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Cant decode this article, invalid format",
			IsError:    true,
		}
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}

	oldArticle, err := api.storage.Article().UpdateByID(id, newArticle)
	if err != nil {
		api.logger.Info("Error while Article().Update(newArticle): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	if oldArticle == nil {
		api.logger.Info("Cant find article in database: ", err)
		message := Message{
			StatusCode: 404,
			Message:    "Cant find article in database with this id",
			IsError:    true,
		}
		writer.WriteHeader(404)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	message := Message{
		StatusCode: 200,
		Message:    "Article updated successfully",
		IsError:    false,
	}
	writer.WriteHeader(200)
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}

}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User register POST /api/v1/user/register")
	user := &models.User{}
	err := json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		api.logger.Info("Error while decoding user: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Cant decode this user, invalid format",
			IsError:    true,
		}
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)

	if err != nil {
		api.logger.Info("Error while Article().Update(newArticle): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}

	if ok {
		api.logger.Info("Error while User().FindByLogin(user.Login): ", err)
		message := Message{
			StatusCode: 400,
			Message:    "User with this login already exists",
			IsError:    true,
		}
		writer.WriteHeader(400)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	user, err = api.storage.User().Create(user)
	if err != nil {
		api.logger.Info("Error while User().Create(user): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}

	writer.WriteHeader(201)
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}
}

func (api *API) PostToAuth(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth POST api/v1/user/auth")
	userFromRequest := &models.User{}
	err := json.NewDecoder(request.Body).Decode(userFromRequest)
	if err != nil {
		api.logger.Info("Error while decoding user: ", err)
		message := Message{
			StatusCode: 400,
			Message:    "Cant decode this user, invalid format",
			IsError:    true,
		}
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	userFromDatabase, ok, err := api.storage.User().FindByLogin(userFromRequest.Login)

	if err != nil {
		api.logger.Info("Error while User().Create(user): ", err)
		message := Message{
			StatusCode: 500,
			Message:    "Some troubles with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err := json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	if !ok {
		api.logger.Info("Error while User().FindByLogin(userFromRequest.Login), no user with that login")
		message := Message{
			StatusCode: 404,
			Message:    "Can't find user with that login",
			IsError:    true,
		}
		writer.WriteHeader(404)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	if userFromDatabase.Password != userFromRequest.Password {
		api.logger.Info("Error while userFromDatabase.Password != userFromRequest.Password, wrong password")
		message := Message{
			StatusCode: 404,
			Message:    "Wrong password",
			IsError:    true,
		}
		writer.WriteHeader(404)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()
	claims["admin"] = true
	fullToken, err := token.SignedString(middelware.SecretKey)
	if err != nil {
		api.logger.Info("Error while token.SignedString(middelware.SecretKey)", err)
		message := Message{
			StatusCode: 500,
			Message:    "We have some problems now, try again later",
			IsError:    true,
		}
		writer.WriteHeader(500)
		err = json.NewEncoder(writer).Encode(message)
		if err != nil {
			api.logger.Info("Cant encode: ", err)
		}
		return
	}
	writer.WriteHeader(201)
	message := Message{
		StatusCode: 201,
		Message:    fullToken,
		IsError:    false,
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		api.logger.Info("Cant encode: ", err)
	}
}
