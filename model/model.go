package model

type QueryRequest struct {
	PageNum  string `json:"pageNum"`
	PageSize string `json:"pageSize"`
	UnitName string `json:"unitName"`
}

type IcpResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Params  interface{} `json:"params"`
}

type AuthParams struct {
	Bussiness string `json:"bussiness"`
	Expire    int64  `json:"expire"`
	Refresh   string `json:"refresh"`
}

type QueryParams struct {
	EndRow           int           `json:"endRow"`
	FirstPage        int           `json:"firstPage"`
	HasNextPage      bool          `json:"hasNextPage"`
	HasPreviousPage  bool          `json:"hasPreviousPage"`
	IsFirstPage      bool          `json:"isFirstPage"`
	IsLastPage       bool          `json:"isLastPage"`
	LastPage         int           `json:"lastPage"`
	List             []*DomainInfo `json:"list"`
	NavigatePages    int           `json:"navigatePages"`
	NavigatepageNums []int         `json:"navigatepageNums"`
	NextPage         int           `json:"nextPage"`
	PageNum          int           `json:"pageNum"`
	PageSize         int           `json:"pageSize"`
	Pages            int           `json:"pages"`
	PrePage          int           `json:"prePage"`
	Size             int           `json:"size"`
	StartRow         int           `json:"startRow"`
	Total            int           `json:"total"`
}

type DomainInfo struct {
	ContentTypeName  string `json:"contentTypeName"`
	Domain           string `json:"domain"`
	DomainId         int64  `json:"domainId"`
	HomeUrl          string `json:"homeUrl"`
	LeaderName       string `json:"leaderName"`
	LimitAccess      string `json:"limitAccess"`
	MainId           int64  `json:"mainId"`
	MainLicence      string `json:"mainLicence"`
	NatureName       string `json:"natureName"`
	ServiceId        int64  `json:"serviceId"`
	ServiceLicence   string `json:"serviceLicence"`
	ServiceName      string `json:"serviceName"`
	UnitName         string `json:"unitName"`
	UpdateRecordTime string `json:"updateRecordTime"`
}
