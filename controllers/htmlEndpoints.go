package controllers

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

var formErrors = map[string]string{}

func Index(c *gin.Context) {
	var articles []models.Article
	repo.DB.Order("popularity desc, location").Find(&articles)
	repo.DB.Find(&articles)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articles": articles,
		"title":    "Comet News",
	})
}

func FormArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "createArticle.html", gin.H{
		"title":  "Comet News",
		"errors": formErrors,
	})
	formErrors = nil
}

func ConsumeArtForm(c *gin.Context) {
	formErrors = map[string]string{}
	update := c.PostForm("update")
	if update == "on" {
		var newArticle models.Article
		// This will infer what binder to use depending on the content-type header.
		newArticle.Title = c.PostForm("title")
		newArticle.Author = c.PostForm("author")
		newArticle.Location = c.PostForm("location")
		newArticle.Body = c.PostForm("body")
		if !valid(newArticle) {
			c.Redirect(http.StatusFound, "/create_article")
		}
		var article models.Article
		if err := repo.DB.Where("Title = ?", newArticle.Title).First(&article).Error; err != nil {
			formErrors["global"] = "No Article named " + newArticle.Title + " found."
			c.Redirect(http.StatusFound, "/create_article")
		}
		// todo make sure article actually got updated
		repo.DB.Model(&article).Updates(newArticle)
		formErrors["success"] = "Article updated successfully"
		c.Redirect(http.StatusFound, "/create_article")
	} else {
		var newArticle models.Article
		newArticle.Title = c.PostForm("title")
		newArticle.Author = c.PostForm("author")
		newArticle.Location = c.PostForm("location")
		newArticle.Body = c.PostForm("body")
		if !valid(newArticle) {
			c.Redirect(http.StatusFound, "/create_article")
		}
		repo.DB.Create(&newArticle)
		formErrors["success"] = "Article created successfully"
		c.Redirect(http.StatusFound, "/create_article")
	}
}

func valid(article models.Article) bool {
	if article.Title == "" {
		formErrors["title"] = "Must include title of article"
	} else if len([]rune(article.Title)) > 40 {
		formErrors["title"] = "Title length must be less than 40 characters"
	}

	return true
}
