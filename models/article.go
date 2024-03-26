package models

import (
	"time"
)

type Article struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"Title"`
	Author    string `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateArticleInput struct {
	Title  string `json:"title" binding:"required" gorm:"size:5"`
	Author string `json:"author" binding:"required" gorm:"size:5"`
}

type UpdateArticleInput struct {
	Title  string `json:"title" gorm:"size:40"`
	Author string `json:"author" gorm:"size:40"`
}
