package jd

type GoodsJingfenQueryRequest struct {
	// 必填 ...

	EliteId uint64 `json:"eliteId"` // 频道ID:1-好券商品,2-精选卖场,10-9.9包邮,15-京东配送,22-实时热销榜,23-为你推荐,24-数码家电,25-超市,26-母婴玩具,27-家具日用,28-美妆穿搭,30-图书文具,
	// 31-今日必推,32-京东好物,33-京东秒杀,34-拼购商品,40-高收益榜,41-自营热卖榜,108-秒杀进行中,109-新品首发,110-自营,112-京东爆品,125-首购商品,129-高佣榜单,130-视频商品,153-历史最低价商品榜
	// 非必填 ...

	PageIndex   uint64 `json:"pageIndex,omitempty"`   // 页码，默认1
	PageSize    uint64 `json:"pageSize,omitempty"`    // 每页数量，默认20，上限50，建议20
	SortName    string `json:"sortName,omitempty"`    // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort        string `json:"sort,omitempty"`        // asc,desc升降序,默认降序
	Pid         string `json:"pid,omitempty"`         // 联盟id_应用id_推广位id，三段式
	Fields      string `json:"fields,omitempty"`      // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
	ForbidTypes string `json:"forbidTypes,omitempty"` // 10微信京东购物小程序禁售，11微信京喜小程序禁售
}

type GoodsJingfenQueryResult struct {
	JdUnionOpenGoodsJingfenQueryResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_jingfen_query_response"`
}

func (r *GoodsJingfenQueryResult) GetResult() []byte {
	str := r.JdUnionOpenGoodsJingfenQueryResponse.Result
	return []byte(str)
}

