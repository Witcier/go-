package category

import (
	"witcier/go-api/model"
)

type MenuList struct {
	model.Menu
	Child []MenuList `json:"child"`
}

func MenuCategory(data []model.Menu, pid uint) []MenuList {
	var list []MenuList
	for _, v := range data {
		if v.ParentID == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := MenuCategory(data, v.ID)
			node := MenuList{
				Menu:  v,
				Child: child,
			}
			list = append(list, node)
		}
	}

	return list
}
