package users

import (
	"context"
	"database/sql"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/mapper/users"
	"time"
)

var (
	ErrUserDuplicate = users.ErrDuplicateEmail
	ErrUserNotFound  = users.ErrUserNotFound
)

type AuthRepository struct {
	mapper users.UserAuthMapper
}

func NewUserInfoRepository(mapper users.UserAuthMapper) *AuthRepository {
	return &AuthRepository{
		mapper: mapper,
	}
}

func (repo *AuthRepository) FindByUsername(ctx context.Context, username string) (model.Auth, error) {
	u, err := repo.mapper.FindByUsername(ctx, username)
	if err != nil {
		return model.Auth{}, err
	}
	return repo.entityToDomain(u), nil
}

func (repo *AuthRepository) Create(ctx context.Context, auth model.Auth) error {
	return repo.mapper.Insert(ctx, repo.domainToEntity(auth))
}

func (repo *AuthRepository) domainToEntity(u model.Auth) users.Auth {
	return users.Auth{
		Id: u.Id,
		Username: sql.NullString{
			String: u.Username,
			Valid:  u.Username != "",
		},
		Password:   u.Password,
		CreateTime: u.CreateTime.UnixMilli(),
	}
}

func (repo *AuthRepository) entityToDomain(a users.Auth) model.Auth {
	return model.Auth{
		Id:         a.Id,
		Username:   a.Username.String,
		Password:   a.Password,
		CreateTime: time.UnixMilli(a.CreateTime),
	}
}
