package users

import (
	"context"
	"database/sql"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/mapper/users"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = users.ErrDuplicateEmail
	ErrUserNotFound  = users.ErrUserNotFound
)

type AuthRepository interface {
	FindByUsername(ctx context.Context, username string) (model.Auth, error)
	CreateInTxn(ctx context.Context, tx *gorm.DB, auth model.Auth) error
	domainToEntity(a model.Auth) users.Auth
	entityToDomain(a users.Auth) model.Auth
}

type AuthRepositoryStruct struct {
	mapper users.AuthMapper
}

func NewAuthRepository(mapper users.AuthMapper) AuthRepository {
	return &AuthRepositoryStruct{
		mapper: mapper,
	}
}

func (repo *AuthRepositoryStruct) FindByUsername(ctx context.Context, username string) (model.Auth, error) {
	a, err := repo.mapper.FindByUsername(ctx, username)
	if err != nil {
		return model.Auth{}, err
	}
	return repo.entityToDomain(a), nil
}

func (repo *AuthRepositoryStruct) CreateInTxn(ctx context.Context, tx *gorm.DB, auth model.Auth) error {
	return repo.mapper.InsertInTxn(ctx, tx, repo.domainToEntity(auth))
}

func (repo *AuthRepositoryStruct) domainToEntity(a model.Auth) users.Auth {
	return users.Auth{
		UserId: a.UserId,
		Username: sql.NullString{
			String: a.Username,
			Valid:  a.Username != "",
		},
		Password:      a.Password,
		LastLoginTime: time.Now(),
	}
}

func (repo *AuthRepositoryStruct) entityToDomain(a users.Auth) model.Auth {
	return model.Auth{
		Id:         a.ID,
		UserId:     a.UserId,
		Username:   a.Username.String,
		Password:   a.Password,
		CreateTime: a.CreatedAt,
	}
}
