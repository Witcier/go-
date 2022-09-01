package service

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"
)

type PositionService struct{}

func (s *PositionService) ListPosition() (response.CategoryResult, error) {
	var positions []global.Category
	var resp response.CategoryResult

	err := global.DB.Model(&model.Position{}).Find(&positions).Error
	if err != nil {
		return resp, err
	}

	list := utils.Category(positions, 0)
	resp = response.CategoryResult{
		List: list,
	}

	return resp, nil
}

func (s *PositionService) StorePosition(r request.StorePosition) error {
	position := &model.Position{
		Category: global.Category{
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Remark: r.Remark,
	}

	err := global.DB.Create(&position).Error

	return err
}

func (s *PositionService) UpdatePosition(id uint, r request.UpdatePosition) error {
	position := &model.Position{
		Category: global.Category{
			ModelID: global.ModelID{
				ID: id,
			},
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Remark: r.Remark,
	}

	err := global.DB.Updates(&position).Error

	return err
}

func (s *PositionService) DeletePosition(id uint) error {
	var position model.Position
	err := global.DB.Where("id = ?", id).Delete(&position).Error
	if err != nil {
		return err
	}

	return nil
}
