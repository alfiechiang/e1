package dao

import (
	"fmt"
	"e1/global"

	"github.com/gin-gonic/gin"
)

type BlogArticle struct {
	ID            uint32 `gorm:"primary_key" json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (d *Dao) GreateArticle(c *gin.Context) {
	var article BlogArticle

	c.ShouldBindJSON(&article)
	db := global.DBEngine
	db.Table("blog_article").Create(&article)
}

func (d *Dao) DeleteArticle(articleId int32) {
	db := global.DBEngine

	db.Table("blog_article").Delete(&BlogArticle{}, articleId)
}

func (d *Dao) UpdateArticle(articleId uint32, c *gin.Context) {
	db := global.DBEngine
	var article BlogArticle
	c.ShouldBindJSON(&article)
	article.ID = articleId
	db.Table("blog_article").Save(&article)
}


func  (d *Dao) ListArticle(c *gin.Context) ([]struct{}){
	db:= global.DBEngine
	var articles  []struct{}
	db.Table("blog_article").Find(&articles)
	return articles
}

func (d *Dao) GetArticle() {
	var article BlogArticle
	db := global.DBEngine
	db.Table("blog_article").First(&article, 1)

	fmt.Println(article)
}
