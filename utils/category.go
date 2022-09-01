package utils

import (
	"witcier/go-api/global"
)

type List struct {
	global.Category
	Child []List
}

func Category(data []global.Category, pid uint) []List {
	var list []List
	for _, v := range data {
		if v.ParentID == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := Category(data, v.ID)
			node := List{
				Category: global.Category{
					ModelID: global.ModelID{
						ID: v.ID,
					},
					Name:     v.Name,
					ParentID: v.ParentID,
				},
				Child: child,
			}
			list = append(list, node)
		}
	}

	return list
}
