package repo

import (
	"SpaceNewsWeb/models"
	"bufio"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

/*
 * ConnectDatabase is used to connect to the database using a basic dsn
 * if database cant be connected to, the function will panic
 */
func ConnectDatabase() {
	// Basic dsn with inconsequential password and username
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	log.Println(dsn)
	// Attempt to connect
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	// If database can be connected to, create table based on Article model
	err = database.AutoMigrate(&models.Article{})
	if err != nil {
		return
	}
	// Set database to DB
	DB = database
}

/*
 * InitRepoData reads data from a file and formats it so it can be added to the database
 * If there arent articles in the database, then load it up with default articles
 */
func InitRepoData() {
	// Read stored files and up
	article := models.Article{}
	empty := DB.Take(&article)   // Pull an article from the database
	if empty.RowsAffected == 0 { // If there are no articles then continue with the function
		news := [2]string{"news/localNews.txt", "news/systemNews.txt"} // Slice of files articles are in
		for _, i := range news {
			log.Println("Searching through", news, "for news articles")
			dat, err := os.ReadFile(i)                                  // Read file
			check(err)                                                  // If there is an error panic
			scanner := bufio.NewScanner(strings.NewReader(string(dat))) // Read through the file one line at a time
			scanner.Split(bufio.ScanLines)                              // one line at a time
			line := 0                                                   // Based on the line from the start of each article, treat the line differently
			newArticle := models.Article{}
			for scanner.Scan() {
				text := scanner.Text() // get text of the line
				if text == "<->" {     // Element used to break up articles, set line to 0 when found
					line = 0
					DB.Create(&newArticle) // save created article to database
					log.Println("Article", newArticle.Title, "added to database")
					newArticle = models.Article{} // wipe out old article if new article is started
				} else if line == 0 { // if first line then save it as the title
					newArticle.Title = scanner.Text()
					line++
				} else if line == 1 { // if second line then save it as the author
					newArticle.Author = scanner.Text()
					line++
				} else if line == 2 { // if third line then save it as the location
					newArticle.Location = scanner.Text()
					line++
				} else if line > 2 { // if its after the third line save it as the body
					newArticle.Body += scanner.Text()
					line++
				}
			}
		}
	}
}

// if error doesnt equal nil, panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}
