package service

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"
)

type MenuService struct{}

func (menuService *MenuService) ListMenu() (response.CategoryResult, error) {
	var menus []global.Category
	var resp response.CategoryResult

	err := global.DB.Model(&model.Menu{}).Find(&menus).Error
	if err != nil {
		return resp, err
	}

	list := utils.Category(menus, 0)
	resp = response.CategoryResult{
		List: list,
	}

	return resp, nil
}

func (menuService *MenuService) StoreMenu(r request.StoreMenu) error {
	menu := &model.Menu{
		Category: global.Category{
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Uri:   r.Uri,
		Order: r.Order,
	}

	err := global.DB.Create(&menu).Error

	return err
}

func (menuService *MenuService) UpdateMenu(id uint, r request.UpdateMenu) error {
	menu := &model.Menu{
		Category: global.Category{
			ModelID: global.ModelID{
				ID: id,
			},
			Name:     r.Name,
			ParentID: r.ParentID,
		},
		Uri:   r.Uri,
		Order: r.Order,
	}

	err := global.DB.Updates(&menu).Error

	return err
}

func (menuService *MenuService) DeleteMenu(id uint) error {
	var menu model.Menu
	err := global.DB.Where("id = ?", id).Delete(&menu).Error
	if err != nil {
		return err
	}

	return nil
}
