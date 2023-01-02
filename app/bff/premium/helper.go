package premium_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/premium/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/premium/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/premium/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
