package core

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/rpc/metadata"
)

type MessagesCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *MessagesCore {
	return &MessagesCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
