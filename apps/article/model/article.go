package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Image   string `gorm:"type:varchar(255);not null" json:"image"`
	Author  string `gorm:"type:varchar(255);not null" json:"author"`
}

func (*Article) LookAllArticle(db *gorm.DB) (*[]Article, error) {
	articles := []Article{}
	err := db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return &articles, nil
}
