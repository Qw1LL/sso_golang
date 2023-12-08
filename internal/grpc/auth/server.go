package auth

import (
	ssov1 "github.com/Qw1LL/auth_protos/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
