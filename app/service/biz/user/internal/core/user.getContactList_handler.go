package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
)

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (c *UserCore) UserGetContactList(in *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	rValList := &user.Vector_ContactData{}

	rValList.Datas = c.svcCtx.Dao.GetUserContactList(c.ctx, in.GetUserId())
	if rValList.Datas == nil {
		rValList.Datas = []*user.ContactData{}
	}

	return rValList, nil
}
