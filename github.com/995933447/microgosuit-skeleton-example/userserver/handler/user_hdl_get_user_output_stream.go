package handler

import (
	"github.com/995933447/microgosuit-skeleton-example/user"
	"google.golang.org/grpc"
)

func (s *User) GetUserOutputStream(req *user.GetUserReq, stream grpc.ServerStreamingServer[user.GetUserResp]) error {
	return nil
}
