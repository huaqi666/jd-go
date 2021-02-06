package jd

// 请求的方法(接口)
type Method string

func (m Method) String() string {
	return string(m)
}

type ResponseKey string

func (m ResponseKey) String() string {
	return string(m)
}

// 商品接口
const (
	GoodsJingfenQuery            Method = "jd.union.open.goods.jingfen.query"            // 精选商品查询
	GoodsQuery                   Method = "jd.union.open.goods.query"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQuery Method = "jd.union.open.goods.promotiongoodsinfo.query" // 根据skuid查询商品信息接口
	CategoryGoodsGet             Method = "jd.union.open.category.goods.get"             // 商品类目查询接口
	GoodsGigfieldQuery           Method = "jd.union.open.goods.bigfield.query"           // 商品详情查询接口【申请】
)
const (
	GoodsJingfenQueryResponse            ResponseKey = "jd_union_open_goods_jingfen_query_response"            // 精选商品查询
	GoodsQueryResponse                   ResponseKey = "jd_union_open_goods_query_response"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQueryResponse ResponseKey = "jd_union_open_goods_promotiongoodsinfo_query_response" // 根据skuid查询商品信息接口
	CategoryGoodsGetResponse             ResponseKey = "jd_union_open_category_goods_get_response"             // 商品类目查询接口
	GoodsGigfieldQueryResponse           ResponseKey = "jd_union_open_goods_bigfield_query_response"           // 商品详情查询接口【申请】
)
const (
	GoodsJingfenQueryResponce            ResponseKey = "jd_union_open_goods_jingfen_query_responce"            // 精选商品查询
	GoodsQueryResponce                   ResponseKey = "jd_union_open_goods_query_responce"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQueryResponce ResponseKey = "jd_union_open_goods_promotiongoodsinfo_query_responce" // 根据skuid查询商品信息接口
	CategoryGoodsGetResponce             ResponseKey = "jd_union_open_category_goods_get_responce"             // 商品类目查询接口
	GoodsGigfieldQueryResponce           ResponseKey = "jd_union_open_goods_bigfield_query_responce"           // 商品详情查询接口【申请】
)

// 优惠价接口
const (
	CouponQuery Method = "jd.union.open.coupon.query" // 优惠券领取情况查询接口【申请】
)
const (
	CouponQueryResponse ResponseKey = "jd_union_open_coupon_query_response" // 优惠券领取情况查询接口【申请】
)
const (
	CouponQueryResponce ResponseKey = "jd_union_open_coupon_query_responce" // 优惠券领取情况查询接口【申请】
)

// 推广接口
const (
	PromotionCommonGet       Method = "jd.union.open.promotion.common.get"       // 网站/APP获取推广链接接口
	PromotionBysubunionidGet Method = "jd.union.open.promotion.bysubunionid.get" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGet    Method = "jd.union.open.promotion.byunionid.get"    // 工具商获取推广链接接口【申请】
	PromotionAppletGet       Method = "jd.union.open.promotion.applet.get"       // 通过小程序获取推广链接【申请】

	StatisticsRedpacketAgentQuery Method = "jd.union.open.statistics.redpacket.agent.query" // 工具商京享红包效果数据查询接口【申请】
)
const (
	PromotionCommonGetResponse       ResponseKey = "jd_union_open_promotion_common_get_response"       // 网站/APP获取推广链接接口
	PromotionBysubunionidGetResponse ResponseKey = "jd_union_open_promotion_bysubunionid_get_response" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGetResponse    ResponseKey = "jd_union_open_promotion_byunionid_get_response"    // 工具商获取推广链接接口【申请】
	PromotionAppletGetResponse       ResponseKey = "jd_union_open_promotion_applet_get_response"       // 通过小程序获取推广链接【申请】

	StatisticsRedpacketAgentQueryResponse Method = "jd_union_open_statistics_redpacket_agent_query_response" // 工具商京享红包效果数据查询接口【申请】
)
const (
	PromotionCommonGetResponce       ResponseKey = "jd_union_open_promotion_common_get_responce"       // 网站/APP获取推广链接接口
	PromotionBysubunionidGetResponce ResponseKey = "jd_union_open_promotion_bysubunionid_get_responce" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGetResponce    ResponseKey = "jd_union_open_promotion_byunionid_get_responce"    // 工具商获取推广链接接口【申请】
	PromotionAppletGetResponce       ResponseKey = "jd_union_open_promotion_applet_get_responce"       // 通过小程序获取推广链接【申请】

	StatisticsRedpacketAgentQueryResponce Method = "jd_union_open_statistics_redpacket_agent_query_responce" // 工具商京享红包效果数据查询接口【申请】
)

