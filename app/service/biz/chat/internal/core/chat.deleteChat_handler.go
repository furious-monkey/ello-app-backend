package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (c *ChatCore) ChatDeleteChat(in *chat.TLChatDeleteChat) (*chat.MutableChat, error) {
	mChat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}
	if mChat.Creator() != in.OperatorId {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	keys := []string{c.svcCtx.Dao.GetChatCacheKey(in.ChatId)}
	mChat.Walk(func(userId int64, participant *chat.ImmutableChatParticipant) error {
		keys = append(keys, c.svcCtx.Dao.GetChatParticipantCacheKey(participant.ChatId, participant.UserId))
		return nil
	})

	_, _, err = c.svcCtx.Dao.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				// kicked
				c.svcCtx.Dao.ChatParticipantsDAO.UpdateStateByChatIdTx(tx, mtproto.ChatMemberStateKicked, in.ChatId)
				c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, 0, in.ChatId)
				c.svcCtx.Dao.ChatsDAO.UpdateDeactivatedTx(tx, false, in.ChatId)
			})
			return 0, 0, tR.Err
		},
		keys...)
	if err != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	return mChat, nil
}
