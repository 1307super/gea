package api

import (

	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/ip"
	"gea/app/utils/response"
	"gea/app/utils/token"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
)

var Login = new(loginApi)

type loginApi struct {
	BaseController
}

// 图形验证码
func (a *loginApi)CaptchaImage(r *ghttp.Request) {
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

	r.Response.WriteJsonExit(response.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

//验证登陆
func (a *loginApi) CheckLogin(r *ghttp.Request) {
	var (
		req *define.UserApiLoginReq
		serverLoginReq *define.UserServiceLoginReq
	)
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	//比对验证码
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)
	if !verifyResult {
		a.Err(r,"验证码不正确")
	}
	if err := gconv.Struct(req, &serverLoginReq); err != nil {
		a.Err(r,err.Error())
	}

	//记录日志
	var logininfor model.SysLogininfor
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
	if err := service.User.Login(r.Context(),serverLoginReq); err != nil{
		// 登录失败
		logininfor.Msg = "账号或密码不正确"
		logininfor.Status = "1"
		dao.SysLogininfor.Insert(logininfor)

		errTimes := service.User.SetPasswordCounts(req.UserName)
		having := 5 - errTimes
		a.Err(r,"账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定")
	}

	//保存在线状态
	ctx := shared.Context.Get(r.Context())
	var userOnline model.SysUserOnline
	userOnline.Token = token.CacheKey + ctx.User.LoginName
	userOnline.LoginName = req.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.Ipaddr = loginIp
	userOnline.ExpireTime = g.Cfg().GetInt("jwt.timeout")
	userOnline.StartTimestamp = gtime.Now()
	userOnline.LastAccessTime = gtime.Now()
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	dao.SysUserOnline.Save(userOnline)

	// 记录日志
	logininfor.Msg = "登陆成功"
	logininfor.Status = "0"
	dao.SysLogininfor.Insert(logininfor)

	a.Succ(r,g.Map{
		"token": ctx.Token,
	})
}

//注销
func (a *loginApi)Logout(r *ghttp.Request) {
	user := shared.Context.Get(r.Context()).User
	if user == nil{
		a.Succ(r)
	}
	token.RemoveCache(token.CacheKey + user.LoginName)
	a.Succ(r)
}