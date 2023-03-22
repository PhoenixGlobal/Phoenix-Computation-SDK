package common

// ReqPage page request parameters
type ReqPage struct {
	PageNo   int `form:"pageNo" json:"pageNo" example:"1"`      //page number, non-null
	PageSize int `form:"pageSize" json:"pageSize" example:"20"` //page size, non-null
}
