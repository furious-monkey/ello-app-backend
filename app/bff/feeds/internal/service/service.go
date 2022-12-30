package service

import (
	"context"
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/feeds/internal/svc"
)

type Service struct {
	svcCtx            *svc.ServiceContext
	operationRegister map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error)
}

func New(svcCtx *svc.ServiceContext) *Service {
	srv := &Service{
		svcCtx: svcCtx,
	}
	srv.initOperationRegister()
	return srv
}
