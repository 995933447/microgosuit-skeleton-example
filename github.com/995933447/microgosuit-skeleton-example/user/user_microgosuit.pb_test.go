package user

import (
	"context"
	"log"
	"testing"

	"github.com/995933447/microgosuit"
)

func TestGetUser(t *testing.T) {
	err := microgosuit.InitSuitWithGrpc(context.TODO(), "", "microgosuit", "microgosuit/")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := UserGatewayRpc().GetUserInfo(context.Background(), &GetUserReq{})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(resp)
}
