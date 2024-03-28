package main

import (
	"SpaceNewsWeb/controllers"
	"SpaceNewsWeb/repo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo.ConnectDatabase()
	repo.InitRepoData()

	r.GET("/articles", controllers.FindArticles)
	r.POST("/article", controllers.CreateArticle)
	r.GET("/articles/:id", controllers.FindArticle)
	r.PATCH("articles/:id", controllers.UpdateArticle)
	r.DELETE("articles/:id", controllers.DeleteArticle)

	//html
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Index)
	err := r.Run()
	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}
}
