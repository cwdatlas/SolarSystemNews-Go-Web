package models

type Article struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"Title"`
	Author     string `json:"author"`
	Location   string `json:"location"`
	Body       string `json:"body"`
	Popularity int    `json:"popularity"`
}

type CreateArticleInput struct {
	Title    string `json:"title" binding:"required" gorm:"size:5"`
	Author   string `json:"author" binding:"required" gorm:"size:5"`
	Location string `json:"location" binding:"required" gorm:"size:20"`
	Body     string `json:"body" binding:"required" gorm:"size:500"`
}

type UpdateArticleInput struct {
	Title    string `json:"title" gorm:"size:40"`
	Author   string `json:"author" gorm:"size:40"`
	Location string `json:"location" gorm:"size:20"`
	Body     string `json:"body" gorm:"size:500"`
}
