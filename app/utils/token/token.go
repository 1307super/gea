package token

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
	"time"
)

type GeaClaims struct {
	//【JWT ID】     该jwt的唯一ID编号
	//【issuer】     发布者的url地址
	//【issued at】  该jwt的发布时间；unix 时间戳
	//【subject】    该JWT所面向的用户，用于处理特定应用，不是常用的字段
	//【audience】   接受者的url地址
	//【expiration】 该jwt销毁的时间；unix时间戳
	//【not before】 该jwt的使用时间不能早于该时间；unix时间戳
	StandardClaims *jwt.StandardClaims
	RefreshTime    int64 //【The refresh time】 该jwt刷新的时间；unix时间戳
	UserId         string `json:"user_id"`
	LoginName      string `json:"login_name"`
}

type Token struct {
	Claim    *GeaClaims
	Token    string
	NewToken string
}

const CacheKey = "GJWT"

//创建Claims
func New() *GeaClaims {
	timeOut := g.Cfg().GetInt("jwt.timeout")
	if timeOut <= 0 {
		timeOut = 3600
	}
	refresh := g.Cfg().GetInt("jwt.refresh")
	if refresh <= 0 {
		refresh = timeOut / 2
	}
	var claims GeaClaims
	standardClaims := new(jwt.StandardClaims)
	standardClaims.Id = guid.S()
	standardClaims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeOut)).Unix()
	standardClaims.IssuedAt = time.Now().Unix()
	claims.RefreshTime = time.Now().Add(time.Second * time.Duration(refresh)).Unix()
	claims.StandardClaims = standardClaims
	return &claims
}

func (c *GeaClaims) SetIss(issuer string) *GeaClaims {
	c.StandardClaims.Issuer = issuer
	return c
}

func (c *GeaClaims) SetSub(subject string) *GeaClaims {
	c.StandardClaims.Subject = subject
	return c
}

func (c *GeaClaims) SetAud(audience string) *GeaClaims {
	c.StandardClaims.Audience = audience
	return c
}

func (c *GeaClaims) SetNbf(notBefore int64) *GeaClaims {
	c.StandardClaims.NotBefore = notBefore
	return c
}
func (c *GeaClaims) SetUserId(userId string) *GeaClaims {
	c.UserId = userId
	return c
}
func (c *GeaClaims) SetLoginName(loginName string) *GeaClaims {
	c.LoginName = loginName
	return c
}

func (c *GeaClaims) Valid() error {
	//标准验证
	return c.StandardClaims.Valid()
}
//创建token
func (c *GeaClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(g.Cfg().GetString("api.jwt.encryptKey"))
	if err != nil {
		return "", err
	}
	c.SetCache(CacheKey+c.LoginName, c.UserId)
	return token.SignedString(mySignKeyBytes)
}

//验证token
func VerifyAuthToken(token string) (*Token, error) {
	var GeaClaims = new(GeaClaims)
	GeaClaims, err := GeaClaims.DecryptToken(token)
	if err != nil {
		return nil, err
	}
	// 从缓存获取
	userId, err := GeaClaims.GetCache(CacheKey+GeaClaims.LoginName)
	if err != nil {
		return nil,err
	}
	rs := new(Token)
	rs.Claim = GeaClaims
	rs.Token = token
	//判断是否需要刷新
	if GeaClaims.RefreshTime > time.Now().Unix() {
		//生成新token
		newToken, err := New().SetUserId(userId).SetLoginName(rs.Claim.LoginName).CreateToken()
		if err == nil {
			rs.NewToken = newToken
		}
	}
	return rs, nil
}

func (c *GeaClaims) DecryptToken(token string) (*GeaClaims, error) {
	mySignKey := g.Cfg().GetString("api.jwt.encryptKey")
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(mySignKey) //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		return nil, err
	}
	parseAuth, err := jwt.ParseWithClaims(token, c, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})
	if err != nil {
		return nil, err
	}
	//验证claims
	if err := parseAuth.Claims.Valid(); err != nil {
		return nil, err
	}
	return c, nil
}
// 设置缓存
func (c *GeaClaims) SetCache(cacheKey string, userId string) string {
	gcache.Set(cacheKey, userId, gconv.Duration(c.RefreshTime)*time.Millisecond*2)
	return userId
}

// 获取缓存
func (c *GeaClaims) GetCache(cacheKey string) (string, error) {
	userCacheValue := gcache.Get(cacheKey)
	if userCacheValue == nil {
		return "", gerror.New("请登录")
	}
	var userId string
	userId = gconv.String(userCacheValue)
	return userId, nil
}
func RemoveCache(cacheKey string) {
	gcache.Remove(cacheKey)
}
