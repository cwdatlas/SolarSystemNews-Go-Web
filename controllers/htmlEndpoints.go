package controllers

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

var data = map[string]string{}

func Index(c *gin.Context) {
	var articles []models.Article
	repo.DB.Find(&articles)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articles": articles,
		"title":    "Comet News",
	})
}

func FrontPage(c *gin.Context) {
	data["title"] = "You made it to front page"
	c.HTML(http.StatusOK, "index.html", data)
}

func PostTest(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", data)
}
