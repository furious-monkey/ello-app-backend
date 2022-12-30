package photos_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/photos/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/photos/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/photos/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
