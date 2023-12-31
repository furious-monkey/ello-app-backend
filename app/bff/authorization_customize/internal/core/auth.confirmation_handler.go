package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
)

type ConfirmationReq struct {
	UsernameOrEmail string `json:"username_or_email"`
	Code            string `json:"code"`
}

type ConfirmationResp struct {
	Message string `json:"message"`
}

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) Confirmation(in json.RawMessage) (*ConfirmationResp, error) {
	var req ConfirmationReq
	if err := json.Unmarshal(in, &req); err != nil {
		return nil, err
	}

	out, err := c.svcCtx.Dao.AuthorizationClient.Confirmation(c.ctx, &authorization.ConfirmationRequest{
		UsernameOrEmail: req.UsernameOrEmail,
		Code:            req.Code,
	})
	if err != nil {
		return nil, err
	}

	return &ConfirmationResp{
		Message: out.Message,
	}, nil
}
