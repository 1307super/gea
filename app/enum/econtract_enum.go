package enum

const (
	// ============== 刘平福接口 =============
	// 生成合同
	Creatgea string = "/token/contract/savgea"
	// 发起合同
	StartContract string = "/token/contract/startContract"

	//  ============ 金格接口 ================
	// 发送短信
	ToSignSms string = "/api/common/sendSignContentMsg"
	// 合同作废
	ToSignToVoid string = "/api/contract/toVoid"
	// 查询合同信息
	ToSignQueryContract string = "/api/cntquery"

	// ============= 统一认证 ================
	UserInfoByPhone string = "/hy/getUserInfoByPhone"
)
