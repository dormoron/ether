package users

import (
	"context"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/users"
)

type RoleService interface {
	Insert(ctx context.Context, role model.Role) (model.Role, error)
	UpdateById(ctx context.Context, a model.Role) error
	Delete(ctx context.Context, id uint) error
	Detail(ctx context.Context, id uint) (model.Role, error)
}

type RoleServiceStruct struct {
	repository users.RoleRepository
}

func (r *RoleServiceStruct) Delete(ctx context.Context, id uint) error {
	return r.Delete(ctx, id)
}

func (r *RoleServiceStruct) Insert(ctx context.Context, role model.Role) (model.Role, error) {
	return r.repository.Create(ctx, role)
}

func (r *RoleServiceStruct) UpdateById(ctx context.Context, a model.Role) error {
	return r.repository.UpdateById(ctx, a)
}

func (r *RoleServiceStruct) Detail(ctx context.Context, id uint) (model.Role, error) {
	return r.repository.FindById(ctx, id)
}

func NewRoleService(repository users.RoleRepository) RoleService {
	return &RoleServiceStruct{
		repository: repository,
	}
}
