package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *SecretchatsCore) MessagesDiscardEncryption(in *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &mtproto.Bool{}, nil
}
