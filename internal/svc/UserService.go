package svc

import (
	"context"
	"demo/internal/domain"
	"demo/internal/mapper"

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
	s := &userService{
		logger: log.With().Str("ltag", "userService").Logger(),
	}
	return s
}

func init() {
	UserService = newUserService()
}

func (svc *userService) RetrieveDemoUser(ctx context.Context) (*domain.User, error) {
	logger := svc.logger.With().
		Str("requestId", ctx.Value("request-id").(string)).
		Str("func", "RetrieveDemoUser").
		Logger()
	u, err := mapper.UserMapper.FindOne(ctx)
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, err
	}
	return u, nil
}
