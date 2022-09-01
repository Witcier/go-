package request

type StoreDepartment struct {
	Name   string `json:"name" validate:"required"`
	Remark string `json:"remark" validate:"required"`
}

type UpdateDepartment struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}
