package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/code/code"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/random2"
)

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (c *CodeCore) CodeCreatePhoneCode(in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.svcCtx.Dao.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("getCachePhoneCode - error: %v", err)
		err = mtproto.ErrInternelServerError
		return nil, err
	}
	if codeData == nil {
		codeData = &code.PhoneCodeTransaction{
			AuthKeyId:             in.AuthKeyId,
			Phone:                 in.Phone,
			SessionId:             in.SessionId,
			PhoneNumberRegistered: in.PhoneNumberRegistered,
			PhoneCode:             random2.RandomNumeric(5),
			PhoneCodeHash:         crypto.GenerateStringNonce(16),
			PhoneCodeExpired:      int32(time.Now().Unix() + 3*60),
			SentCodeType:          in.SentCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          in.NextCodeType,
			State:                 code.CodeStateSend, // model.CodeStateSent
		}

	} else if in.SessionId != codeData.SessionId {
		codeData.State = code.CodeStateSend
		codeData.SessionId = in.SessionId
	}

	//switch codeData.State {
	//case model.CodeStateSend:
	//	codeData.State = model.CodeStateSent
	//case model.CodeStateSent:
	//	codeData.State = model.CodeStateSent
	//default:
	//	// codeData = newCodeData()
	//}
	//
	//if err = c.Dao.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData); err != nil {
	//	log.Errorf("putCachePhoneCode - error: %v", err)
	//	err = mtproto.ErrInternelServerError
	//	return
	//}

	return codeData, nil
}