package users

import (
	"context"
	"database/sql"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/mapper/users"
	"time"
)

type UserRepository struct {
	mapper users.UserMapper
}

func NewUserRepository(mapper users.UserMapper) *UserRepository {
	return &UserRepository{mapper}
}

func (repo *UserRepository) Create(ctx context.Context, auth model.Auth) (int64, error) {
	return repo.mapper.Insert(ctx, repo.domainToEntity(auth))
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := repo.mapper.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return repo.entityToDomain(u), nil
}

func (repo *UserRepository) domainToEntity(u model.Auth) users.User {
	return users.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Username,
			Valid:  u.Username != "",
		},
		CreateTime: u.CreateTime.UnixMilli(),
	}
}

func (repo *UserRepository) entityToDomain(u users.User) model.User {
	return model.User{
		Id:         u.Id,
		Username:   u.Email.String,
		WebSite:    u.WebSite,
		Nickname:   u.Nickname,
		Avatar:     u.Avatar,
		Intro:      u.Intro,
		IsDisable:  u.IsDisable,
		CreateTime: time.UnixMilli(u.CreateTime),
	}
}
