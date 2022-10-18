package category

import (
	"witcier/go-api/model"
)

type PositionList struct {
	model.Position
	Child []PositionList `json:"child"`
}

func PositionCategory(data []model.Position, pid uint) []PositionList {
	var list []PositionList
	for _, v := range data {
		if v.ParentID == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := PositionCategory(data, v.ID)
			node := PositionList{
				Position: v,
				Child:    child,
			}
			list = append(list, node)
		}
	}

	return list
}
