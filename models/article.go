package models

/*
 * Author: Aidan Scott
 * article.go houses the couple types of articles used for validation and storage
 */

// Used to store articles in database
type Article struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"Title"`
	Author     string `json:"author"`
	Location   string `json:"location"`
	Body       string `json:"body"`
	Popularity int    `json:"popularity"`
}

// Used for creation validation
type CreateArticleInput struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Location string `json:"location" binding:"required"`
	Body     string `json:"body" binding:"required"`
}

// Used for article update validation
type UpdateArticleInput struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Location string `json:"location"`
	Body     string `json:"body"`
}
