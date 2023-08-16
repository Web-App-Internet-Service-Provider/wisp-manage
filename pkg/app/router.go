package app

import (
	"encoding/gob"
	"net/http"

	"github.com/AbdulrahmanDaud10/wisp-manage/pkg/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func RouteSetup(auth *api.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "template/static")
	router.LoadHTMLGlob("../template/web/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/", HomeHandler)
	router.GET("/login", LoginHandler(auth))
	router.GET("/callback", CallBackHandler(auth))
	router.GET("/user", api.IsAuthenticated, UserHandler)
	router.GET("/logout", LogoutHandler)

	return router
}
