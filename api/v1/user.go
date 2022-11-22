package v1

import (
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}

type UserProfileRes struct {
	*entity.User
}

type UserSignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type UserSignUpRes struct{}
