package app

import (
	"github.com/adataou/bookstore_oauth-api/src/services/access_token"
	"github.com/adataou/bookstore_oauth-api/src/repository/db"
	"github.com/adataou/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
	"github.com/adataou/bookstore_oauth-api/src/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
