package url_shortener

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aditya3232/url-shortener/config"
)

type Service interface {
	GenerateShortKey() string
	Create(input UrlShortenerInput) (UrlShortener, error)
	Update(input UrlShortenerInput) (UrlShortener, error)
	FindByShortKey(input UrlShortenerInput) (UrlShortener, error)
}

type service struct {
	urlShortenerRepository Repository
}

func NewService(urlShortenerRepository Repository) *service {
	return &service{urlShortenerRepository}
}

const (
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength = 6
)

func (s *service) GenerateShortKey() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[random.Intn(len(charset))]
	}

	return string(shortKey)
}

func (s *service) Create(input UrlShortenerInput) (UrlShortener, error) {
	shortKey := s.GenerateShortKey()
	shortenedURL := fmt.Sprintf("http://localhost:3636/short/%s", shortKey)

	urlShortener := UrlShortener{
		LongUrl:  input.LongUrl,
		ShortUrl: shortenedURL,
		Click:    input.Click,
	}

	newUrlShortener, err := s.urlShortenerRepository.Create(urlShortener)
	if err != nil {
		return newUrlShortener, err
	}

	return newUrlShortener, nil
}

func (s *service) Update(input UrlShortenerInput) (UrlShortener, error) {
	urlShortener := UrlShortener{
		LongUrl:  input.LongUrl,
		ShortUrl: input.ShortUrl,
		Click:    input.Click,
	}

	newUrlShortener, err := s.urlShortenerRepository.Update(urlShortener)
	if err != nil {
		return newUrlShortener, err
	}

	return newUrlShortener, nil
}

func (s *service) FindByShortKey(input UrlShortenerInput) (UrlShortener, error) {
	fullUrlWithPrefix := config.CONFIG.URL_SHORT_PREFIX + input.ShortUrl

	newUrlShortener, err := s.urlShortenerRepository.FindByShortKey(fullUrlWithPrefix)
	if err != nil {
		return newUrlShortener, err
	}

	return newUrlShortener, nil
}
