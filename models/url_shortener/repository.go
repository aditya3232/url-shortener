package url_shortener

import (
	"github.com/aditya3232/url-shortener/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]UrlShortener, helper.Pagination, error)
	GetOne(UrlShortener) (UrlShortener, error)
	Create(UrlShortener) (UrlShortener, error)
	Update(UrlShortener) (UrlShortener, error)
	Delete(UrlShortener) error
	FindByShortKey(UrlShortener) (UrlShortener, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]UrlShortener, helper.Pagination, error) {
	var urlShorteners []UrlShortener
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&urlShorteners), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return urlShorteners, helper.Pagination{}, err
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&urlShorteners).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(urlShorteners)

	return urlShorteners, pagination, nil
}

func (r *repository) GetOne(urlShortener UrlShortener) (UrlShortener, error) {
	err := r.db.Where("id = ?", urlShortener.ID).First(&urlShortener).Error
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
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

func (r *repository) Delete(urlShortener UrlShortener) error {
	err := r.db.Where("id = ?", urlShortener.ID).Delete(&urlShortener).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByShortKey(urlShortener UrlShortener) (UrlShortener, error) {

	err := r.db.Model(&urlShortener).Where("short_url = ?", urlShortener.ShortUrl).First(&urlShortener).Error
	if err != nil {
		return urlShortener, err
	}

	return urlShortener, nil
}
