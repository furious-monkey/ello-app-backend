package core

import (
	"errors"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/mail"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/numbers"
	"gorm.io/gorm"
	"time"
)

func (c *AccountCore) AccountConfirmationSend(in *account.ConfirmationSendReq) (*account.ConfirmationSendResp, error) {
	// todo: add your logic here and delete this line
	var userEllo models.UsersEllo

	code, err := numbers.ConfirmationCode(6)
	if err != nil {
		c.Logger.Errorf("can not generate confirmation code (%v)", err)
		return &account.ConfirmationSendResp{
			Status:             false,
			Email:              in.Email,
			ConfirmationExpire: 0,
			Message:            "Can not generate confirmation code",
		}, err
	}

	if err := c.svcCtx.Gorm.Where("email", in.Email).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err := fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return &account.ConfirmationSendResp{
			Status:             false,
			Email:              in.Email,
			ConfirmationExpire: 0,
			Message:            "Can not get usersEllo record",
		}, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("there is no account in usersEllo")
		c.Logger.Error(err)
		return &account.ConfirmationSendResp{
			Status:             false,
			Email:              in.Email,
			ConfirmationExpire: 0,
			Message:            "There is not account in usersEllo",
		}, err
	}

	exp := time.Now().Add(time.Minute * 10)
	confirmationCodes := &models.ConfirmationCodes{
		UserID:    userEllo.UserID,
		Code:      code,
		ExpiredAt: &exp,
	}

	if err := c.svcCtx.Gorm.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(confirmationCodes).Error; err != nil {
			c.Logger.Errorf("can not create confirmation_codes record (%v)", err)
			return err
		}

		mailReq := &mail.SendMailRequest{
			Email:    in.Email,
			Username: userEllo.Username,
			Message:  "Confirmation code: " + code,
			Subject:  "Confirmation code from ElloApp",
		}

		if _, err = mail.SendMail(c.ctx, mailReq); err != nil {
			c.Logger.Errorf("can not send code to mail (%v)", err)
			return err
		}

		return
	}); err != nil {
		return nil, err
	}

	return &account.ConfirmationSendResp{
		Status:             true,
		Email:              in.Email,
		ConfirmationExpire: confirmationCodes.ExpiredAt.Unix(),
		Message:            "Sent Successfully",
	}, nil
}
