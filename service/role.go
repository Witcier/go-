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
