package service

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
)

type DepartmentService struct{}

func (departmentService *DepartmentService) ListDepartment(r request.List) (response.PageResult, error) {
	var resp response.PageResult
	var total int64

	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.DB.Model(&model.Department{})
	var departments []model.Department
	err := db.Count(&total).Error
	if err != nil {
		return resp, err
	}

	err = db.Limit(limit).Offset(offset).Find(&departments).Error

	resp = response.PageResult{
		List:     departments,
		Total:    total,
		Page:     r.Page,
		PageSize: r.PageSize,
	}

	return resp, err
}

func (departmentService *DepartmentService) StoreDepartment(r request.StoreDepartment) error {
	department := &model.Department{
		Name:   r.Name,
		Remark: r.Remark,
	}

	err := global.DB.Create(&department).Error

	return err
}

func (departmentService *DepartmentService) UpdateDepartment(id uint, r request.UpdateDepartment) error {
	department := &model.Department{
		ModelID: global.ModelID{
			ID: id,
		},
		Name:   r.Name,
		Remark: r.Remark,
	}

	err := global.DB.Updates(&department).Error

	return err
}

func (departmentService *DepartmentService) DeleteDepartment(id uint) error {
	var department model.Department
	err := global.DB.Where("id = ?", id).Delete(&department).Error
	if err != nil {
		return err
	}

	return nil
}
