package jd

import (
	"encoding/json"
	"fmt"
	"github.com/cliod/jd-go/common"
	"reflect"
	"sort"
	"time"
)

// 系统配置
type Config struct {
	Method      Method `json:"method" url:"method"`                       // API接口名称
	AppKey      string `json:"app_key" url:"app_key"`                     // 分配给应用的AppKey
	AccessToken string `json:"access_token,omitempty" url:"access_token"` // Oauth2颁发的动态令牌,根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传
	Timestamp   string `json:"timestamp" url:"timestamp"`                 // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟
	Format      string `json:"format" url:"format"`                       // 响应格式。暂时只支持json
	Version     string `json:"v" url:"v"`                                 // API协议版本，可选值：2.0，请根据API具体版本号传入此参数，一般为1.0
	SignMethod  string `json:"sign_method" url:"sign_method"`             // 签名的摘要算法， md5
	Sign        string `json:"sign" url:"sign"`                           // API输入参数签名结果
	SecretKey   string `json:"-" url:"-"`                                 // api秘钥
}

func NewConfig(appKey, secretKey string) *Config {
	return &Config{
		AppKey:    appKey,
		Timestamp: time.Now().Local().Format("2006-01-02 15:04:05"),
		//Timestamp:   "2020-11-17 13:20:47",
		AccessToken: "", // 不需要授权，不用填写
		Format:      "json",
		SignMethod:  "md5",
		Method:      "",
		Version:     "1.0",
		SecretKey:   secretKey,
	}
}

// 参数封装
type parameter struct {
	Config
	ParamJson map[string]interface{} `json:"param_json,omitempty" url:"param_json"`
}

// 参数封装
type Param struct {
	Config
	ParamJson string `json:"param_json,omitempty" url:"param_json"`
}

func NewParameter(config *Config, pj map[string]interface{}) *parameter {
	return &parameter{Config: *config, ParamJson: pj}
}

// 检查必选参数
func (p *parameter) CheckRequiredParams() error {
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

//  把所有请求参数按照参数名称的 ASCII 码表顺序进行排序并拼接
func (p *parameter) getConcatParams() string {
	var params map[string]interface{}
	bs, _ := json.Marshal(p)
	_ = json.Unmarshal(bs, &params)
	s := make([]string, len(params))
	for k := range params {
		if "sign" != k {
			v := params[k]
			if "string" != reflect.TypeOf(params[k]).String() {
				valueBs, _ := json.Marshal(params[k])
				v = string(valueBs)
			}
			s = append(s, fmt.Sprintf("%s%s", k, v))
		}
	}
	sort.Strings(s)
	var concat string
	for i := range s {
		concat += s[i]
	}
	return concat
}

// 组装 HTTP 请求，将所有参数名和参数值采用 utf-8 进行 URL 编码（
func (p *parameter) getParamString() string {
	bs, _ := json.Marshal(p.ParamJson)
	return string(bs)
}

// 把 appSecret 的值拼接在字符串的两端，使用 MD5 进行加密，并转化成大写
func (p *parameter) attachSign() {
	concatParams := p.getConcatParams()
	p.Sign = common.Md5ToUpper(p.SecretKey + concatParams + p.SecretKey)
}
