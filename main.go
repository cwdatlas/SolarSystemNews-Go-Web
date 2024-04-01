package main

import (
	"SpaceNewsWeb/controllers"
	"SpaceNewsWeb/repo"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
 * Author: Aidan Scott
 * Main routs all traffic to their respective functions
 */
func main() {
	// Gets default router
	r := gin.Default()

	// Connect to database
	repo.ConnectDatabase()
	// Init data in database, go to repo to learn more
	repo.InitRepoData()

	//Rest endpoints
	// Provides CRUD methods
	r.GET("articles", controllers.FindArticles)
	r.POST("article", controllers.CreateArticle)
	r.GET("articles/:id", controllers.FindArticle)
	r.PATCH("articles/:id", controllers.UpdateArticle)
	r.DELETE("articles/:id", controllers.DeleteArticle)

	// Loads up templates to view
	r.LoadHTMLGlob("templates/*")

	//Form and HTML endpoints
	r.GET("/", controllers.Index)                         // Provides no param endpoint for index
	r.GET("/:article_title", controllers.Index)           // provides search mechanism for index
	r.GET("/create_article", controllers.FormArticle)     // Page that provides form
	r.POST("/create_article", controllers.ConsumeArtForm) // Endpoint to post form to

	// If error then print error and shutdown
	err := r.Run()
	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}
}
