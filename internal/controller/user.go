package controller

import (
	"context"

	v1 "demo/api/v1"
	"demo/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res = &v1.UserProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}
