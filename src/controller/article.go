package controller

import (
//	"fmt"
	_ "e1/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (control *Controller) GreateArticle(c *gin.Context) {
	control.Svc.GreateArticle(c)
	//GreateArticle
}

func (control *Controller) GetArticles(c *gin.Context) {
	for {
		//fmt.Println("GOGIIGI")
	}
	//control.Svc.GetArticle(c)
}

func (control *Controller) DeleteArticle(c *gin.Context) {

	articleId, _ := strconv.Atoi(c.Param("articleId"))
	control.Svc.DeleteArticle(int32(articleId))
}

func (control *Controller) UpdateArticle(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	control.Svc.UpdateArticle(uint32(articleId), c)
}

func (control *Controller) ListArticle(c *gin.Context) {
	ob := control.Svc.ListArticle(c)
	c.JSON(200, gin.H{
		"data": ob,
	})
}
