package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}