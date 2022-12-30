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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (c *MediaCore) MediaGetPhoto(in *media.TLMediaGetPhoto) (*mtproto.Photo, error) {
	photo, err := c.svcCtx.Dao.GetPhotoV2(c.ctx, in.GetPhotoId())
	if err != nil {
		c.Logger.Errorf("media.getPhotoFileData(%d) - error: %v", in.GetPhotoId(), err)
		return nil, err
	}

	return photo, nil
}
