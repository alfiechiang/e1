package service

import (
	_ "fmt"
	_ "e1/src/dao"

	"github.com/gin-gonic/gin"
)

//GreateArticle
func (svc *Service) GreateArticle(c *gin.Context) {
	svc.Dao.GreateArticle(c)
}

func (svc *Service) DeleteArticle(articleId int32) {
	svc.Dao.DeleteArticle(articleId)
}

func (svc *Service) UpdateArticle(articleId uint32,c *gin.Context) {
	svc.Dao.UpdateArticle(articleId,c)
}

func  (svc *Service) ListArticle(c *gin.Context) ([]struct{}){
	return svc.Dao.ListArticle(c)
}

func (svc *Service) GetArticle(c *gin.Context) {
	svc.Dao.GetArticle()
}
