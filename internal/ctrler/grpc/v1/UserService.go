package v1

import (
	"context"

	v1 "demo/gen/user/v1"
	"demo/internal/svc"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	UserService *userService
)

type userService struct {
	logger zerolog.Logger
}

func newUserService() *userService {
	us := &userService{
		logger: log.With().Str("ltag", "userService").Logger(),
	}
	return us
}

func (us *userService) RetrieveDemoUser(
	ctx context.Context,
	req *connect_go.Request[v1.RetrieveDemoUserRequest]) (
	*connect_go.Response[v1.RetrieveDemoUserResponse],
	error) {
	u, err := svc.UserService.RetrieveDemoUser(ctx)
	if err != nil {
		us.logger.Error().Err(err).Send()
		return &connect_go.Response[v1.RetrieveDemoUserResponse]{}, nil
	}
	resp := connect_go.NewResponse(&v1.RetrieveDemoUserResponse{
		User: &v1.User{
			Name: u.Name,
			Age:  uint32(u.Age),
		},
	})

	return resp, nil
}