// Deprecated: 弃用
type GoodsJingfenQueryResultInner struct {
	Code int `json:"code"`
	Data []struct {
		BrandCode    string `json:"brandCode"`
		BrandName    string `json:"brandName"`
		CategoryInfo struct {
			Cid1     int    `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     int    `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     int    `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission          float64 `json:"commission"`
			CommissionShare     int     `json:"commissionShare"`
			EndTime             int64   `json:"endTime"`
			IsLock              int     `json:"isLock"`
			PlusCommissionShare int     `json:"plusCommissionShare"`
			StartTime           int64   `json:"startTime"`
		} `json:"commissionInfo"`
		CouponInfo struct {
			CouponList []interface{} `json:"couponList"`
		} `json:"couponInfo"`
		DeliveryType      int   `json:"deliveryType"`
		ForbidTypes       []int `json:"forbidTypes"`
		GoodCommentsShare int   `json:"goodCommentsShare"`
		ImageInfo         struct {
			ImageList []struct {
				URL string `json:"url"`
			} `json:"imageList"`
			WhiteImage string `json:"whiteImage"`
		} `json:"imageInfo"`
		InOrderCount30Days    int    `json:"inOrderCount30Days"`
		InOrderCount30DaysSku int    `json:"inOrderCount30DaysSku"`
		IsHot                 int    `json:"isHot"`
		MaterialURL           string `json:"materialUrl"`
		Owner                 string `json:"owner"`
		PinGouInfo            struct {
		} `json:"pinGouInfo"`
		PriceInfo struct {
			LowestPrice     int `json:"lowestPrice"`
			LowestPriceType int `json:"lowestPriceType"`
			Price           int `json:"price"`
		} `json:"priceInfo"`
		ResourceInfo struct {
			EliteID   int    `json:"eliteId"`
			EliteName string `json:"eliteName"`
		} `json:"resourceInfo"`
		SeckillInfo struct {
			SeckillEndTime   int64 `json:"seckillEndTime"`
			SeckillOriPrice  int   `json:"seckillOriPrice"`
			SeckillPrice     int   `json:"seckillPrice"`
			SeckillStartTime int64 `json:"seckillStartTime"`
		} `json:"seckillInfo"`
		ShopInfo struct {
			ShopID    int     `json:"shopId"`
			ShopLevel float64 `json:"shopLevel"`
			ShopName  string  `json:"shopName"`
		} `json:"shopInfo"`
		SkuID    int    `json:"skuId"`
		SkuName  string `json:"skuName"`
		Spuid    int    `json:"spuid"`
		BookInfo struct {
			Isbn string `json:"isbn"`
		} `json:"bookInfo,omitempty"`
	} `json:"data"`
	Message    string `json:"message"`
	RequestID  string `json:"requestId"`
	TotalCount int    `json:"totalCount"`
}

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
	IsHot                uint64   `json:"-"`                              // Deprecated: 已废弃，请勿使用
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
	JdUnionOpenGoodsQueryResponse struct {
		QueryResult struct {
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
		} `json:"queryResult"`
	} `json:"jd_union_open_goods_query_response"`
}

type GoodsPromotiongoodsinfoResult struct {
}

type CategoryGoodsRequest struct {
	// 必填 ...

	ParentId uint64 `json:"parentId"` // 父类目id(一级父类目为0)
	Grade    uint64 `json:"grade"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
}

type CategoryGoodsResult struct {
}

type GoodsGigfieldRequest struct {
	// 必填 ...

	SkuIds []uint64 `json:"skuIds"` // skuId集合，最多支持批量入参10个sku
	// 非必填 ...

	Fields []string `json:"fields"` // 查询域集合，不填写则查询全部，目目前支持：categoryInfo（类目信息）,imageInfo（图片信息）,baseBigFieldInfo（基础大字段信息）,bookBigFieldInfo（图书大字段信息）,videoBigFieldInfo（影音大字段信息）,detailImages（商详图）
}

type GoodsGigfieldResult struct {
}

type PromoteCommonGetRequest struct {
	// 必填 ...

	MaterialId string `json:"materialId"` //	推广物料url，例如活动链接、商品链接等；不支持仅传入skuid
	SiteId     string `json:"siteId"`     // 网站ID/APP ID，入口：京东联盟-推广管理-网站管理/APP管理-查看网站ID/APP ID（1、接口禁止使用导购媒体id入参；2、投放链接的网址或应用必须与传入的网站ID/AppID备案一致，否则订单会判“无效-来源与备案网址不符”）
	// 非必填 ...

	PositionId    uint64 `json:"positionId,omitempty"`    // 推广位id
	SubUnionId    string `json:"subUnionId,omitempty"`    // 子渠道标识，您可自定义传入字母、数字或下划线，最多支持80个字符，该参数会在订单行查询接口中展示（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301）
	Ext1          string `json:"ext1,omitempty"`          // 系统扩展参数（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），最多支持40字符，参数会在订单行查询接口中展示
	Protocol      uint64 `json:"-"`                       // Deprecated:【已废弃】请勿再使用
	Pid           string `json:"pid,omitempty"`           // 联盟子推客身份标识（不能传入接口调用者自己的pid）
	CouponUrl     string `json:"couponUrl,omitempty"`     // 优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	GiftCouponKey string `json:"giftCouponKey,omitempty"` // 礼金批次号
}

type PromoteCommonGetResult struct {
	JdUnionOpenPromotionCommonGetResponse struct {
		GetResult string `json:"getResult"`
	} `json:"jd_union_open_promotion_common_get_response"`
}

type PromotionBysubunionidQueryRequest struct {
	// 必填 ...

	MaterialId string `json:"materialId"` // 推广物料url，例如活动链接、商品链接等；不支持仅传入skuid
	// 非必填 ...

	SubUnionId    string `json:"subUnionId,omitempty"`    // 子渠道标识，您可自定义传入字母、数字或下划线，最多支持80个字符，该参数会在订单行查询接口中展示（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301）
	PositionId    uint64 `json:"positionId,omitempty"`    // 推广位ID
	Pid           string `json:"pid,omitempty"`           // 联盟子推客身份标识（不能传入接口调用者自己的pid）
	CouponUrl     string `json:"couponUrl,omitempty"`     // 优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	ChainType     uint64 `json:"chainType,omitempty"`     // 转链类型，1：长链， 2 ：短链 ，3： 长链+短链，默认短链，短链有效期60天
	GiftCouponKey string `json:"giftCouponKey,omitempty"` // 礼金批次号
}

type PromotionBysubunionidQueryResult struct {
}

type PromotionByunionidQueryRequest struct {
	// 必填 ...

	MaterialId string `json:"materialId"` // 推广物料url，例如活动链接、商品链接等；不支持仅传入skuid
	UnionId    uint64 `json:"unionId"`    // 目标推客的联盟ID
	// 非必填 ...

	PositionId uint64 `json:"positionId,omitempty"` // 新增推广位id （不填的话，为其默认生成一个唯一此接口推广位-名称：微信手Q短链接）
	Pid        string `json:"pid,omitempty"`        // 联盟子推客身份标识（不能传入接口调用者自己的pid）
	CouponUrl  string `json:"couponUrl,omitempty"`  // 优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
	SubUnionId string `json:"subUnionId,omitempty"` // 子渠道标识，您可自定义传入字母、数字或下划线，最多支持80个字符，该参数会在订单行查询接口中展示（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301）
	ChainType  uint64 `json:"chainType,omitempty"`  // 转链类型，1：长链， 2 ：短链 ，3： 长链+短链，默认短链，短链有效期60天
}

type PromotionByunionidQueryResult struct {
}

type CouponQueryRequest struct {
}

type CouponQueryResult struct {
	JdUnionOpenCouponQueryResponse struct {
		QueryResult struct {
			Message string `json:"message"`
			Data    struct {
				CouponResp struct {
					RemainNum     string `json:"remainNum"`
					Platform      string `json:"platform"`
					Quota         string `json:"quota"`
					Num           string `json:"num"`
					TakeBeginTime string `json:"takeBeginTime"`
					TakeEndTime   string `json:"takeEndTime"`
					Link          string `json:"link"`
					BeginTime     string `json:"beginTime"`
					Yn            string `json:"yn"`
					EndTime       string `json:"endTime"`
					Discount      string `json:"discount"`
				} `json:"couponResp"`
			} `json:"data"`
			Code string `json:"code"`
		} `json:"queryResult"`
	} `json:"jd_union_open_coupon_query_response"`
}

type PositionQueryRequest struct {
	// 必填 ...

	UnionId   uint64 `json:"unionId"`   // 需要查询的目标联盟id
	Key       string `json:"key"`       // 推客unionid对应的“授权Key”，在联盟官网-我的工具-我的API中查询，授权Key具有唯一性，有效期365天，删除或创建新的授权Key后原有的授权Key自动失效
	UnionType uint64 `json:"unionType"` // 3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，会在推客联盟后台展示，可用于内容平台推广
	PageIndex uint64 `json:"pageIndex"` // 页码，上限100
	PageSize  uint64 `json:"pageSize"`  // 每页条数，上限100
}

type PositionQueryResult struct {
}

type PositionCreateRequest struct {
	// 必填 ...

	UnionId       uint64   `json:"unionId"`       // 需要查询的目标联盟id
	Key           string   `json:"key"`           // 推客unionid对应的“授权Key”，在联盟官网-我的工具-我的API中查询，授权Key具有唯一性，有效期365天，删除或创建新的授权Key后原有的授权Key自动失效
	UnionType     uint64   `json:"unionType"`     // 3：私域推广位，上限5000个，不在联盟后台展示，无对应 PID；4：联盟后台推广位，上限500个，会在推客联盟后台展示，可用于内容平台推广
	TypeNum       uint64   `json:"type"`          // 站点类型 1.网站推广位2.APP推广位3.导购媒体推广位4.聊天工具推广位
	SpaceNameList []string `json:"spaceNameList"` // 推广位名称集合，英文,分割；上限50
	// 非必填 ...

	SiteId uint64 `json:"siteId"` // 站点ID：网站的ID/app ID/snsID 。当type非4(聊天工具)时，siteId必填
}

type PositionCreateResult struct {
}

type UserPidGetRequest struct {
	// 必填 ...

	UnionId       uint64 `json:"unionId"`       // 联盟ID
	ChildUnionId  uint64 `json:"childUnionId"`  // 子站长ID
	PromotionType uint64 `json:"promotionType"` // 推广类型,1APP推广 2聊天工具推广
	MediaName     string `json:"mediaName"`     // 媒体名称，即子站长的app应用名称，推广方式为app推广时必填，且app名称必须为已存在的app名称
	// 非必填 ...

	PositionName string `json:"positionName"` // 子站长的推广位名称，如不存在则创建，不填则由联盟根据母账号信息创建
}

type UserPidGetResult struct {
}

type ActivityQueryRequest struct {
	// 非必填 ...

	PageIndex  uint64 `json:"pageIndex"`  // 页码，默认1
	PageSize   uint64 `json:"pageSize"`   // 每页数量，默认20，上限50
	PoolId     uint64 `json:"poolId"`     // 活动物料ID，1：营销日历热门会场；2：营销日历热门榜单；6：PC站长端官方活动
	ActiveDate uint64 `json:"activeDate"` // 按单个日期查询活动，查询日期范围为过去或未来15天。建议按日期依次查询当天及未来的活动
}

type ActivityQueryResult struct {
}

type StatisticsRedpacketRequest struct {
	// 必填 ...

	StartDate string `json:"startDate"` // 开始日期yyyy-MM-dd，不可超出最近90天开始日期yyyy-MM-dd，不可超出最近90天
	EndDate   string `json:"endDate"`   // 结束时间yyyy-MM-dd，不可超出最近90天
	PageIndex uint64 `json:"pageIndex"` // 页码，起始1
	PageSize  uint64 `json:"pageSize"`  // 每页数，大于等于10且小于等于100的正整数
	// 非必填 ...

	ActId      uint64 `json:"actId"`      // 京享红包活动Id，如不传则查询全部京享红包活动
	PositionId uint64 `json:"positionId"` // 推广位ID，支持联盟后台推广位和API创建的私域推广位
}

type StatisticsRedpacketResult struct {
}
type GiftGetRequest struct {
	// 必填 ...

	SkuMaterialId    string  `json:"skuMaterialId"`    // 商品skuId或落地页地址
	Discount         float64 `json:"discount"`         // 优惠券面额，最小不可低于1元，最大不可超过pop商品价格的80%，自营商品价格的50%
	Amount           uint64  `json:"amount"`           // 总数量
	ReceiveStartTime string  `json:"receiveStartTime"` // 领取开始时间(yyyy-MM-dd HH)，区间为(创建当天0点直至未来6天内)，系统补充为yyyy-MM-dd HH:00:00
	ReceiveEndTime   string  `json:"receiveEndTime"`   // 领取结束时间(yyyy-MM-dd HH)，区间为(创建当前时间点直至未来6天内)，系统补充为yyyy-MM-dd HH:59:59
	IsSpu            uint64  `json:"isSpu"`            // 是否绑定同spu商品(1:是;0:否)，例如skuMaterialId输入一款37码的鞋，当isSpu选择1时，此款鞋的全部尺码均可推广这张礼金；当isSpu选择0时，此款鞋仅37码可推广这张礼金，其他鞋码不支持
	ExpireType       uint64  `json:"expireType"`       // 使用时间类型：1.相对时间，需配合effectiveDays一同传入；2.绝对时间，需配合useStartTime和useEndTime一同传入
	Share            uint64  `json:"share"`            // 每个礼金推广链接是否限制仅可领取1张礼金：-1不限，1限制
	// 非必填 ...

	EffectiveDays uint64 `json:"effectiveDays,omitempty"` // 消费者领取后n天内可用，时间天数1至7，当expireType=1时，必须设置该字段
	UseStartTime  string `json:"useStartTime,omitempty"`  // 消费者领取后的使用开始时间，格式：yyyy-MM-dd，系统补充为yyyy-MM-dd HH:00:00，当expireType=2时，必须设置该字段
	UseEndTime    string `json:"useEndTime,omitempty"`    // 消费者领取后的使用结束时间，格式：yyyy-MM-dd，系统补充为yyyy-MM-dd HH:59:59，当expireType=2时，必须设置该字段
	ContentMatch  uint64 `json:"contentMatch,omitempty"`  // 是否允许通过内容平台推广，0：不允许，1：允许；默认为0
}

type GiftGetResult struct {
}

type GiftStopRequest struct {
	// 必填 ...

	GiftCouponKey string `json:"giftCouponKey"` // 礼金批次ID
}

type GiftStopResult struct {
}

type GiftStatisticRequest struct {
	// 必填 ...

	SkuId         uint64 `json:"skuId"`         // 查询该SKU您创建的推客礼金，以及该SKU您可推广的联盟礼金。 skuId和giftCouponKey二选一，不可同时入参。
	GiftCouponKey string `json:"giftCouponKey"` // 根据礼金批次ID精确查询礼金信息，请勿和createTime同时传入。 skuId和giftCouponKey二选一，不可同时入参。
	CreateTime    string `json:"createTime"`    // 可查询此日期及以后创建的礼金，如不传则不限日期。 格式：yyyy-MM-dd
	StartTime     string `json:"startTime"`     // 可查询此日期及以后下单的礼金效果数据，如不传则不限日期。 格式：yyyy-MM-dd
}

type GiftStatisticResult struct {
}

type OrderQueryRequest struct {
	// 必填 ...

	PageNo   uint64 `json:"pageNo"`   // 页码，返回第几页结果
	PageSize uint64 `json:"pageSize"` // 每页包含条数，上限为500
	TypeNum  uint64 `json:"type"`     // 订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
	Time     string `json:"time"`     // 查询时间，建议使用分钟级查询，格式：yyyyMMddHH、yyyyMMddHHmm或yyyyMMddHHmmss，如201811031212 的查询范围从12:12:00--12:12:59
	// 非必填 ...

	ChildUnionId uint64 `json:"childUnionId"` // 子推客unionID，传入该值可查询子推客的订单，注意不可和key同时传入。（需联系运营开通PID权限才能拿到数据）
	Key          string `json:"key"`          // 工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
}

type OrderQueryResult struct {
}

type OrderBonusRequest struct {
	// 必填 ...

	OptType   uint64 `json:"optType"`   // 时间类型（1.下单时间拉取、2.更新时间拉取）
	StartTime uint64 `json:"startTime"` // 订单开始时间，时间戳（毫秒），与endTime间隔不超过10分钟
	EndTime   uint64 `json:"endTime"`   // 订单结束时间，时间戳（毫秒），与startTime间隔不超过10分钟
	PageSize  uint64 `json:"pageSize"`  // 每页数量，上限100
	// 非必填 ...

	PageNo    uint64 `json:"pageNo"`    // 页码，默认值为1
	SortValue string `json:"sortValue"` // 时间类型按'下单'查询时，和pageSize组合使用。获取最后一条记录的sortValue，指定拉取条数（pageSize），以此方式查询数据。
}

type OrderBonusResult struct {
}

type OrderRowRequest struct {
	// 必填 ...

	PageIndex uint64 `json:"pageIndex"` // 页码
	PageSize  uint64 `json:"pageSize"`  // 每页包含条数，上限为500
	TypeNum   uint64 `json:"type"`      // 订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
	StartTime string `json:"startTime"` // 开始时间 格式yyyy-MM-dd HH:mm:ss，与endTime间隔不超过1小时
	EndTime   string `json:"endTime"`   // 结束时间 格式yyyy-MM-dd HH:mm:ss，与startTime间隔不超过1小时
	// 非必填 ...

	ChildUnionId uint64 `json:"childUnionId"` // 子推客unionID，传入该值可查询子推客的订单，注意不可和key同时传入。（需联系运营开通PID权限才能拿到数据）
	Key          string `json:"key"`          // 工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
	Fields       string `json:"fields"`       // 支持出参数据筛选，逗号','分隔，目前可用：goodsInfo（商品信息）,categoryInfo(类目信息）
}

type OrderRowResult struct {
}
