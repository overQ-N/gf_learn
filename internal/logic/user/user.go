package user

import (
	"context"
	"demo/internal/model/entity"
	"demo/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}
