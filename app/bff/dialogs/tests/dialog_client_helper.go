package tests

import (
	dialogs_client "github.com/teamgram/teamgram-server/app/bff/dialogs/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewDialogClient() dialogs_client.DialogsClient {
	return dialogs_client.NewDialogsClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "bff.bff",
			InsecureSkipVerify: true,
		},
	}))
}
