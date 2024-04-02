package controllers

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/repo"
	"SpaceNewsWeb/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
 * @Author Aidan Scott
 * htmlEndpoints.go house the html related endpoints, including form endpoints
 * todo change repo.db access to service access of the database
 */

// Map used to pass errors to the article form
var formErrors = map[string]string{}

/*
 * Index provides an endpoint that allows the user to search the stored articles
 * or view all articles if no data is searched
 */
func Index(c *gin.Context) {
	var articles []models.Article
	error := ""                // data will be added if an error occurs
	title := c.Query("search") // If ?search= is in the get request, then get its value
	// if it is blank, then search for all articles sorted by popularity then location
	if len(title) < 1 {
		services.GetManyOrdered("popularity desc, location", &articles)
	} else { // Search by if contains search value non caps sensitive
		generalSearch := "%" + title + "%"
		if err := repo.DB.Where("title ILIKE ? OR location ILIKE ? OR author ILIKE ?", generalSearch, generalSearch, generalSearch).Find(&articles).Error; err != nil {
			log.Println("Error occurred when accepting form", err)
			error = "Something went wrong, please try again later"
		}
	}
	// Add models to HTML template
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articles": articles,
		"title":    "Comet News",
		"error":    error,
	})
}

/*
 * FormArticle provides the enpoint for the client to write and update an article
 */
func FormArticle(c *gin.Context) {
	// Add models to html template
	c.HTML(http.StatusOK, "createArticle.html", gin.H{
		"title":  "Comet News",
		"errors": formErrors,
	})
	// Clear formErrors because we optimistically assume all errors have been fixed.
	// If they haven't, they will show up once they are set again in the valid function
	formErrors = nil
}

/*
 * ConsumeArtForm, short for Consume Article Form, provides the post endpoint for Form submision
 */
func ConsumeArtForm(c *gin.Context) {
	formErrors = map[string]string{}
	update := c.PostForm("update")
	var newArticle models.Article
	// Setting local variables from form values
	newArticle.Title = c.PostForm("title")
	newArticle.Author = c.PostForm("author")
	newArticle.Location = c.PostForm("location")
	newArticle.Body = c.PostForm("body")
	// If user is updating an article
	if update == "on" {
		// Check to see if the  update article is valid, returns false if not
		if !validUpdate(newArticle) {
			log.Println("Form had errors", formErrors)
			c.Redirect(http.StatusFound, "/create_article")
			return
		}
		var article models.Article
		if err := repo.DB.Where("title = ?", newArticle.Title).First(&article).Error; err != nil {
			formErrors["global"] = "No Article named " + newArticle.Title + " found."
			log.Println(formErrors["global"])
			c.Redirect(http.StatusFound, "/create_article")
			return
		}
		// set oldArticle to articles value before updated
		oldArticle := article
		repo.DB.Model(&article).Updates(newArticle)
		// Checking if article was really added
		if oldArticle == article {
			formErrors["global"] = "Could Not Update Article"
			c.Redirect(http.StatusFound, "/create_article")
			return
		}
		// YAY! article was updated!
		formErrors["success"] = "Article updated successfully"
		log.Println("Article", newArticle.Title, "added to the database")
		c.Redirect(http.StatusFound, "/create_article")
		return
	} else {
		// Article needs to be added from scratch
		if !valid(newArticle) {
			c.Redirect(http.StatusFound, "/create_article")
			return
		}
		if err := repo.DB.Create(&newArticle).Error; err != nil {
			formErrors["global"] = "Something went wrong adding your article to our system, try again later"
			c.Redirect(http.StatusFound, "/create_article")
			log.Println(err)
			return
		}
		formErrors["success"] = "Article created successfully"
		c.Redirect(http.StatusFound, "/create_article")
	}
}

/*
 * Valid is specifically for articles being added via a form
 */
func valid(article models.Article) bool {
	// Title validation, only one issue can be sent back at one time
	if article.Title == "" {
		formErrors["title"] = "Must include title"
	} else if len(article.Title) > 40 {
		formErrors["title"] = "Title length must be less than 40 characters"
	}
	// Location validation
	if article.Location == "" {
		formErrors["location"] = "Must include a location"
	} else if len(article.Location) > 40 {
		formErrors["location"] = "Location length must be less than 40 characters"
	}
	// Author/date validation
	if article.Author == "" {
		formErrors["author"] = "Must include an author/date"
	} else if len(article.Author) > 40 {
		formErrors["author"] = "Author/date must be less than 40 characters"
	}
	// Body validation
	if article.Body == "" {
		formErrors["body"] = "Must include a body"
	} else if len(article.Body) > 500 {
		formErrors["body"] = "Body must be less than 500 characters"
	}
	// If there are errors, return false
	return len(formErrors) == 0
}

/*
 * When updating an article validateUpdate should be used as it allows for blank entries
 */
func validUpdate(article models.Article) bool {
	// Title validation, only one issue can be sent back at one time
	if article.Title != "" && len(article.Title) > 40 {
		formErrors["title"] = "Title length must be less than 40 characters"
	}
	// Location validation
	if article.Location != "" && len(article.Location) > 40 {
		formErrors["location"] = "Location length must be less than 40 characters"
	}
	// Author/date validation
	if article.Author != "" && len(article.Author) > 40 {
		formErrors["author"] = "Author/date must be less than 40 characters"
	}
	// Body validation
	if article.Body != "" && len(article.Body) > 500 {
		formErrors["body"] = "Body must be less than 500 characters"
	}
	return len(formErrors) == 0
}
