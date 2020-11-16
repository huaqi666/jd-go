package jd

// 关键词商品查询接口
type GoodsQueryRequest struct {
	// 必填 ...
	// 非必填 ...

	Cid1                 uint64   `json:"cid1,omitempty"`                 // 一级类目id
	Cid2                 uint64   `json:"cid2,omitempty"`                 // 二级类目id
	Cid3                 uint64   `json:"cid3,omitempty"`                 // 三级类目id
	PageIndex            uint64   `json:"pageIndex,omitempty"`            // 页码
	PageSize             uint64   `json:"pageSize,omitempty"`             // 每页数量，单页数最大30，默认20
	SkuIds               []uint64 `json:"skuIds,omitempty"`               // skuid集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	Keyword              string   `json:"keyword,omitempty"`              // 关键词，字数同京东商品名称一致，目前未限制
	Pricefrom            float64  `json:"pricefrom,omitempty"`            // 商品价格下限
	Priceto              float64  `json:"priceto,omitempty"`              // 商品价格上限
	CommissionShareStart uint64   `json:"commissionShareStart,omitempty"` // 佣金比例区间开始
	CommissionShareEnd   uint64   `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	Owner                string   `json:"owner,omitempty"`                // 商品类型：自营[g]，POP[p]
	SortName             string   `json:"sortName,omitempty"`             // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string   `json:"sort,omitempty"`                 // asc,desc升降序,默认降序
	IsCoupon             uint64   `json:"isCoupon,omitempty"`             // 是否是优惠券商品，1：有优惠券，0：无优惠券
	IsPG                 string   `json:"isPG,omitempty"`                 // 是否是拼购商品，1：拼购商品，0：非拼购商品
	PingouPriceStart     float64  `json:"pingouPriceStart,omitempty"`     // 拼购价格区间开始
	PingouPriceEnd       float64  `json:"pingouPriceEnd,omitempty"`       // 拼购价格区间结束
	IsHot                uint64   `json:"-,omitempty"`                    // Deprecated: 已废弃，请勿使用
	BrandCode            string   `json:"brandCode,omitempty"`            // 品牌code
	ShopId               uint64   `json:"shopId,omitempty"`               // 店铺Id
	HasContent           uint64   `json:"hasContent,omitempty"`           // 1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        uint64   `json:"hasBestCoupon,omitempty"`        // 1：查询有最优惠券商品；其他值过滤掉此入参条件。
	Pid                  string   `json:"pid,omitempty"`                  // 联盟id_应用iD_推广位id
	Fields               string   `json:"fields,omitempty"`               // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息),hotWords(热词),similar(相似推荐商品),documentInfo(段子信息),bookInfo(图书信息),specInfo(扩展信息)
	ForbidTypes          string   `json:"forbidTypes,omitempty"`          // 10微信京东购物小程序禁售，11微信京喜小程序禁售
	JxFlags              []uint64 `json:"jxFlags,omitempty"`              // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买），入参多个值表示或条件查询
	ShopLevelFrom        float64  `json:"shopLevelFrom,omitempty"`        // 支持传入0.0、2.5、3.0、3.5、4.0、4.5、4.9，默认为空表示不筛选评分
	DeliveryType         uint64   `json:"deliveryType,omitempty"`         // 京东配送 1：是，0：不是
	Isbn                 string   `json:"isbn,omitempty"`                 // 图书编号
}

type GoodsQueryResult struct {
	Message        string `json:"message"`
	SimilarSkuList string `json:"similarSkuList"`
	TotalCount     string `json:"totalCount"`
	Data           struct {
		GoodsResp struct {
			CategoryInfo struct {
				Cid3     string `json:"cid3"`
				Cid2Name string `json:"cid2Name"`
				Cid2     string `json:"cid2"`
				Cid3Name string `json:"cid3Name"`
				Cid1Name string `json:"cid1Name"`
				Cid1     string `json:"cid1"`
			} `json:"categoryInfo"`
			Spuid       string `json:"spuid"`
			IsJdSale    string `json:"isJdSale"`
			MaterialURL string `json:"materialUrl"`
			BookInfo    struct {
				Isbn string `json:"isbn"`
			} `json:"bookInfo"`
			VideoInfo struct {
				VideoList struct {
					Video struct {
						PlayURL   string `json:"playUrl"`
						Duration  string `json:"duration"`
						ImageURL  string `json:"imageUrl"`
						Width     string `json:"width"`
						PlayType  string `json:"playType"`
						High      string `json:"high"`
						VideoType string `json:"videoType"`
					} `json:"video"`
				} `json:"videoList"`
			} `json:"videoInfo"`
			DocumentInfo struct {
				Document string `json:"document"`
				Discount string `json:"discount"`
			} `json:"documentInfo"`
			BrandName string `json:"brandName"`
			ImageInfo struct {
				ImageList struct {
					URLInfo struct {
						URL string `json:"url"`
					} `json:"urlInfo"`
				} `json:"imageList"`
				WhiteImage string `json:"whiteImage"`
			} `json:"imageInfo"`
			BrandCode      string `json:"brandCode"`
			SkuID          string `json:"skuId"`
			CommissionInfo struct {
				StartTime           string `json:"startTime"`
				CommissionShare     string `json:"commissionShare"`
				CouponCommission    string `json:"couponCommission"`
				PlusCommissionShare string `json:"plusCommissionShare"`
				EndTime             string `json:"endTime"`
				Commission          string `json:"commission"`
				IsLock              string `json:"isLock"`
			} `json:"commissionInfo"`
			JxFlags            string `json:"jxFlags"`
			InOrderCount30Days string `json:"inOrderCount30Days"`
			CouponInfo         struct {
				CouponList struct {
					Coupon struct {
						GetEndTime   string `json:"getEndTime"`
						GetStartTime string `json:"getStartTime"`
						Quota        string `json:"quota"`
						PlatformType string `json:"platformType"`
						UseEndTime   string `json:"useEndTime"`
						UseStartTime string `json:"useStartTime"`
						BindType     string `json:"bindType"`
						IsBest       string `json:"isBest"`
						Link         string `json:"link"`
						HotValue     string `json:"hotValue"`
						Discount     string `json:"discount"`
					} `json:"coupon"`
				} `json:"couponList"`
			} `json:"couponInfo"`
			SpecInfo struct {
				Spec     string `json:"spec"`
				SpecName string `json:"specName"`
				Color    string `json:"color"`
				Size     string `json:"size"`
			} `json:"specInfo"`
			PinGouInfo struct {
				PingouTmCount   string `json:"pingouTmCount"`
				PingouURL       string `json:"pingouUrl"`
				PingouEndTime   string `json:"pingouEndTime"`
				PingouPrice     string `json:"pingouPrice"`
				PingouStartTime string `json:"pingouStartTime"`
			} `json:"pinGouInfo"`
			DeliveryType string `json:"deliveryType"`
			ForbidTypes  string `json:"forbidTypes"`
			SkuName      string `json:"skuName"`
			IsHot        string `json:"isHot"`
			PriceInfo    struct {
				Price             string `json:"price"`
				HistoryPriceDay   string `json:"historyPriceDay"`
				LowestCouponPrice string `json:"lowestCouponPrice"`
				LowestPriceType   string `json:"lowestPriceType"`
				LowestPrice       string `json:"lowestPrice"`
			} `json:"priceInfo"`
			ShopInfo struct {
				ShopID    string `json:"shopId"`
				ShopName  string `json:"shopName"`
				ShopLevel string `json:"shopLevel"`
			} `json:"shopInfo"`
			Owner             string `json:"owner"`
			GoodCommentsShare string `json:"goodCommentsShare"`
			Comments          string `json:"comments"`
			CommentInfo       struct {
				CommentList struct {
					Content   string `json:"content"`
					ImageList struct {
						URLInfo struct {
							URL string `json:"url"`
						} `json:"urlInfo"`
					} `json:"imageList"`
				} `json:"commentList"`
			} `json:"commentInfo"`
		} `json:"goodsResp"`
	} `json:"data"`
	Code     string `json:"code"`
	HotWords string `json:"hotWords"`
}

type AutoGenerated struct {
	JdUnionOpenGoodsQueryResponse struct {
		QueryResult interface{} `json:"queryResult"`
	} `json:"jd_union_open_goods_query_response"`
}
