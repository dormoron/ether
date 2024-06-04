package users

import (
	"context"
	"database/sql"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/mapper/users"
	"strconv"
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
	a, err := repo.mapper.FindByUsername(ctx, username)
	if err != nil {
		return model.Auth{}, err
	}
	return repo.entityToDomain(a), nil
}

func (repo *AuthRepository) Create(ctx context.Context, auth model.Auth) error {
	return repo.mapper.Insert(ctx, repo.domainToEntity(auth))
}

func (repo *AuthRepository) domainToEntity(a model.Auth) users.Auth {
	return users.Auth{
		Id: a.Id,
		UserId: sql.NullString{
			String: strconv.FormatInt(a.UserId, 10),
			Valid:  a.UserId != 0,
		},
		Username: sql.NullString{
			String: a.Username,
			Valid:  a.Username != "",
		},
		Password:   a.Password,
		CreateTime: a.CreateTime.UnixMilli(),
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
