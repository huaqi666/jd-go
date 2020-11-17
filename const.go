package jd

// 请求的方法(接口)
type Method string

const (
	GoodsJingfenQuery            Method = "jd.union.open.goods.jingfen.query"            // 精选商品查询
	GoodsQuery                   Method = "jd.union.open.goods.query"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQuery Method = "jd.union.open.goods.promotiongoodsinfo.query" // 根据skuid查询商品信息接口
	CategoryGoodsQuery           Method = "jd.union.open.category.goods.get"             // 商品类目查询接口
	GoodsGigfieldQuery           Method = "jd.union.open.goods.bigfield.query"           // 商品详情查询接口【申请】
)

const (
	CouponQuery Method = "jd.union.open.coupon.query" // 优惠券领取情况查询接口【申请】
)

const (
	PromotionCommonGet         Method = "jd.union.open.promotion.common.get"       // 网站/APP获取推广链接接口
	PromotionBysubunionidQuery Method = "jd.union.open.promotion.bysubunionid.get" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidQuery    Method = "jd.union.open.promotion.byunionid.get"    // 工具商获取推广链接接口【申请】
)

const (
	OrderQuery      Method = "jd.union.open.order.query"       // 订单查询接口
	OrderBonusQuery Method = "jd.union.open.order.bonus.query" // 奖励订单查询接口【申请】
	OrderRowQuery   Method = "jd.union.open.order.row.query"   // 订单行查询接口
)

const (
	PositionQuery  Method = "jd.union.open.position.query"  // 查询推广位【申请】
	PositionCreate Method = "jd.union.open.position.create" // 创建推广位【申请】
	UserPidGet     Method = "jd.union.open.user.pid.get"    // 获取PID【申请】
)

const (
	GiftStop      Method = "jd.union.open.coupon.gift.stop"            // 礼金停止
	GiftGet       Method = "jd.union.open.coupon.gift.get"             // 礼金创建
	GiftStatistic Method = "jd.union.open.statistics.giftcoupon.query" // 礼金效果数据
)

const (
	ActivityQuery            Method = "jd.union.open.activity.query"             // 活动查询接口
	StatisticsRedpacketQuery Method = "jd.union.open.statistics.redpacket.query" // 京享红包效果数据
)

const (
	BaseUrl = "https://router.jd.com/api"
)
