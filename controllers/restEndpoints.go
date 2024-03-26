package controllers

import (
	"SpaceNewsWeb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindArtciles(c *gin.Context) {
	var articles []models.Article
	// updates articles with all found articles
	models.DB.Find(&articles)
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	article := models.Article{Title: input.Title, Author: input.Author}
	models.DB.Create(&article)

	c.JSON(http.StatusOK, gin.H{"data": article})
}

func FindArticle(c *gin.Context) {
	var article models.Article
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func UpdateArticle(c *gin.Context) {
	var article models.Article
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&article).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func DeleteArticle(c *gin.Context) {
	var article models.Article
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

/*
// validation
func ValidateArticle(article models.Article) error {
	err := validate.Struct(article)
	if err != nil {
		// Validation errors occurred
		return err
	}
	// Passed validation
	return nil
}
*/
