package session

import (
	"context"
	"demo/internal/consts"
	"demo/internal/model/entity"
	"demo/internal/service"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}
func New() *sSession {
	return &sSession{}
}

func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.UserSessionKey, user)
}

func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}
