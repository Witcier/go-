package request

type StoreDepartment struct {
	Name   string `json:"name" binding:"required"`
	Remark string `json:"remark" binding:"required"`
}

type UpdateDepartment struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}
