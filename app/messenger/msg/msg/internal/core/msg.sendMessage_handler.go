package core

import (
	"context"
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MsgSendMessage
// msg.sendMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;
func (c *MsgCore) MsgSendMessage(in *msg.TLMsgSendMessage) (*mtproto.Updates, error) {
	var (
		rUpdates *mtproto.Updates
		err      error
		outBox   = in.GetMessage()
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	if outBox.GetScheduleDate().GetValue() != 0 {
		// c.Logger.Errorf("msg.sendMessage blocked, License key from https://elloapp.com required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if !peer.IsUserOrChatOrChannel() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMessage - error: %v", err)
		return nil, err
	}

	if peer.IsUser() {
		rUpdates, err = c.sendUserOutgoingMessage(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		if err != nil {
			c.Logger.Errorf("msg.sendMessage - error: %v", err)
			return nil, err
		}
	} else if peer.IsChat() {
		rUpdates, err = c.sendChatOutgoingMessage(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		if err != nil {
			c.Logger.Errorf("msg.sendMessage - error: %v", err)
			return nil, err
		}
	} else if peer.IsChannel() {
		var forbidden = true
		if in.Message.Message != nil && in.Message.Message.Action != nil {
			switch in.Message.Message.Action.PredicateName {
			case mtproto.Predicate_messageActionChatDeleteUser:
				forbidden = false
			default:
			}
		}
		if forbidden {
			rOk, err := c.svcCtx.Dao.ChannelsClient.CheckUserIsAdministrator(c.ctx, &channels.CheckUserIsAdministratorReq{
				ChannelId: in.PeerId,
				UserId:    in.UserId,
			})
			if err != nil {
				return nil, err
			}
			forbidden = !rOk.Status
		}

		if forbidden {
			return nil, errors.New("NO_EDIT_CHAT_PERMISSION")
		}

		rUpdates, err = c.sendChannelOutgoingMessage(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		if err != nil {
			c.Logger.Errorf("msg.sendMessage - error: %v", err)
			return nil, err
		}
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserOutgoingMessage(userId, authKeyId, peerUserId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	var (
		err      error
		rUpdates *mtproto.Updates
	)

	rUpdates, err = c.sendUserMessage(
		userId,
		authKeyId,
		peerUserId,
		outBox,
		func(did int64, inboxMsg *mtproto.Message) error {
			//inBox := &msgpb.InboxUserMessage{
			//	From:            makeSender(r.From),
			//	PeerUserId:      r.PeerId,
			//	RandomId:        r.Message.RandomId,
			//	DialogMessageId: did,
			//	MessageDataId:   mid,
			//	Message:         inboxMsg,
			//}
			//if model.IsBotFather(r.PeerUserId) {
			//	return s.botsClient.SendUserMessageToInbox(ctx, inBox)
			//} else {
			// toUser, _ := s.UserFacade.GetUserById(ctx, r.From.Id, r.PeerId)
			// log.Debug(toUser.DebugString())
			blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(c.ctx, &userpb.TLUserBlockedByUser{
				UserId:     peerUserId,
				PeerUserId: userId,
			})
			// UserIsBlockedByUser(ctx, r.PeerId, r.From.Id)
			if !mtproto.FromBool(blocked) {
				_, err = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInbox(c.ctx, &inbox.TLInboxSendUserMessageToInbox{
					FromId:     userId,
					PeerUserId: peerUserId,
					Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
						RandomId:        outBox.RandomId,
						DialogMessageId: did,
						// MessageDataId:   mid,
						Message: inboxMsg,
					}).To_InboxMessageData(),
				})
			}
			//}
			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserMessage(
	fromUserId int64,
	fromAuthKeyId int64,
	toUserId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {
	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{fromUserId, toUserId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sender, _ := users.GetImmutableUser(fromUserId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(toUserId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sendMe := fromUserId == toUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	// handle duplicateMessage
	hasDuplicateMessage, err := c.svcCtx.Dao.HasDuplicateMessage(c.ctx, fromUserId, outBox.RandomId)
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(c.ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	box, err := c.svcCtx.Dao.SendUserMessage(c.ctx, fromUserId, toUserId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			// TODO: check
			//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
			//	&userpb.TLUserGetMutableUsers{
			//		Id: idList,
			//	})
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			if len(idList) > 0 {
				chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
					&chatpb.TLChatGetChatListByIdList{
						IdList: idList,
					})
				return chats.GetChatListByIdList(fromUserId, idList...)
			} else {
				return []*mtproto.Chat{}
			}
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateNewMessage)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    fromUserId,
		AuthKeyId: fromAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updateNewMessage),
	})

	return rUpdates, nil
}

func (c *MsgCore) sendChatOutgoingMessage(userId, authKeyId, peerChatId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	rUpdates, err := c.sendChatMessage(c.ctx,
		userId,
		authKeyId,
		peerChatId,
		outBox,
		func(did int64, inboxMsg *mtproto.Message) error {
			_, err := c.svcCtx.Dao.InboxClient.InboxSendChatMessageToInbox(
				c.ctx,
				&inbox.TLInboxSendChatMessageToInbox{
					FromId:     userId,
					PeerChatId: peerChatId,
					Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
						RandomId:        outBox.RandomId,
						DialogMessageId: did,
						// MessageDataId:   mid,
						Message: inboxMsg,
					}).To_InboxMessageData(),
				})
			if err != nil {
				c.Logger.Errorf("checkDuplicateMessage error - %v", err)
				return err
			}

			return err
		})
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendChatMessage(
	ctx context.Context,
	fromUserId int64,
	fromAuthKeyId int64,
	chatId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {

	hasDuplicateMessage, err := c.svcCtx.Dao.HasDuplicateMessage(ctx, fromUserId, outBox.RandomId)
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	box, err := c.svcCtx.Dao.SendChatMessage(ctx, fromUserId, chatId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateNewMessage)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    fromUserId,
		AuthKeyId: fromAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updateNewMessage),
	})

	c.svcCtx.Dao.PutDuplicateMessage(ctx, fromUserId, outBox.RandomId, rUpdates)

	return rUpdates, nil
}

func (c *MsgCore) sendChannelOutgoingMessage(userId, authKeyId, peerChatId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	rUpdates, err := c.sendChannelMessage(c.ctx,
		userId,
		authKeyId,
		peerChatId,
		outBox,
		func(did int64, inboxMsg *mtproto.Message) error {
			_, err := c.svcCtx.Dao.InboxClient.InboxSendChatMessageToInbox(
				c.ctx,
				&inbox.TLInboxSendChatMessageToInbox{
					FromId:     userId,
					PeerChatId: peerChatId,
					Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
						RandomId:        outBox.RandomId,
						DialogMessageId: did,
						// MessageDataId:   mid,
						Message: inboxMsg,
					}).To_InboxMessageData(),
				})
			if err != nil {
				c.Logger.Errorf("checkDuplicateMessage error - %v", err)
				return err
			}

			return err
		})
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendChannelMessage(
	ctx context.Context,
	fromUserId int64,
	fromAuthKeyId int64,
	chatId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {

	hasDuplicateMessage, err := c.svcCtx.Dao.HasDuplicateMessage(ctx, fromUserId, outBox.RandomId)
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	box, err := c.svcCtx.Dao.SendChannelMessage(ctx, fromUserId, chatId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			res, _ := c.svcCtx.ChannelsClient.GetChatsListBySelfAndIDList(c.ctx, &channels.GetChatsListBySelfAndIDListReq{
				SelfUserId: fromUserId,
				IdList:     idList,
			})
			return res.Chats
		},
		updateNewMessage)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    fromUserId,
		AuthKeyId: fromAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updateNewMessage),
	})

	c.svcCtx.Dao.PutDuplicateMessage(ctx, fromUserId, outBox.RandomId, rUpdates)

	return rUpdates, nil
}
