/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/tos/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// HelpGetTermsOfServiceUpdate
// help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
func (s *Service) HelpGetTermsOfServiceUpdate(ctx context.Context, request *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getTermsOfServiceUpdate - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetTermsOfServiceUpdate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getTermsOfServiceUpdate - reply: %s", r.DebugString())
	return r, err
}

// HelpAcceptTermsOfService
// help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
func (s *Service) HelpAcceptTermsOfService(ctx context.Context, request *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.acceptTermsOfService - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpAcceptTermsOfService(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.acceptTermsOfService - reply: %s", r.DebugString())
	return r, err
}
