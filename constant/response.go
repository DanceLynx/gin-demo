package constant

type ResponseCode int

const (
	SUCCESS  ResponseCode = 0   //成功
	CODE_404 ResponseCode= 404 // 页面未找到
	CODE_500 ResponseCode= 500 // 服务器错误
	CODE_403 ResponseCode= 403 // 访问被拒绝
	CODE_401 ResponseCode= 401 // 认证失败

	USER_LOGIN_FAILED ResponseCode = 1001 //用户信息不合法
	USER_NOT_EXISTS ResponseCode = 1002 //用户不存在
	USER_JWT_ERROR ResponseCode = 1003 //登录生成jwt token失败
	USER_VERIFY_FAILD ResponseCode = 1003 //jwt认证失败
	USER_JWT_PARSE_FAILD ResponseCode = 1004 //jwt解析失败
)
