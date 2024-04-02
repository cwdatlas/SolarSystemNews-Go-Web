package services

import (
	"SpaceNewsWeb/models"
	"SpaceNewsWeb/repo"
	"fmt"
	"time"
)

/*
 * Author Aidan Scott
 * this file is for accessing the repo
 * Includes all CRUD operations
 */

type dbError struct {
	When time.Time
	What string
}

func (e *dbError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func Create(article *models.Article) {
	repo.DB.Create(article)
}

func UpdateByUArt(article *models.Article, update models.UpdateArticleInput) {
	repo.DB.Model(article).Updates(update)
}
func Update(article *models.Article, update models.UpdateArticleInput) {
	repo.DB.Model(article).Updates(update)
}

func DeleteA(column string, value string, article *models.Article) error {
	if err := GetBy(column, value, article); err != nil {
		return err
	}
	repo.DB.Delete(&article)
	return nil
}

func GetBy(column string, value string, article *models.Article) error {
	if err := repo.DB.Where(column+" = ?", value).First(article).Error; err != nil {
		return &dbError{time.Now(), "Could not find Article"}
	}
	return nil
}

func GetMany(article *[]models.Article) {
	repo.DB.Find(article)
}
