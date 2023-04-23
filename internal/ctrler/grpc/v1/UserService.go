package v1

import (
	"context"

	genv1 "demo/gen/user/v1"
	"demo/internal/svc"

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

func init() {
	UserService = newUserService()
}

func (us *userService) RetrieveDemoUser(
	ctx context.Context,
	req *genv1.RetrieveDemoUserRequest) (
	*genv1.RetrieveDemoUserResponse,
	error,
) {
	u, err := svc.UserService.RetrieveDemoUser(ctx)
	if err != nil {
		us.logger.Error().Err(err).Send()
		return &genv1.RetrieveDemoUserResponse{}, err
	}

	resp := &genv1.RetrieveDemoUserResponse{
		User: &genv1.User{
			Name: u.Name,
			Age:  uint32(u.Age),
		},
	}
	return resp, nil
}
