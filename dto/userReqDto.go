package dto

type AddUserReq struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AddUserContactsReqDto struct {
	UserId    int
	ContactId int
}
