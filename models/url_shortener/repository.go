package url_shortener

import "gorm.io/gorm"

type Repository interface {
	Create(UrlShortener) (UrlShortener, error)
	Update(UrlShortener) (UrlShortener, error)
	FindByShortKey(ShortUrl string) (UrlShortener, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(urlShortener UrlShortener) (UrlShortener, error) {
	err := r.db.Model(&urlShortener).Create(&urlShortener).Error
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
}

func (r *repository) Update(urlShortener UrlShortener) (UrlShortener, error) {
	err := r.db.Model(&urlShortener).Where("short_url = ?", urlShortener.ShortUrl).Updates(&urlShortener).Error
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
}

func (r *repository) FindByShortKey(ShortUrl string) (UrlShortener, error) {
	var urlShortener UrlShortener

	err := r.db.Model(&urlShortener).Where("short_url = ?", ShortUrl).First(&urlShortener).Error
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
}
