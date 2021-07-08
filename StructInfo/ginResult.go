package StructInfo

type MesInfo struct {
	Msg  string // 错误描述
	ER   error
	Data interface{} // 返回数据
}
