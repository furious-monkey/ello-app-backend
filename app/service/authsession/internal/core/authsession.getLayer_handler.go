package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (c *AuthsessionCore) AuthsessionGetLayer(in *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getUserId - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	layer := c.svcCtx.GetApiLayer(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLInt32(&mtproto.Int32{
		V: layer,
	}).To_Int32(), nil
}
