package service

import (
	"context"
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/core"
)

// AuthSingUP
func (s *Service) AuthSingUP(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.AuthSingUP - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.AuthSingUP(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.AuthSingUP - reply: %+v", r)
	return r, err
}

// AuthSingIN
func (s *Service) AuthSingIN(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.AuthSingIN - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.AuthSingIN(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.AuthSingUP - reply: %+v", r)
	return r, err
}

// Confirmation
func (s *Service) Confirmation(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.Confirmation(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}

func (s *Service) ForgotPassword(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.ForgotPassword(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}

func (s *Service) ForgotPasswordConfirm(ctx context.Context, in json.RawMessage) (interface{}, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), string(in))

	r, err := c.ForgotPasswordConfirm(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}
