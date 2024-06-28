package users

import (
	"context"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/mapper/users"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FirstByDefaultInTxn(ctx context.Context, tx *gorm.DB, role model.Role) (uint, error)
	UpdateById(ctx context.Context, r model.Role) error
	Create(ctx context.Context, role model.Role) (model.Role, error)
	FindById(ctx context.Context, id uint) (model.Role, error)
	Delete(ctx context.Context, id uint) error
	domainToEntity(r model.Role) users.Role
	entityToDomain(r users.Role) model.Role
}

type RoleRepositoryStruct struct {
	mapper users.RoleMapper
}

func (c *RoleRepositoryStruct) Delete(ctx context.Context, id uint) error {
	return c.mapper.Delete(ctx, id)
}

func (c *RoleRepositoryStruct) FindById(ctx context.Context, id uint) (model.Role, error) {
	byId, err := c.mapper.FindById(ctx, id)
	if err != nil {
		return model.Role{}, err
	}
	return c.entityToDomain(byId), err
}

func (c *RoleRepositoryStruct) UpdateById(ctx context.Context, r model.Role) error {
	return c.mapper.UpdateById(ctx, c.domainToEntity(r))
}

func (c *RoleRepositoryStruct) FirstByDefaultInTxn(ctx context.Context, tx *gorm.DB, role model.Role) (uint, error) {
	return c.mapper.FirstOrCreateByDefaultInTxn(ctx, tx, c.domainToEntity(role))
}

func (c *RoleRepositoryStruct) Create(ctx context.Context, role model.Role) (model.Role, error) {
	insert, err := c.mapper.Insert(ctx, c.domainToEntity(role))
	if err != nil {
		return model.Role{}, err
	}
	return c.entityToDomain(insert), err
}

func (c *RoleRepositoryStruct) domainToEntity(r model.Role) users.Role {
	return users.Role{
		Name:        r.Name,
		Description: r.Description,
		IsDefault:   r.IsDefault,
		IsDisable:   r.IsDisable,
	}
}

func (c *RoleRepositoryStruct) entityToDomain(r users.Role) model.Role {
	return model.Role{
		Id:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		IsDefault:   r.IsDefault,
		IsDisable:   r.IsDisable,
	}
}

func NewRoleRepository(mapper users.RoleMapper) RoleRepository {
	return &RoleRepositoryStruct{
		mapper: mapper,
	}
}
