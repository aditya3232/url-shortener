package url_shortener

// create && update url shortener
type UrlShortenerInput struct {
	ID       int    `form:"id"`
	LongUrl  string `form:"long_url" binding:"required"`
	ShortUrl string `form:"short_url"`
	Click    int    `form:"click"`
}

type UrlGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
