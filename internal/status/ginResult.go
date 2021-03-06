package status

var (
	OK  = response(200, "OK")    // 通用成功
	Err = response(500, "ERROR") // 通用错误

	ParamErr           = response(1001, "获取参数失败")
	CreateErr          = response(1002, "创建礼包码失败")
	CodeLenErr         = response(1003, "礼包码输入错误")
	FindCodeErr        = response(1004, "查询礼品码失败")
	VerifyCodeErr      = response(1005, "礼品码验证失败")
	CodeTypeErr        = response(1006, "礼品码类型错误")
	CodeUserErr        = response(1007, "请输用户名")
	RedisErr           = response(1008, "redis异常")
	MarshalErr         = response(1009, "序列化异常")
	Received           = response(1010, "礼包码已领取结束")
	DesignatedUser     = response(1011, "指定用户领取")
	DesignatedReceived = response(1012, "您已领取，不要重复领取")
	CodeTimeOver       = response(1013, "礼品码过期或者无效")
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息

func (res *Response) WithMsg(message string, err error) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
