package jd

import (
	"fmt"
	"time"
)

// 系统配置
type Config struct {
	Method      Method `json:"method"`                 // API接口名称
	AppKey      string `json:"app_key"`                // 分配给应用的AppKey
	AccessToken string `json:"access_token,omitempty"` // Oauth2颁发的动态令牌,根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传
	Timestamp   string `json:"timestamp"`              // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟
	Format      string `json:"format"`                 // 响应格式。暂时只支持json
	Version     string `json:"v"`                      // API协议版本，可选值：2.0，请根据API具体版本号传入此参数，一般为1.0
	SignMethod  string `json:"sign_method"`            // 签名的摘要算法， md5
	Sign        string `json:"sign"`                   // API输入参数签名结果
	SecretKey   string `json:"-"`                      // api秘钥
}

func NewConfig(appKey, secretKey string) *Config {
	return &Config{
		AppKey:      appKey,
		Timestamp:   time.Now().Local().Format("2006-01-02 15:04:05"),
		AccessToken: "", // 不需要授权，不用填写
		Format:      "json",
		SignMethod:  "md5",
		Method:      "",
		Version:     "1.0",
		SecretKey:   secretKey,
	}
}

// 检查必选参数
func (p *Config) CheckRequiredParams() error {
	if "" == p.AppKey || "" == p.SecretKey {
		return fmt.Errorf("AppKey and SecretKey must be set")
	}
	if "" == p.Format || "" == p.SignMethod || "" == p.Timestamp {
		return fmt.Errorf("format, sign_method and timestamp must be set")
	}
	if "" == p.Method {
		return fmt.Errorf("method is required")
	}
	if "" == p.Version {
		return fmt.Errorf("version is required")
	}
	return nil
}
