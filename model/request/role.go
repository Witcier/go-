package request

type StoreRole struct {
	Name   string `json:"name" binding:"required"`
	Remark string `json:"remark" binding:"required"`
}

type UpdateRole struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type StoreRoleMenu struct {
	MenuIDs []int `json:"menuIds"`
}

type StoreRolePermission struct {
	PermissionIds []int `json:"PermissionIds"`
}
