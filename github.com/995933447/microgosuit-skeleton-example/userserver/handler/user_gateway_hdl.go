package handler

import (
	"github.com/995933447/microgosuit-skeleton-example/user"
)

type UserGateway struct {
	user.UnimplementedUserGatewayServer
}
