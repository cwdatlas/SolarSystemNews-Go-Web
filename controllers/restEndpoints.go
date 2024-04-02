package controllers

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * Author: Aidan Scott
 * restEndpoints.go is for all Rest related endpoints
 * todo create validation service and implement it here
 */

/*
 * FindArticles returns all articles in the database
 */
func FindArticles(c *gin.Context) {
	var articles []models.Article
	// updates articles with all found articles
	services.GetMany(&articles)
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

/*
 * CreateArticle adds an article to the database
 */
func CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	// binds json from the context to the input variable
	if err := c.ShouldBindJSON(&input); err != nil {
		// If cant bind, then return error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create article then save it
	article := models.Article{Title: input.Title, Author: input.Author, Body: input.Body, Location: input.Location}
	// todo add validation
	services.Create(&article)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

/*
 * FindAticle uses an id to find an article out of all articles in the database
 */
func FindArticle(c *gin.Context) {
	var article models.Article
	// uses where statement to find article
	err := services.GetBy("id", c.Param("id"), &article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Return article if able to add to database
	c.JSON(http.StatusOK, gin.H{"data": article})
}

/*
 * UpdateArticle updates an article based on its id
 */
func UpdateArticle(c *gin.Context) {
	var article models.Article
	// Where to find article
	if err := services.GetBy("id", c.Param("id"), &article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Bind json
	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update article
	services.UpdateByUArt(&article, input)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

/*
 * DeleteArtcile uses an ID to delete an article
 */
func DeleteArticle(c *gin.Context) {
	var article models.Article
	// delete the article
	if err := services.DeleteA("id", c.Param("id"), &article); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
