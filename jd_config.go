package jd

import (
	"encoding/json"
	"fmt"
	"github.com/cliod/jd-go/common"
	"github.com/cliod/jd-go/log"
	"sort"
	"strings"
	"time"
)

// 系统配置/系统参数
type Config struct {
	Method      Method `json:"method"`       // API接口名称
	AppKey      string `json:"app_key"`      // 分配给应用的AppKey
	AccessToken string `json:"access_token"` // Oauth2颁发的动态令牌,根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传
	Timestamp   string `json:"timestamp"`    // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟
	Format      string `json:"format"`       // 响应格式。暂时只支持json
	Version     string `json:"v"`            // API协议版本，可选值：2.0，请根据API具体版本号传入此参数，一般为1.0
	SignMethod  string `json:"sign_method"`  // 签名的摘要算法， md5
	Sign        string `json:"sign"`         // API输入参数签名结果
	SecretKey   string `json:"-"`            // api秘钥
	RouteApi    string `json:"-"`            // 路由API，默认京东联盟
}

func newConfig(appKey, secretKey, routApi string) *Config {
	return &Config{
		AppKey:      appKey,
		Timestamp:   time.Now().Local().Format("2006-01-02 15:04:05"),
		AccessToken: "", // 不需要授权，不用填写
		Format:      "json",
		SignMethod:  "md5",
		Method:      "",
		Version:     "1.0",
		SecretKey:   secretKey,
		RouteApi:    routApi,
	}
}

// 默认京东联盟
func NewConfig(appKey, secretKey string) *Config {
	return newConfig(appKey, secretKey, BaseUrl)
}

// 参数封装，用于参数校验和签名
type parameter struct {
	Config
	ParamJson    string `json:"param_json,omitempty"`        // 业务参数(string)
	BuyParamJson string `json:"360buy_param_json,omitempty"` // 业务参数(string)
}

func newParameter(config *Config, pj map[string]interface{}) (p *parameter) {
	bs, _ := json.Marshal(pj)
	p = &parameter{Config: *config}
	if strings.Trim(config.RouteApi, " ") == JosRootEndpoint {
		p.BuyParamJson = string(bs)
	} else {
		p.ParamJson = string(bs)
	}
	return
}

// 参数封装，用于请求
type Param struct {
	Method       Method `json:"method" url:"method"`                                           // API接口名称
	AppKey       string `json:"app_key" url:"app_key"`                                         // 分配给应用的AppKey
	AccessToken  string `json:"access_token,omitempty" url:"access_token"`                     // Oauth2颁发的动态令牌,根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传
	Timestamp    string `json:"timestamp" url:"timestamp"`                                     // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟
	Format       string `json:"format" url:"format"`                                           // 响应格式。暂时只支持json
	Version      string `json:"v" url:"v"`                                                     // API协议版本，可选值：2.0，请根据API具体版本号传入此参数，一般为1.0
	SignMethod   string `json:"sign_method" url:"sign_method"`                                 // 签名的摘要算法， md5
	Sign         string `json:"sign" url:"sign"`                                               // API输入参数签名结果
	ParamJson    string `json:"param_json,omitempty" url:"param_json,omitempty"`               // 业务参数
	BuyParamJson string `json:"360buy_param_json,omitempty" url:"360buy_param_json,omitempty"` // 业务参数
}

func NewParam(param *parameter) *Param {
	return &Param{
		BuyParamJson: param.BuyParamJson,
		ParamJson:    param.ParamJson,
		AppKey:       param.AppKey,
		Method:       param.Method,
		AccessToken:  param.AccessToken,
		Timestamp:    param.Timestamp,
		Format:       param.Format,
		Version:      param.Version,
		SignMethod:   param.SignMethod,
		Sign:         param.Sign,
	}
}

// 检查必选参数
func (p *parameter) CheckRequiredParams() (err error) {
	if "" == p.AppKey || "" == p.SecretKey {
		err = fmt.Errorf("AppKey and SecretKey must be set")
		log.Error("Parameter:", err)
		return
	}
	if p.RouteApi == UnionRootEndpoint {
		// 宙斯中format和sign_method不是必须
		if "" == p.Format || "" == p.SignMethod {
			err = fmt.Errorf("format, sign_method must be set")
			log.Error("Parameter:", err)
			return
		}
	}
	if "" == p.Timestamp {
		err = fmt.Errorf("timestamp must be set")
		log.Error("Parameter:", err)
		return
	}
	if "" == p.Method {
		err = fmt.Errorf("method is required")
		log.Error("Parameter:", err)
		return
	}
	if "" == p.Version {
		err = fmt.Errorf("version is required")
		log.Error("Parameter:", err)
		return
	}
	return nil
}

//  把所有请求参数按照参数名称的 ASCII 码表顺序进行排序并拼接
func (p *parameter) getConcatParams() string {
	var (
		sorts  []string
		value  string
		params map[string]interface{}
		concat string
	)
	bs, _ := json.Marshal(p)
	_ = json.Unmarshal(bs, &params)
	for key := range params {
		if "sign" != key {
			value = params[key].(string)
			// 签名兼容宙斯
			if p.RouteApi == UnionRootEndpoint && key == "access_token" && value == "" {
				continue
			}
			sorts = append(sorts, key+value)
		}
	}
	sort.Strings(sorts)
	for i := range sorts {
		concat += sorts[i]
	}
	return concat
}

// 把 appSecret 的值拼接在字符串的两端，使用 MD5 进行加密，并转化成大写
func (p *parameter) attachSign() {
	concatParams := p.getConcatParams()
	p.Sign = common.Md5ToUpper(p.SecretKey + concatParams + p.SecretKey)
}
