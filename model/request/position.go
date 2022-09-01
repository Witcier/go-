package request

type StorePosition struct {
	Name     string `json:"name" binding:"required"`
	Remark   string `json:"remark" binding:"required"`
	ParentID uint   `json:"parentID"`
}

type UpdatePosition struct {
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	ParentID uint   `json:"parentID"`
}