// 订单接口
const (
	OrderQuery      Method = "jd.union.open.order.query"       // 订单查询接口
	OrderBonusQuery Method = "jd.union.open.order.bonus.query" // 奖励订单查询接口【申请】
	OrderRowQuery   Method = "jd.union.open.order.row.query"   // 订单行查询接口
)
const (
	OrderQueryResponse      ResponseKey = "jd_union_open_order_query_response"       // 订单查询接口
	OrderBonusQueryResponse ResponseKey = "jd_union_open_order_bonus_query_response" // 奖励订单查询接口【申请】
	OrderRowQueryResponse   ResponseKey = "jd_union_open_order_row_query_response"   // 订单行查询接口
)
const (
	OrderQueryResponce      ResponseKey = "jd_union_open_order_query_responce"       // 订单查询接口
	OrderBonusQueryResponce ResponseKey = "jd_union_open_order_bonus_query_responce" // 奖励订单查询接口【申请】
	OrderRowQueryResponce   ResponseKey = "jd_union_open_order_row_query_responce"   // 订单行查询接口
)

// 推广位接口
const (
	PositionQuery  Method = "jd.union.open.position.query"  // 查询推广位【申请】
	PositionCreate Method = "jd.union.open.position.create" // 创建推广位【申请】
	UserPidGet     Method = "jd.union.open.user.pid.get"    // 获取PID【申请】
)
const (
	PositionQueryResponse  ResponseKey = "jd_union_open_position_query_response"  // 查询推广位【申请】
	PositionCreateResponse ResponseKey = "jd_union_open_position_create_response" // 创建推广位【申请】
	UserPidGetResponse     ResponseKey = "jd_union_open_user_pid_get_response"    // 获取PID【申请】
)
const (
	PositionQueryResponce  ResponseKey = "jd_union_open_position_query_responce"  // 查询推广位【申请】
	PositionCreateResponce ResponseKey = "jd_union_open_position_create_responce" // 创建推广位【申请】
	UserPidGetResponce     ResponseKey = "jd_union_open_user_pid_get_responce"    // 获取PID【申请】
)

// 营销接口
const (
	CouponGiftStop           Method = "jd.union.open.coupon.gift.stop"            // 礼金停止
	CouponGiftGet            Method = "jd.union.open.coupon.gift.get"             // 礼金创建
	StatisticGiftCouponQuery Method = "jd.union.open.statistics.giftcoupon.query" // 礼金效果数据
)
const (
	CouponGiftStopResponse           ResponseKey = "jd_union_open_coupon_gift_stop_response"            // 礼金停止
	CouponGiftGetResponse            ResponseKey = "jd_union_open_coupon_gift_get_response"             // 礼金创建
	StatisticGiftCouponQueryResponse ResponseKey = "jd_union_open_statistics_giftcoupon_query_response" // 礼金效果数据
)
const (
	CouponGiftStopResponce           ResponseKey = "jd_union_open_coupon_gift_stop_responce"            // 礼金停止
	CouponGiftGetResponce            ResponseKey = "jd_union_open_coupon_gift_get_responce"             // 礼金创建
	StatisticGiftCouponQueryResponce ResponseKey = "jd_union_open_statistics_giftcoupon_query_responce" // 礼金效果数据
)

// 活动接口
const (
	ActivityQuery            Method = "jd.union.open.activity.query"             // 活动查询接口
	StatisticsRedpacketQuery Method = "jd.union.open.statistics.redpacket.query" // 京享红包效果数据
)
const (
	ActivityQueryResponse            ResponseKey = "jd_union_open_activity_query_response"             // 活动查询接口
	StatisticsRedpacketQueryResponse ResponseKey = "jd_union_open_statistics_redpacket_query_response" // 京享红包效果数据
)
const (
	ActivityQueryResponce            ResponseKey = "jd_union_open_activity_query_responce"             // 活动查询接口
	StatisticsRedpacketQueryResponce ResponseKey = "jd_union_open_statistics_redpacket_query_responce" // 京享红包效果数据
)

const (
	UnionRootEndpoint = "https://router.jd.com/api"     // 京东联盟API路由: 参数: param_json
	JosRootEndpoint   = "https://api.jd.com/routerjson" // 宙斯API路由 🤣: 参数: 360buy_param_json
	BaseUrl           = UnionRootEndpoint
)

// 参数验证/校验等级
type VCode int

const (
	Non VCode = iota // 不校验😅
	NotEmpty
)
