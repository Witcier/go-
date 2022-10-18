package service

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils/category"
)

type PermissionService struct{}

func (s *PermissionService) List() (response.CategoryResult, error) {
	var permissions []model.Permission
	var resp response.CategoryResult

	err := global.DB.Model(&model.Permission{}).Find(&permissions).Error
	if err != nil {
		return resp, err
	}

	list := category.PermissionCategory(permissions, 0)
	resp = response.CategoryResult{
		List: list,
	}

	return resp, nil
}

func (s *PermissionService) Store(r request.StorePermission) error {
	permission := &model.Permission{
		Category: global.Category{
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Rule: r.Rule,
	}

	err := global.DB.Create(&permission).Error

	return err
}

func (s *PermissionService) Update(id uint, r request.UpdatePermission) error {
	permission := &model.Permission{
		Category: global.Category{
			ModelID: global.ModelID{
				ID: id,
			},
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Rule: r.Rule,
	}

	err := global.DB.Updates(&permission).Error

	return err
}

func (s *PermissionService) Delete(id uint) error {
	var permission model.Permission
	err := global.DB.Where("id = ?", id).Delete(&permission).Error
	if err != nil {
		return err
	}

	return nil
}
