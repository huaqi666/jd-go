package common

import (
	"fmt"
	"github.com/cliod/jd-go/log"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
)

// http请求接口
type Service interface {
	// 执行Get请求
	Get(url string, args interface{}) ([]byte, error)
}

// http请求默认实现(json传参)
type ServiceImpl struct {
	client *http.Client
}

func (s *ServiceImpl) Get(address string, v interface{}) ([]byte, error) {
	urlVal, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	values, _ := query.Values(v)
	urlVal.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlVal.String(), nil)
	if err != nil {
		return nil, err
	}

	log.Debug("Request URL", req.URL.String())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return nil, fmt.Errorf("statusCode: %d", statusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewService() Service {
	return &ServiceImpl{
		client: http.DefaultClient,
	}
}
