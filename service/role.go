package service

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
)

type RoleService struct{}

func (s *RoleService) List(r request.List) (response.PageResult, error) {
	var roles []model.Role
	var resp response.PageResult
	var total int64

	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.DB.Model(&model.Role{})
	err := db.Count(&total).Error
	if err != nil {
		return resp, err
	}

	err = db.Limit(limit).Offset(offset).Find(&roles).Error

	resp = response.PageResult{
		List:     roles,
		Total:    total,
		Page:     r.Page,
		PageSize: r.PageSize,
	}

	return resp, err
}

func (s *RoleService) Store(r request.StoreRole) error {
	role := &model.Role{
		Name:   r.Name,
		Remark: r.Remark,
	}

	err := global.DB.Create(&role).Error

	return err
}

func (s *RoleService) Update(id uint, r request.UpdateRole) error {
	role := &model.Role{
		ModelID: global.ModelID{
			ID: id,
		},
		Name:   r.Name,
		Remark: r.Remark,
	}

	err := global.DB.Updates(&role).Error

	return err
}

func (s *RoleService) Delete(id uint) error {
	var role model.Role
	err := global.DB.Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *RoleService) Menu(id uint, r request.StoreRoleMenu) error {
	var role model.Role
	var menus []model.Menu
	err := global.DB.Where("ID = ?", id).First(&role).Error
	if err != nil {
		return err
	}

	err = global.DB.Find(&menus, r.MenuIDs).Error
	if err != nil {
		return err
	}

	global.DB.Model(&role).Association("Menus").Replace(&menus)

	return nil
}

func (s *RoleService) GetMenu(id uint) (response.RoleMenuResponse, error) {
	var role model.Role
	var menus []model.Menu
	var resp response.RoleMenuResponse
	var menuIds []int

	err := global.DB.Where("ID = ?", id).Preload("Menus").Find(&role).Error
	if err != nil {
		return resp, err
	}

	menus = role.Menus
	for _, v := range menus {
		menuIds = append(menuIds, int(v.ID))
	}

	resp.MenuIds = menuIds

	return resp, err
}

func (s *RoleService) Permission(id uint, r request.StoreRolePermission) error {
	var role model.Role
	var permissions []model.Permission
	err := global.DB.Where("ID = ?", id).First(&role).Error
	if err != nil {
		return err
	}

	err = global.DB.Find(&permissions, r.PermissionIds).Error
	if err != nil {
		return err
	}

	global.DB.Model(&role).Association("Permissions").Replace(&permissions)

	return nil
}

func (s *RoleService) GetPermission(id uint) (response.RolePermissionResponse, error) {
	var role model.Role
	var permissions []model.Permission
	var resp response.RolePermissionResponse
	var permissionIds []int

	err := global.DB.Where("ID = ?", id).Preload("Permissions").Find(&role).Error
	if err != nil {
		return resp, err
	}

	permissions = role.Permissions
	for _, v := range permissions {
		permissionIds = append(permissionIds, int(v.ID))
	}

	resp.PermissionIds = permissionIds

	return resp, err
}
