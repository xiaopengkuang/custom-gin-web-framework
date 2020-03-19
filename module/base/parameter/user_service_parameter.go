package parameter

// 用户账号基本信息
type UserAccount struct {
	username string
	secret   string
}

// 用户登录参数
type UserLoginPara struct {
	UserAccount   // 用户账号信息
	client string // 登录客户端
}

// 用户注册参数
type UserRegisterPara struct {
	UserAccount   //用户账号信息
	client string //注册来源
	code   string //验证码
}

// 用户token 校验参数
type UserTokenCheckPara struct {
	token  string
	client string
}
