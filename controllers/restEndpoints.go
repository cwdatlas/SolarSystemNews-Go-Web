package controllers

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * FindArticles returns all articles in the database
 */
func FindArticles(c *gin.Context) {
	var articles []models.Article
	// updates articles with all found articles
	repo.DB.Find(&articles)
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
	article := models.Article{Title: input.Title, Author: input.Author, Body: input.Body}
	// todo add validation
	repo.DB.Create(&article)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

/*
 * FindAticle uses an id to find an article out of all articles in the database
 */
func FindArticle(c *gin.Context) {
	var article models.Article
	// uses where statement to find article
	if err := repo.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
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
	if err := repo.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
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
	repo.DB.Model(&article).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

/*
 * DeleteArtcile uses an ID to delete an article
 */
func DeleteArticle(c *gin.Context) {
	var article models.Article
	if err := repo.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// delete the article
	repo.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
