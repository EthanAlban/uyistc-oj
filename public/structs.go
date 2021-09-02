package public

// Ret 返回信息序列化的结构体
type Ret struct {
	Code int
	Data interface{}
	Msg  string
}

// JudgeResult 接收判题结果的结构体
type JudgeResult struct {
	Code int
	Info string
}

// SysUsage 获取系统资源使用情况
type SysUsage struct {
	LoadAvg float64
	MemLoad float64
	Topic   string
}
