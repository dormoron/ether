package users

import (
	"context"
	"gorm.io/gorm"
)

type RoleMapper interface {
	Insert(ctx context.Context, r Role) (Role, error)
	FirstOrCreateByDefaultInTxn(ctx context.Context, tx *gorm.DB, r Role) (uint, error)
	UpdateById(ctx context.Context, r Role) error
	Disable(ctx context.Context, r Role) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (Role, error)
	FindByDisable(ctx context.Context, disable bool) ([]Role, error)
}

type RoleMapperStruct struct {
	db *gorm.DB
}

func (mapper *RoleMapperStruct) Delete(ctx context.Context, id uint) error {
	var role Role
	return mapper.db.WithContext(ctx).Delete(&role, id).Error
}

func (mapper *RoleMapperStruct) FirstOrCreateByDefaultInTxn(ctx context.Context, tx *gorm.DB, r Role) (uint, error) {
	var role Role
	err := tx.WithContext(ctx).Where("is_default = ?", true).Assign(r).FirstOrCreate(&role).Error
	return role.ID, err
}

func (mapper *RoleMapperStruct) Insert(ctx context.Context, r Role) (Role, error) {
	err := mapper.db.WithContext(ctx).Create(&r).Error
	return r, err
}

func (mapper *RoleMapperStruct) UpdateById(ctx context.Context, r Role) error {
	res := mapper.db.Model(&r).WithContext(ctx).Where("id = ?", r.ID).Updates(
		map[string]any{
			"Name":        r.Name,
			"Description": r.Description,
		})
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return ErrUpdateData
	}
	return nil
}

func (mapper *RoleMapperStruct) Disable(ctx context.Context, r Role) error {
	res := mapper.db.Model(&r).WithContext(ctx).
		Where("id =?", r.ID).
		Updates(map[string]any{
			"isDisable": r.IsDisable,
		})
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return ErrUpdateData
	}
	return nil
}

func (mapper *RoleMapperStruct) FindById(ctx context.Context, id uint) (Role, error) {
	var r Role
	err := mapper.db.WithContext(ctx).Where("`id`=?", id).First(&r).Error
	return r, err
}

func (mapper *RoleMapperStruct) FindByDisable(ctx context.Context, disable bool) ([]Role, error) {
	var r []Role
	err := mapper.db.WithContext(ctx).Where("is_disable=?", disable).Find(&r).Error
	return r, err
}

func NewRoleMapper(db *gorm.DB) RoleMapper {
	return &RoleMapperStruct{
		db: db,
	}
}
