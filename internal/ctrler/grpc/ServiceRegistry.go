// 在此注册grpc server
package grpc

import (
	genv1 "demo/gen/user/v1"
	grpcv1 "demo/internal/ctrler/grpc/v1"

	"github.com/why2go/gostarter/grpcstarter"
)

func RegisterGrpcService() {
	grpcstarter.RegisterService(&genv1.UserService_ServiceDesc, grpcv1.UserService)
}
