// Code generated by goctl. DO NOT EDIT.
// Source: channels.proto

package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/core"
)

// NEW
func (s *Service) GetChannelDataById(ctx context.Context, in *channels.ChannelDataByIdReq) (*channels.ChannelData, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelDataById - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelDataById(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelDataById - reply: %+v", r)
	return r, err
}

func (s *Service) CreateNewChannel(ctx context.Context, in *channels.CreateNewChannelReq) (*channels.ChannelData, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.CreateNewChannel - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.CreateNewChannel(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.CreateNewChannel - reply: %+v", r)
	return r, err
}

func (s *Service) AddChannelParticipant(ctx context.Context, in *channels.AddChannelParticipantReq) (*channels.AddChannelParticipantResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.AddChannelParticipant - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.AddChannelParticipant(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.AddChannelParticipant - reply: %+v", r)
	return r, err
}

func (s *Service) GetChatsListBySelfAndIDList(ctx context.Context, in *channels.GetChatsListBySelfAndIDListReq) (*channels.GetChatsListBySelfAndIDListResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChatsListBySelfAndIDList - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChatsListBySelfAndIDList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChatsListBySelfAndIDList - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelFullBySelfId(ctx context.Context, in *channels.GetChannelFullBySelfIdReq) (*channels.GetChannelFullBySelfIdResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelFullBySelfId - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelFullBySelfId(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelFullBySelfId - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelParticipantList(ctx context.Context, in *channels.ChannelParticipantListReq) (*channels.ChannelParticipantListResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelParticipantList - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelParticipantList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelParticipantList - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelParticipants(ctx context.Context, in *channels.ChannelParticipantsReq) (*channels.GetChannelParticipantsResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelParticipants - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelParticipants(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelParticipants - reply: %+v", r)
	return r, err
}

func (s *Service) ToChat(ctx context.Context, in *channels.ToChatReq) (*channels.ToChatResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.ToChat - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.ToChat(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.ToChat - reply: %+v", r)
	return r, err
}

func (s *Service) CheckChannelUserName(ctx context.Context, in *channels.CheckChannelUserNameReq) (*channels.CheckChannelUserNameResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.CheckChannelUserName - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.CheckChannelUserName(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.CheckChannelUserName - reply: %+v", r)
	return r, err
}

func (s *Service) UpdateChannelLink(ctx context.Context, in *channels.UpdateChannelLinkReq) (*channels.UpdateChannelLinkResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.UpdateChannelLink - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.UpdateChannelLink(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.UpdateChannelLink - reply: %+v", r)
	return r, err
}

func (s *Service) CheckUserIsAdministrator(ctx context.Context, in *channels.CheckUserIsAdministratorReq) (*channels.CheckUserIsAdministratorResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.CheckUserIsAdministrator - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.CheckUserIsAdministrator(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.CheckUserIsAdministrator - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelParticipantIdList(ctx context.Context, in *channels.ChannelData) (*channels.GetChannelParticipantIdListResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelParticipantIdList - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelParticipantIdList(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelParticipantIdList - reply: %+v", r)
	return r, err
}

func (s *Service) CheckDeleteChannelUser(ctx context.Context, in *channels.CheckDeleteChannelUserReq) (*channels.CheckDeleteChannelUserResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.CheckDeleteChannelUser - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.CheckDeleteChannelUser(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.CheckDeleteChannelUser - reply: %+v", r)
	return r, err
}

func (s *Service) DeleteChannelUser(ctx context.Context, in *channels.DeleteChannelUserReq) (*channels.DeleteChannelUserResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.DeleteChannelUser - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.DeleteChannelUser(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.DeleteChannelUser - reply: %+v", r)
	return r, err
}

func (s *Service) EditChannelTitle(ctx context.Context, in *channels.EditChannelTitleReq) (*channels.EditChannelTitleResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.EditChannelTitle - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.EditChannelTitle(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.EditChannelTitle - reply: %+v", r)
	return r, err
}

func (s *Service) EditChannelPhoto(ctx context.Context, in *channels.EditChannelPhotoReq) (*channels.EditChannelPhotoResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.EditChannelPhoto - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.EditChannelPhoto(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.EditChannelPhoto - reply: %+v", r)
	return r, err
}

func (s *Service) EditChannelAdmin(ctx context.Context, in *channels.EditChannelAdminReq) (*channels.EditChannelAdminResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.EditChannelAdmin - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.EditChannelAdmin(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.EditChannelAdmin - reply: %+v", r)
	return r, err
}

func (s *Service) ToggleChannelAdmins(ctx context.Context, in *channels.ToggleChannelAdminsReq) (*channels.ToggleChannelAdminsResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.ToggleChannelAdmins - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.ToggleChannelAdmins(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.ToggleChannelAdmins - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelBySelfID(ctx context.Context, in *channels.GetChannelBySelfIDReq) (*channels.GetChannelBySelfIDResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelBySelfID - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelBySelfID(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelBySelfID - reply: %+v", r)
	return r, err
}

func (s *Service) GetChannelParticipantIdListDao(ctx context.Context, in *channels.GetChannelParticipantIdListDaoReq) (*channels.GetChannelParticipantIdListDaoResp, error) {
	c := core.NewChannels(ctx, s.svcCtx)

	c.Logger.Debugf("channels.GetChannelParticipantIdListDao - metadata: %s, request: %s", c.MD.DebugString(), in.String())
	r, err := c.GetChannelParticipantIdListDao(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.GetChannelParticipantIdListDao - reply: %+v", r)
	return r, err
}
