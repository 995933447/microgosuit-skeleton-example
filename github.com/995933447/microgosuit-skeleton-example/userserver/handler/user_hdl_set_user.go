package handler

import (
	"context"

	"github.com/995933447/microgosuit-skeleton-example/user"
)

func (s *User) SetUser(ctx context.Context, req *user.SetUserReq) (*user.SetUserResp, error) {
	var resp user.SetUserResp
	return &resp, nil
}
