package response

type RoleMenuResponse struct {
	MenuIds []int `json:"menusIds"`
}

type RolePermissionResponse struct {
	PermissionIds []int `json:"permissionIds"`
}
