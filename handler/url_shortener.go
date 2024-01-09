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
	input.ShortUrl = shortKey // input buat cari long_url where short_url, disini cuma param aja

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

	urlShortener.Click = urlShortener.Click + 1
	input = url_shortener.UrlShortenerInput{
		ShortUrl: urlShortener.ShortUrl, // ini short_url real, dari database
		Click:    urlShortener.Click,
	}

	_, err = h.urlShortenerService.Update(input)
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

	// redirect
	c.Redirect(http.StatusMovedPermanently, urlShortener.LongUrl)

}

func (h *URLShortener) GetAll(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, url_shortener.UrlShortener{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	urls, page, err := h.urlShortenerService.GetAll(filter, page, sort)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(urls) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := fmt.Sprintf("Users not found")
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIDataTableResponse(message, http.StatusOK, page, url_shortener.UrlsGetAllFormat(urls))
	c.JSON(response.Meta.Code, response)
}
