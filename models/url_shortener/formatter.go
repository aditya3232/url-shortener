package url_shortener

type UrlShortenerFormatter struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func UrlShortenerFormat(urlShortener UrlShortener) UrlShortenerFormatter {
	var urlShortenerFormatter UrlShortenerFormatter

	urlShortenerFormatter.LongUrl = urlShortener.LongUrl
	urlShortenerFormatter.ShortUrl = urlShortener.ShortUrl

	return urlShortenerFormatter
}
