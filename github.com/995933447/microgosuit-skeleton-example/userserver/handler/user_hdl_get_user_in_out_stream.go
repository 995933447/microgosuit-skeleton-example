package handler

import (
	"github.com/995933447/microgosuit-skeleton-example/user"
	"google.golang.org/grpc"
)

func (s *User) GetUserInOutStream(stream grpc.BidiStreamingServer[user.GetUserReq, user.GetUserResp]) error {
	return nil
}
