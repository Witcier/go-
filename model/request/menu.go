package request

type StoreMenu struct {
	Name     string `json:"name" binding:"required"`
	Uri      string `json:"uri"`
	ParentID uint   `json:"parentID"`
	Order    int    `json:"order"`
}

type UpdateMenu struct {
	Name     string `json:"name"`
	Uri      string `json:"uri"`
	ParentID uint   `json:"parentID"`
	Order    int    `json:"order"`
}
