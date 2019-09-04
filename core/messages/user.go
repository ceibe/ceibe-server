package messages

type NewUserRequest struct {
	Username string `json:"username" validate:"min=4,max=16,required"`
	Password string `json:"password" validate:"len=64,required"`
}

type UserBasicInformation struct {
	Username string `json:"username"`
}