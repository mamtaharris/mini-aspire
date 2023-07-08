package requests

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
