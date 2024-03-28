package repo

import (
	"SpaceNewsWeb/models"
	"bufio"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	fmt.Print(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&models.Article{})
	if err != nil {
		return
	}
	DB = database
}

func InitRepoData() {
	// Read stored files and up
	article := models.Article{}
	empty := DB.Take(&article)
	if empty.RowsAffected == 0 {
		news := [2]string{"news/localNews.txt", "news/systemNews.txt"}
		for _, i := range news {
			dat, err := os.ReadFile(i)
			check(err)
			scanner := bufio.NewScanner(strings.NewReader(string(dat)))
			scanner.Split(bufio.ScanLines)
			line := 0
			newArticle := models.Article{}
			for scanner.Scan() {
				text := scanner.Text()
				if text == "<->" {
					line = 0
					DB.Create(&newArticle)
					newArticle = models.Article{}
				} else if line == 0 {
					newArticle.Title = scanner.Text()
					line++
				} else if line == 1 {
					newArticle.Author = scanner.Text()
					line++
				} else if line == 2 {
					newArticle.Location = scanner.Text()
					line++
				} else if line > 2 {
					newArticle.Body += scanner.Text()
					line++
				}
			}
		}
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
