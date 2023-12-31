package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserUpdatePredefinedUsername
// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedUsername(in *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedUsername - error: method UserUpdatePredefinedUsername not impl")

	return nil, mtproto.ErrMethodNotImpl
}
