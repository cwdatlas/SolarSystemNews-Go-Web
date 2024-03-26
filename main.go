package main

import (
	"SpaceNewsWeb/controllers"
	"SpaceNewsWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/articles", controllers.FindArtciles)
	r.POST("/article", controllers.CreateArticle)
	r.GET("/articles/:id", controllers.FindArticle)
	r.PATCH("articles/:id", controllers.UpdateArticle)
	r.DELETE("articles/:id", controllers.DeleteArticle)

	//html
	r.LoadHTMLGlob("templates/*")
	r.GET("/:id", controllers.DisplayIndex)
	r.GET("/front_page", controllers.FrontPage)
	r.GET("post", controllers.PostTest)
	err := r.Run()
	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}
}
