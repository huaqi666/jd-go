package jd

import (
	"github.com/cliod/jd-go/common"
	"strings"
)

type Service interface {
	common.Service

	// 获取商品相关接口
	GetGoodsService() GoodsService
	// 获取配置
	GetConfig() *Config
	// 设置配置
	SetConfig(config *Config)
	// 设置http请求
	SetHttpService(service common.Service)
}

type ServiceImpl struct {
	service common.Service
	config  *Config

	goodsService GoodsService
}

func NewJdService(appKet, secretKey string) Service {
	impl := &ServiceImpl{}
	impl.SetConfig(NewConfig(appKet, secretKey))
	impl.SetHttpService(common.NewService())
	return impl
}

func (s *ServiceImpl) Get(url string, args ...interface{}) ([]byte, error) {
	return s.service.Get(url, args...)
}

func (s *ServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	return s.service.Post(url, contentType, data, args...)
}

func (s *ServiceImpl) GetFor(v interface{}, url string, args ...interface{}) error {
	return s.service.GetFor(v, url, args...)
}

func (s *ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	return s.service.PostFor(v, url, contentType, data, args...)
}

func (s *ServiceImpl) GetGoodsService() GoodsService {
	return s.goodsService
}

func (s *ServiceImpl) GetConfig() *Config {
	return s.config
}

func (s *ServiceImpl) SetConfig(config *Config) {
	s.config = config
}

func (s *ServiceImpl) SetHttpService(service common.Service) {
	s.service = service
}

func (s *ServiceImpl) Sign(method Method, param map[string]string) map[string]interface{} {
	c := s.GetConfig()
	c.Method = method
	parameter := common.ToMap(c)

	for k, v := range param {
		parameter[k] = v
	}

	parameter["sign"] = common.Sign(parameter, c.SecretKey, "sign")
	return parameter
}

func (s *ServiceImpl) SignGet(method Method, param map[string]string) string {
	parameter := s.Sign(method, param)
	suffix := ""
	for k, v := range parameter {
		suffix += k + "=" + v.(string) + "&"
	}
	return strings.TrimRight(suffix, "&")
}

func (s *ServiceImpl) DoGet(v interface{}, url string) {

}

func (s *ServiceImpl) DoPost(v interface{}, url string) {

}
