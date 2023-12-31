package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(db, sqlc.NewConn(db, c.Cache)),
	}
}
