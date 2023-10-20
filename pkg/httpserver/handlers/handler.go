package handlers

import (
	_ "effectiveMobile/docs"
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/handlers/user/adduser"
	"effectiveMobile/pkg/httpserver/handlers/user/deleteuser"
	"effectiveMobile/pkg/httpserver/handlers/user/getuser"
	"effectiveMobile/pkg/httpserver/handlers/user/updateuser"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

// gin-swagger middleware
// swagger embed files

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

	router.GET("/", getuser.GetUser(r.Log, r.DB))

	router.POST("/", adduser.AddUser(r.Log, r.DB, adduser.AddUserBody{
		r.Urls.AgeUrl,
		r.Urls.GenderUrl,
		r.Urls.NationalityUrl,
	}))

	router.DELETE("/", deleteuser.DeleteUser(r.Log, r.DB))

	router.PUT("/", updateuser.UpdateUser(r.Log, r.DB))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
