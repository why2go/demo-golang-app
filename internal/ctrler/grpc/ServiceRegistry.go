// 在此注册grpc server
package grpc

import (
	"demo/gen/user/v1/userv1connect"

	"github.com/why2go/gostarter/grpcstarter"
)

func RegisterGrpcService() {
	grpcstarter.RegisterService(userv1connect.UserServiceHandler)
}
