package jd_method

// 推广效果 接口路由方法
const (
	OrderQuery                    Method = "jd.union.open.order.query"                      // Deprecated: 订单查询接口【即将下线】
	OrderRowQuery                 Method = "jd.union.open.order.row.query"                  // 订单行查询接口
	OrderBonusQuery               Method = "jd.union.open.order.bonus.query"                // 奖励订单查询接口【申请】
	StatisticsRedpacketQuery      Method = "jd.union.open.statistics.redpacket.query"       // 京享红包效果数据
	OrderAgentQuery               Method = "jd.union.open.order.agent.query"                // 工具商订单行查询接口【申请】
	StatisticsRedpacketAgentQuery Method = "jd.union.open.statistics.redpacket.agent.query" // 工具商京享红包效果数据查询接口【申请】
)
