package url_shortener

import "time"

type UrlShortener struct {
	ID        int        `gorm:"primary_key;column:id" json:"id"`
	LongUrl   string     `gorm:"column:long_url" json:"long_url"`
	ShortUrl  string     `gorm:"unique;column:short_url" json:"short_url"`
	Click     int        `gorm:"column:click" json:"click"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTTime" json:"updated_at"`
}

func (UrlShortener) TableName() string {
	return "urls"
}
