package dto

type UserRegisterReq struct {
	Wallet string `json:"wallet"`
}

type UserQueryReq struct {
	Wallet string `form:"wallet"`
}
type UserLoginReq struct {
	*UserRegisterReq
}
type UserLogoutReq struct {
	Jwt string `json:"jwt"`
}
