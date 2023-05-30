package model

// 分页参数
type PageParam struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Params   string `json:"params"`
}

type ResultPage struct {
	Total    int64 `json:"total"`
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
	Data     any   `json:"data"`
}

// PageReq 分页
type PageReq struct {
	PageNum  int `json:"pageNum"  example:"10" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码" form:"pageNum"`
	PageSize int `json:"pageSize" example:"1" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量" form:"pageSize"`
}

// RangeDateReq 时间范围查询
type RangeDateReq struct {
	StartTime string `json:"start_time" v:"date#开始日期格式不正确"  dc:"开始日期" form:"startTime"`
	EndTime   string `json:"end_time" v:"date#结束日期格式不正确" dc:"结束日期" form:"endTime"`
}

// StatusReq 通用状态查询
type StatusReq struct {
	Status int `json:"status"  v:"in:-1,0,1,2,3#输入的状态是无效的" dc:"状态" form:"status"`
}
