## jd-go - 京东联盟开发 Golang SDK（开发工具包） 

[![travis-image]][travis-url]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

### 使用方式
注意：最新版本（包括测试版）为 [![GitHub release](https://img.shields.io/badge/github-releases-blue)](https://github.com/cliod/jd-go/releases)

```shell script
go get -u github.com/cliod/jd-go
```

### 使用Demo
公众号使用例子
```go
j := NewJdService("<app_key>", "<app_secret>")
// goods api
g := j.GetGoodsService()
res, err := g.GoodsJingfenQuery(&GoodsJingfenQueryRequest{
    EliteId: 33,
})
if err != nil {
    fmt.Println(res)
}
```

[travis-image]: https://api.travis-ci.com/cliod/jd-go.svg?branch=main
[travis-url]: https://travis-ci.com/cliod/jd-go
