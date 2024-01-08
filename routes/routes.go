package routes

import (
	"github.com/aditya3232/url-shortener/connection"
	"github.com/aditya3232/url-shortener/handler"
	"github.com/aditya3232/url-shortener/models/url_shortener"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	urlShortenerRepository := url_shortener.NewRepository(connection.DatabaseMysql())

	// Initialize services
	urlShortenerService := url_shortener.NewService(urlShortenerRepository)

	// Initialize handlers
	urlShortenerHandler := handler.NewURLShortener(urlShortenerService)

	// Configure routes
	api := router.Group("")

	urlShortenerRoutes := api.Group("")

	configureUrlShortenerRoutes(urlShortenerRoutes, urlShortenerHandler)
}

func configureUrlShortenerRoutes(api *gin.RouterGroup, handler *handler.URLShortener) {
	api.POST("/short", handler.CreateUrlShort)
	api.GET("/short/:short_key", handler.Redirect)
}
