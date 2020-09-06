package constant

type ResponseCode int

const (
	SUCCESS  ResponseCode = 0   //成功
	CODE_404 ResponseCode= 404 // 页面未找到
	CODE_500 ResponseCode= 500 // 服务器错误
	CODE_403 ResponseCode= 403 // 访问被拒绝
	CODE_401 ResponseCode= 401 // 认证失败
)
