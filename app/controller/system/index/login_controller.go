package index

import (
	"gea/app/model"
	logininforModel "gea/app/model/monitor/logininfor"
	"gea/app/model/system/user_online"
	logininforService "gea/app/service/monitor/logininfor"
	userService "gea/app/service/system/user"
	"gea/app/utils/ip"
	"gea/app/utils/response"
	"gea/app/utils/token"
	"gea/library/casbin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
)

type RegisterReq struct {
	UserName     string `p:"username"  v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password     string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	ValidateCode string `p:"validateCode" v:"required|length:4,30#请输入验证码|验证码长度不够"`
	IdKey        string `p:"idkey" v:"required|length:4,30#请输入验证码id|验证码id长度不够"`
}
type RegisterReq2 struct {
	UserName     string `p:"username"  v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password     string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	//ValidateCode string `p:"validateCode" v:"required|length:4,30#请输入验证码|验证码长度不够"`
	//IdKey        string `p:"idkey" v:"required|length:4,30#请输入验证码id|验证码id长度不够"`
}

// 图形验证码
func CaptchaImage(r *ghttp.Request) {
	//config struct for digits
	//数字验证码配置
	//var configD = base64Captcha.ConfigDigit{
	//	Height:     80,
	//	Width:      240,
	//	MaxSkew:    0.7,
	//	DotCount:   80,
	//	CaptchaLen: 5,
	//}
	//config struct for audio
	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 36,
		Width:  111,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	r.Response.WriteJsonExit(model.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

//验证登陆
func CheckLogin(r *ghttp.Request) {
	var req *RegisterReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}
	//比对验证码
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)

	if !verifyResult {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  "验证码不正确",
		})
	}
	isLock := logininforService.CheckLock(req.UserName)
	if isLock {
		r.Response.WriteJsonExit(model.CommonRes{
			Code: 500,
			Msg:  "账号已锁定，请30分钟后再试",
		})
	}

	//记录日志
	var logininfor logininforModel.Entity
	logininfor.LoginName = req.UserName
	logininfor.Ipaddr = r.GetClientIp()

	userAgent := r.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	loginIp := r.GetClientIp()
	loginLocation := ip.GetCityByIp(loginIp)
	logininfor.Os = os
	logininfor.Browser = browser
	logininfor.LoginTime = gtime.Now()
	logininfor.LoginLocation = loginLocation

	user, err := userService.SignIn(req.UserName, req.Password)
	if err != nil || user == nil {
		// 登录失败
		logininfor.Msg = "账号或密码不正确"
		logininfor.Status = "1"
		logininfor.Insert()

		errTimes := logininforService.SetPasswordCounts(req.UserName)
		having := 5 - errTimes
		response.ErrorResp(r).SetMsg("账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定").WriteJsonExit()
	}

	//移除登陆次数记录
	logininforService.RemovePasswordCounts(req.UserName)
	// 校验通过
	//获取用户token
	jwtToken, err := token.New().SetUserId(gconv.String(user.UserId)).SetLoginName(user.LoginName).CreateToken()
	if err != nil {
		response.ErrorResp(r).SetData("Error while signing the token").WriteJsonExit()
		return
	}
	//保存在线状态
	var userOnline user_online.Entity
	userOnline.Token = token.CacheKey + user.LoginName
	userOnline.LoginName = req.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.Ipaddr = loginIp
	userOnline.ExpireTime = g.Cfg().GetInt("jwt.timeout")
	userOnline.StartTimestamp = gtime.Now()
	userOnline.LastAccessTime = gtime.Now()
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	userOnline.Replace()

	// 记录日志
	logininfor.Msg = "登陆成功"
	logininfor.Status = "0"
	logininfor.Insert()

	response.SucessResp(r).SetMsg("登录成功").SetData(g.Map{
		"token": jwtToken,
	}).WriteJsonExit()
}

// 添加用户与角色关系
func AddP(r *ghttp.Request) {
	roleName := r.GetString("roleName")
	userName := r.GetString("userName")
	isOk, err := casbin.AddRoleForUser(userName,roleName)
	if err != nil {
		r.Response.WriteExit("失败")
	}
	if !isOk {
		r.Response.WriteExit("添加失败")
	}
	r.Response.WriteExit("ok")
}
// 添加角色与资源
func AddRP(r *ghttp.Request) {
	roleName := r.GetString("roleName")
	path := r.GetString("path")
	method := r.GetString("method")
	isOk, err := casbin.AddPermissionForUser(roleName,path,method)
	if err != nil {
		r.Response.WriteExit("失败")
	}
	if !isOk {
		r.Response.WriteExit("添加失败")
	}
	r.Response.WriteExit("ok")
}
// 删除用户与角色
func DeleteR (r *ghttp.Request)  {
	userName := r.GetString("userName")
	roleName := r.GetString("roleName")
	isOk, err := casbin.DeleteRoleForUser(userName,roleName)
	if err != nil {
		r.Response.WriteExit("失败")
	}
	if !isOk {
		r.Response.WriteExit("删除失败")
	}
	r.Response.WriteExit("ok")
}
// 删除角色与资源
func DeleteP(r *ghttp.Request)  {
	roleName := r.GetString("roleName")
	path := r.GetString("path")
	method := r.GetString("method")
	isOk, err := casbin.DeletePermissionForUser(roleName,path,method)
	if err != nil {
		r.Response.WriteExit("失败")
	}
	if !isOk {
		r.Response.WriteExit("删除失败")
	}
	r.Response.WriteExit("ok")
}

// 校验
func CheckP(r *ghttp.Request)  {
	userName := r.GetString("userName")
	path := r.GetString("path")
	method := r.GetString("method")
	isOk, err := casbin.Enforce(userName,path,method)
	if err != nil {
		r.Response.WriteExit("失败")
	}
	if !isOk {
		r.Response.WriteExit("没有权限")
	}
	r.Response.WriteExit("ok")
}

