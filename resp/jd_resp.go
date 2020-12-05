package resp

import "github.com/cliod/jd-go/resp/entity"

type GoodsJingfenQueryResponse struct {
	Code       int                      `json:"code"`
	Message    string                   `json:"message"`
	RequestId  string                   `json:"requestId"`
	TotalCount int64                    `json:"totalCount"`
	Data       []*JingfenGoodsQueryResp `json:"data"`
}

type JingfenGoodsQueryResp struct {
	entity.CategoryInfo   `json:"categoryInfo"`   //类目信息
	Comments              int                     `json:"comments"` //评论数
	entity.CommissionInfo `json:"commissionInfo"` //佣金信息
	entity.CouponInfo     `json:"couponInfo"`     //优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64                 `json:"goodCommentsShare"` //商品好评率
	entity.ImageInfo      `json:"imageinfo"`      //图片信息
	InOrderCount30Days    int64                   `json:"inOrderCount30Days"` //30天引单数量
	MaterialUrl           string                  `json:"materialUrl"`        //商品落地页
	entity.PriceInfo      `json:"priceInfo"`      //价格信息
	entity.ShopInfo       `json:"shopinfo"`       //店铺信息
	SkuId                 int64                   `json:"skuId"`     //商品ID
	SkuName               string                  `json:"skuName"`   //商品名称
	IsHot                 int                     `json:"isHot"`     //已废弃，请勿使用
	Spuid                 int64                   `json:"spuid"`     //spuid，其值为同款商品的主skuid
	BrandCode             string                  `json:"brandCode"` //品牌code
	BrandName             string                  `json:"brandName"` //品牌名
	Owner                 string                  `json:"owner"`     //g=自营，p=pop
	entity.PinGouInfo     `json:"pinGouInfo"`     //拼购信息
	entity.ResourceInfo   `json:"resourceInfo"`   //资源信息
	InOrderCount30DaysSku int64                   `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	entity.SeckillInfo    `json:"seckillInfo"`    //秒杀信息
	JxFlgs                []uint                  `json:jxFlags` //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	entity.VideoInfo      `json:videoInfo`        //视频信息
	entity.DocumentInfo   `json:documentInfo`     //段子信息
}

type GoodsQueryResponse struct {
	Code           int               `json:"code"`
	Message        string            `json:"message"`
	SimilarSkuList string            `json:"similarSkuList"`
	TotalCount     int64             `json:"totalCount"`
	Data           []*GoodsQueryResp `json:"data"`
	HotWords       string            `json:"hotWords"`
}

type GoodsQueryResp struct {
	entity.CategoryInfo   `json:"categoryInfo"`   //类目信息
	Comments              int                     `json:"comments"` //评论数
	entity.CommissionInfo `json:"commissionInfo"` //佣金信息
	entity.CouponInfo     `json:"couponInfo"`     //优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64                 `json:"goodCommentsShare"` //商品好评率
	entity.ImageInfo      `json:"imageinfo"`      //图片信息
	InOrderCount30DaysSku int64                   `json:"inOrderCount30DaysSku"` //30天引单数量(sku维度)
	MaterialUrl           string                  `json:"materialUrl"`           //商品落地页
	entity.PriceInfo      `json:"priceInfo"`      //价格信息
	entity.ShopInfo       `json:"shopinfo"`       //店铺信息
	SkuId                 int64                   `json:"skuId"`     //商品ID
	SkuName               string                  `json:"skuName"`   //商品名称
	IsHot                 int                     `json:"isHot"`     //已废弃，请勿使用
	Spuid                 int64                   `json:"spuid"`     //spuid，其值为同款商品的主skuid
	BrandCode             string                  `json:"brandCode"` //品牌code
	BrandName             string                  `json:"brandName"` //品牌名
	Owner                 string                  `json:"owner"`     //g=自营，p=pop
	entity.PinGouInfo     `json:"pinGouInfo"`     //拼购信息
	entity.VideoInfo      `json:videoInfo`        //视频信息
	entity.CommentInfo    `json:commentInfo`      //评价信息
	JxFlgs                []uint                  `json:jxFlags` //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	entity.DocumentInfo   `json:documentInfo`     //段子信息
}

type CouponQueryResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []*CouponResp `json:"data"`
}

// 优惠券
type CouponResp struct {
	TakeEndTime   int64   `json:"takeEndTime"`   // 券领取结束时间(时间戳，毫秒)
	TakeBeginTime int64   `json:"takeBeginTime"` // 券领取开始时间(时间戳，毫秒)
	RemainNum     int64   `json:"remainNum"`     //券剩余张数
	Yn            string  `json:"yn"`            // 券有效状态
	Num           int64   `json:"num"`           // 券总张数
	Quota         float32 `json:"quota"`         // 券消费限额
	Link          string  `json:"link"`          // 券链接
	Discount      float32 `json:"discount"`      //券面额
	BeginTime     int64   `json:"beginTime"`     // 券有效使用开始时间(时间戳，毫秒)
	EndTime       int64   `json:"endTime"`       //券有效使用结束时间(时间戳，毫秒)
	Platform      string  `json:"platform"`      // 券使用平台
}
