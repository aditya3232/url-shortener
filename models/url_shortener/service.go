package url_shortener

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aditya3232/url-shortener/config"
	"github.com/aditya3232/url-shortener/helper"
)

type Service interface {
	GenerateShortKey() string
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]UrlShortener, helper.Pagination, error)
	GetOne(input UrlGetOneByIdInput) (UrlShortener, error)
	Update(input UrlShortenerInput) (UrlShortener, error)
	Delete(input UrlGetOneByIdInput) error
	Create(input UrlShortenerInput) (UrlShortener, error)
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

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]UrlShortener, helper.Pagination, error) {
	urlShorteners, pagination, err := s.urlShortenerRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return urlShorteners, pagination, nil
}

func (s *service) GetOne(input UrlGetOneByIdInput) (UrlShortener, error) {
	// var urlShortener UrlShortener
	// urlShortener.ID = input.ID
	urlShortenerID := UrlShortener{ID: input.ID}

	urlShortener, err := s.urlShortenerRepository.GetOne(urlShortenerID)
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
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

func (s *service) Delete(input UrlGetOneByIdInput) error {
	urlShortenerID := UrlShortener{ID: input.ID}

	_, err := s.urlShortenerRepository.GetOne(urlShortenerID)
	if err != nil {
		return err
	}

	err = s.urlShortenerRepository.Delete(urlShortenerID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByShortKey(input UrlShortenerInput) (UrlShortener, error) {
	fullUrlWithPrefix := config.CONFIG.URL_SHORT_PREFIX + input.ShortUrl
	shortUrl := UrlShortener{ShortUrl: fullUrlWithPrefix}

	newUrlShortener, err := s.urlShortenerRepository.FindByShortKey(shortUrl)
	if err != nil {
		return newUrlShortener, err
	}

	return newUrlShortener, nil
}
