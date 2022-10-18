package request

type StorePermission struct {
	Name     string `json:"name" binding:"required"`
	Rule     string `json:"rule" binding:"required"`
	ParentID uint   `json:"parentID"`
}

type UpdatePermission struct {
	Name     string `json:"name"`
	Rule     string `json:"rule"`
	ParentID uint   `json:"parentID"`
}
