package tests

import (
	"context"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/feeds/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
	"testing"
)

func TestUpdateFeedList(t *testing.T) {
	c := NewRPCClient()
	/*
		{"server_id":"127.0.0.1:20120",
		"client_addr":"192.168.1.5",
		"auth_id":7338124102345237054,
		"session_id":-6066913564088792595,
		"receive_time":1671767460,
		"user_id":777062,
		"client_msg_id":7180186570927840256,
		"layer":147,
		"client":"android",
		"langpack":"android",
		"perm_auth_key_id":9167378892795598833}
	*/
	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		ServerId:      "127.0.0.1:20120",
		ClientAddr:    "192.168.1.105",
		AuthId:        7338124102345237054,
		SessionId:     -6066913564088792595,
		ReceiveTime:   1671767460,
		UserId:        777062,
		ClientMsgId:   7180186570927840256,
		IsBot:         false,
		Layer:         147,
		Client:        "android",
		IsAdmin:       false,
		Langpack:      "android",
		PermAuthKeyId: 9167378892795598833,
	})
	if err != nil {
		panic(err)
	}
	/*
		empty request
	*/
	op, err := op_srv.NewOperation(op_srv.Operation{
		Service: op_srv.Feeds,
		Method:  service.UpdateFeedList,
		Data: []*feeds.FeedInsertItemDO{
			{ChatId: 12, PeerType: 2},
			{ChatId: 13, PeerType: 2},
			{ChatId: 4, PeerType: 2},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.BizInvokeBizDataRaw(ctx, op)
	if err != nil {
		t.Fatal(err)
	}
}
