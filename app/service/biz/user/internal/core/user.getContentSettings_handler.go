package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (c *UserCore) UserGetContentSettings(in *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	var (
		rV = false
	)

	if do, _ := c.svcCtx.Dao.UserSettingsDAO.SelectByKey(c.ctx, in.UserId, "sensitive_enabled"); do != nil {
		if do.Value == "true" {
			rV = true
		}
	}

	return mtproto.MakeTLAccountContentSettings(&mtproto.Account_ContentSettings{
		SensitiveEnabled:   rV,
		SensitiveCanChange: true,
	}).To_Account_ContentSettings(), nil
}
