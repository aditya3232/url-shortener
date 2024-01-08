package handler

import (
	"fmt"
	"net/http"

	"github.com/aditya3232/url-shortener/constant"
	"github.com/aditya3232/url-shortener/helper"
	log_function "github.com/aditya3232/url-shortener/log"
	"github.com/aditya3232/url-shortener/models/url_shortener"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type URLShortener struct {
	urlShortenerService url_shortener.Service
}

func NewURLShortener(urlShortenerService url_shortener.Service) *URLShortener {
	return &URLShortener{urlShortenerService}
}

func (h *URLShortener) CreateUrlShort(c *gin.Context) {
	var input url_shortener.UrlShortenerInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUrlShortener, err := h.urlShortenerService.Create(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusCreated, url_shortener.UrlShortenerFormat(newUrlShortener))
	c.JSON(response.Meta.Code, response)

}

// redirect and add click
func (h *URLShortener) Redirect(c *gin.Context) {
	var input url_shortener.UrlShortenerInput

	shortKey := c.Param("short_key")
	input.ShortUrl = shortKey

	urlShortener, err := h.urlShortenerService.FindByShortKey(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	fmt.Println(urlShortener.LongUrl)

	// urlShortener.Click = urlShortener.Click + 1
	// input = url_shortener.UrlShortenerInput{
	// 	Click: urlShortener.Click,
	// }

	// _, err = h.urlShortenerService.Update(input)
	// if err != nil {
	// 	endpoint := c.Request.URL.Path
	// 	message := constant.InvalidRequest
	// 	errorCode := http.StatusBadRequest
	// 	ipAddress := c.ClientIP()
	// 	errors := helper.FormatError(err)
	// 	log_function.Error(message, errors, endpoint, errorCode, ipAddress)

	// 	response := helper.APIResponse(message, http.StatusBadRequest, nil)
	// 	c.JSON(response.Meta.Code, response)
	// 	return
	// }

	// redirect
	// c.Redirect(http.StatusMovedPermanently, urlShortener.LongUrl)
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}
