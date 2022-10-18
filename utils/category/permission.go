package category

import "witcier/go-api/model"

type PermissionList struct {
	model.Permission
	Child []PermissionList `json:"child"`
}

func PermissionCategory(data []model.Permission, pid uint) []PermissionList {
	var list []PermissionList
	for _, v := range data {
		if v.ParentID == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := PermissionCategory(data, v.ID)
			node := PermissionList{
				Permission: v,
				Child:      child,
			}
			list = append(list, node)
		}
	}

	return list
}
