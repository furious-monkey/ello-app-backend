// Code generated by goctl. DO NOT EDIT.
// Source: phone_call.proto

package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func New(svcCtx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: svcCtx,
	}
}
