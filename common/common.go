package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// http请求接口
type Service interface {
	// 执行Get请求
	Get(url string, args ...interface{}) ([]byte, error)
	// 执行Post请求
	Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error)

	// Get 执行Get请求并将结果转成对象
	GetFor(v interface{}, url string, args ...interface{}) error
	// Post 执行Post请求并将结果转成对象
	PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error
}

// http请求默认实现(json传参)
type ServiceImpl struct {
	client *http.Client
}

func (s *ServiceImpl) Get(url string, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	res, err := s.client.Get(uri)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	body, err := json.Marshal(data)
	res, err := s.client.Post(uri, contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) GetFor(v interface{}, url string, args ...interface{}) error {
	res, err := s.Get(url, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, v)
}

func (s *ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	res, err := s.Post(url, contentType, data, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, v)
}

func NewService() Service {
	return NewServiceFor(http.DefaultClient)
}

func NewServiceFor(client *http.Client) Service {
	return &ServiceImpl{
		client: client,
	}
}
