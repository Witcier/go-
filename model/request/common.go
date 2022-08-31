package request

type List struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Keyword  string `json:"keyword"`
}
