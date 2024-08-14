package sitemodel

type UserLoginRequest struct {
	Username *string `json:"username" binding:"required,min=1"`
	Password *string `json:"password" binding:"required,min=1"`
}
