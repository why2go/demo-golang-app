package mapper

import (
	"context"
	"demo/internal/domain"
)

var (
	UserMapper *userMapper
)

type userMapper struct {
}

func newUserMapper() *userMapper {
	return &userMapper{}
}

func init() {
	UserMapper = newUserMapper()
}

func (mapper *userMapper) FindOne(ctx context.Context) (*domain.User, error) {
	return &domain.User{Name: "Alice", Age: 25}, nil
}
