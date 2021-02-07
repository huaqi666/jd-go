package jd_method

// 推广物料 接口路由
const (
	GoodsJingfenQuery            Method = "jd.union.open.goods.jingfen.query"            // 京粉精选商品查询接口
	GoodsQuery                   Method = "jd.union.open.goods.query"                    // 关键词商品查询接口【申请】
	GoodsMaterialQuery           Method = "jd.union.open.goods.material.query"           // 猜你喜欢商品推荐
	GoodsPromotiongoodsinfoQuery Method = "jd.union.open.goods.promotiongoodsinfo.query" // 根据skuid查询商品信息接口
	CategoryGoodsGet             Method = "jd.union.open.category.goods.get"             // 商品类目查询接口
	GoodsGigfieldQuery           Method = "jd.union.open.goods.bigfield.query"           // 商品详情查询接口【申请】
	CouponQuery                  Method = "jd.union.open.coupon.query"                   // 优惠券领取情况查询接口【申请】
	ActivityQuery                Method = "jd.union.open.activity.query"                 // 活动查询接口
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
