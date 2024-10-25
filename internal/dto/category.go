package dto

type CategoryQueryReq struct {
	Userid int64 `form:"userID"`
}

type CategoryAddReq struct {
	Name   string `json:"name"`
	UserID int64  `json:"userid"`
	Sort   *int8  `json:"sort"`
	Icon   string `json:"icon"`
}
