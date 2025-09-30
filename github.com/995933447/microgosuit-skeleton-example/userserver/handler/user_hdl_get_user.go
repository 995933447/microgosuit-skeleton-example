package handler

import (
	"context"

	"github.com/995933447/microgosuit-skeleton-example/user"
)

func (s *User) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	var resp user.GetUserResp
	resp.User = &user.GetUserResp_User{
		Name: "John Doe",
	}
	return &resp, nil
}
