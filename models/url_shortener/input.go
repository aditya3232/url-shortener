package url_shortener

type UrlShortenerInput struct {
	LongUrl  string `form:"long_url" binding:"required"`
	ShortUrl string `form:"short_url"`
	Click    int    `form:"click"`
}
