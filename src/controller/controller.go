package controller

import (
	"e1/src/dao"
	"e1/src/service"
)

type Controller struct {
	Svc *service.Service
	dao *dao.Dao
}
