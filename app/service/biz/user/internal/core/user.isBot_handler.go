package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserIsBot
// user.isBot id:long = Bool;
func (c *UserCore) UserIsBot(in *user.TLUserIsBot) (*mtproto.Bool, error) {
	userData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetId())
	if userData == nil {
		c.Logger.Errorf("user.isBot - error: invalid user(%d)", in.GetId())
		return nil, mtproto.ErrUserIdInvalid
	}

	return mtproto.ToBool(userData.GetUserData().GetBot() != nil), nil
}
