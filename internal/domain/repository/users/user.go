package users

import (
	"context"
	"database/sql"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/cache"
	"ether/internal/domain/repository/mapper/users"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateInTxn(ctx context.Context, tx *gorm.DB, auth model.Auth) (uint, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindById(ctx context.Context, id uint) (model.User, error)
	domainToEntity(u model.Auth) users.User
	entityToDomain(u users.User) model.User
}

type UserRepositoryStruct struct {
	mapper users.UserMapper
	cache  cache.UserCache
}

func NewUserRepository(mapper users.UserMapper, c cache.UserCache) UserRepository {
	return &UserRepositoryStruct{
		mapper: mapper,
		cache:  c,
	}
}

func (repo *UserRepositoryStruct) CreateInTxn(ctx context.Context, tx *gorm.DB, auth model.Auth) (uint, error) {
	return repo.mapper.InsertInTxn(ctx, tx, repo.domainToEntity(auth))
}

func (repo *UserRepositoryStruct) FindByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := repo.mapper.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return repo.entityToDomain(u), nil
}

func (repo *UserRepositoryStruct) FindById(ctx context.Context, id uint) (model.User, error) {
	u, err := repo.cache.Get(ctx, id)
	if err == nil {
		return u, nil
	}
	user, err := repo.mapper.FindById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	u = repo.entityToDomain(user)

	go func() {
		_ = repo.cache.Set(ctx, u)
	}()
	return u, nil

}

func (repo *UserRepositoryStruct) domainToEntity(u model.Auth) users.User {
	return users.User{
		Email: sql.NullString{
			String: u.Username,
			Valid:  u.Username != "",
		},
	}
}

func (repo *UserRepositoryStruct) entityToDomain(u users.User) model.User {
	return model.User{
		Id:         u.ID,
		Username:   u.Email.String,
		WebSite:    u.WebSite,
		Nickname:   u.Nickname,
		Avatar:     u.Avatar,
		Intro:      u.Intro,
		IsDisable:  u.IsDisable,
		CreateTime: u.CreatedAt,
	}
}
