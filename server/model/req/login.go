package req

// LoginReqParam 登录请求参数
type LoginReqParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}