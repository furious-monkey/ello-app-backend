package core

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/status/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	onlineKeyPrefix  = "online"           //
	userKeyIdsPrefix = "user_online_keys" //
)

type StatusCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *StatusCore {
	return &StatusCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func getUserKey(id int64) string {
	return fmt.Sprintf("%s_%d", userKeyIdsPrefix, id)
}

func getAuthKeyIdKey(id int64) string {
	return fmt.Sprintf("%s_%d", onlineKeyPrefix, id)
}
