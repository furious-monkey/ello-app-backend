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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserGetNotifySettings
// user.getNotifySettings user_id:int peer_type:int peer_id:int = PeerNotifySettings;
func (c *UserCore) UserGetNotifySettings(in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	settings, err := c.svcCtx.Dao.GetUserNotifySettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetUserId())

	if err != nil {
		c.Logger.Errorf("user.getNotifySettings - error: %v", err)
		return nil, err
	}

	return settings, nil
}
