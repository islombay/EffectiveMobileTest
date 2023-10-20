package handlers

import (
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/handlers/user/adduser"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type HandlerInitUrls struct {
	AgeUrl, GenderUrl, NationalityUrl string
}

type HandlerInitBody struct {
	Log  *slog.Logger
	DB   *database.Database
	Urls HandlerInitUrls
}

func InitRoutes(r HandlerInitBody) *gin.Engine {
	router := gin.Default()

	router.POST("/", adduser.AddUser(r.Log, r.DB, adduser.AddUserBody{
		r.Urls.AgeUrl,
		r.Urls.GenderUrl,
		r.Urls.NationalityUrl,
	}))

	router.DELETE("/")

	return router
}
