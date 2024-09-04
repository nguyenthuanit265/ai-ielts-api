package models

type User struct {
	UserId   int    `json:"juser_id" db:"user_id"`
	UserName string `json:"juser_name" db:"user_name" binding:"required"`
	Phone    string `json:"jphone" db:"phone" binding:"required"`
}

type UserInsertBatchReq struct {
	Data []User `json:"data"`
}
