package models

// ParamSignUp 注册接口的参数
type ParamSignUp struct {
	StudentCard string `json:"student_card" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登陆接口的参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamSetUp 设置用户信息接口的参数
type ParamSetUp struct {
	StudentCard string `json:"student_card"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}
