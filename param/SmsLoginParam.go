package param

//手机号+验证码 登陆时传参
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
