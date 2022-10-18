package request

type PersonUpdatePassword struct {
	Password    string `json:"password" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}
