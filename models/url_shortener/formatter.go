package url_shortener

import "time"

type UrlGetFormatter struct {
	ID        int        `json:"id"`
	LongUrl   string     `json:"long_url"`
	ShortUrl  string     `json:"short_url"`
	Click     int        `json:"click"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// create
type UrlShortenerFormatter struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// update
type UrlUpdateFormatter struct {
	ID        int        `json:"id"`
	LongUrl   string     `json:"long_url"`
	ShortUrl  string     `json:"short_url"`
	Click     int        `json:"click"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func UrlsGetFormat(urlShortener UrlShortener) UrlGetFormatter {
	var urlShortenerFormatter UrlGetFormatter

	urlShortenerFormatter.ID = urlShortener.ID
	urlShortenerFormatter.LongUrl = urlShortener.LongUrl
	urlShortenerFormatter.ShortUrl = urlShortener.ShortUrl
	urlShortenerFormatter.Click = urlShortener.Click
	urlShortenerFormatter.CreatedAt = urlShortener.CreatedAt
	urlShortenerFormatter.UpdatedAt = urlShortener.UpdatedAt

	return urlShortenerFormatter
}

func UrlsGetAllFormat(urlShorteners []UrlShortener) []UrlGetFormatter {
	formatter := []UrlGetFormatter{}

	for _, urlShortener := range urlShorteners {
		UrlsGetFormatter := UrlsGetFormat(urlShortener)
		formatter = append(formatter, UrlsGetFormatter)
	}

	return formatter
}

func UrlsUpdateFormat(urlShortener UrlShortener) UrlUpdateFormatter {
	var urlShortenerFormatter UrlUpdateFormatter

	urlShortenerFormatter.ID = urlShortener.ID
	urlShortenerFormatter.LongUrl = urlShortener.LongUrl
	urlShortenerFormatter.ShortUrl = urlShortener.ShortUrl
	urlShortenerFormatter.Click = urlShortener.Click
	urlShortenerFormatter.CreatedAt = urlShortener.CreatedAt
	urlShortenerFormatter.UpdatedAt = urlShortener.UpdatedAt

	return urlShortenerFormatter
}

func UrlShortenerFormat(urlShortener UrlShortener) UrlShortenerFormatter {
	var urlShortenerFormatter UrlShortenerFormatter

	urlShortenerFormatter.LongUrl = urlShortener.LongUrl
	urlShortenerFormatter.ShortUrl = urlShortener.ShortUrl

	return urlShortenerFormatter
}
