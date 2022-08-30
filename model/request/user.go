package request

type StoreUser struct {
	Username     string `json:"username" validate:"required"`
	RealName     string `json:"realName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required"`
	Sex          int8   `json:"sex" validate:"required,oneof=0 1 2"`
	Birth        int    `json:"birth" validate:"required"`
	EntryTime    int    `json:"entryTime" validate:"required"`
	Status       bool   `json:"status" validate:"required"`
	IsAdmin      bool   `json:"isAdmin" validate:"required"`
	DepartmentID uint   `json:"departmentID" validate:"required"`
	PositionID   uint   `json:"positionID" validate:"required"`
}

type UpdateUser struct {
	Username     string `json:"username"`
	RealName     string `json:"realName"`
	Email        string `json:"email" validate:"email"`
	Phone        string `json:"phone"`
	Sex          int8   `json:"sex" validate:"oneof=0 1 2"`
	Birth        int    `json:"birth"`
	EntryTime    int    `json:"entryTime"`
	Status       bool   `json:"status"`
	IsAdmin      bool   `json:"isAdmin"`
	DepartmentID uint   `json:"departmentID"`
	PositionID   uint   `json:"positionID"`
}
