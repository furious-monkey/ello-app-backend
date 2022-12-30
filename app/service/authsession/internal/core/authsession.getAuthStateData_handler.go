/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (c *AuthsessionCore) AuthsessionGetAuthStateData(in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	// TODO: not impl
	c.Logger.Errorf("authsession.getAuthStateData - error: method AuthsessionGetAuthStateData not impl")

	return nil, mtproto.ErrMethodNotImpl
}
