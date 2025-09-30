package handler

import (
	"github.com/995933447/microgosuit-skeleton-example/user"
	"google.golang.org/grpc"
)

func (s *User) GetUserInputStream(stream grpc.ClientStreamingServer[user.GetUserReq, user.GetUserResp]) error {
	return nil
}
