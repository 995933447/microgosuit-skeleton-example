package handler

import (
	"context"

	"github.com/995933447/microgosuit-skeleton-example/user"
)

func (s *UserGateway) GetUserInfo(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	var resp user.GetUserResp

	getUserResp, err := user.UserRpc().GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.User = getUserResp.User

	return &resp, nil
}
