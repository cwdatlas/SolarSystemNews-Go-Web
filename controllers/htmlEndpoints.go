package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DisplayIndex(c *gin.Context) {
	message := c.Param("id")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": message,
	})
}

func FrontPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "You made it to front page",
	})
}

func PostTest(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": "Posted",
	})
}
