package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserUpdatePredefinedCode
// user.updatePredefinedCode phone:string code:string = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedCode(in *user.TLUserUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedCode - error: method UserUpdatePredefinedCode not impl")

	return nil, mtproto.ErrMethodNotImpl
}
