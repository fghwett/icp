package icp

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

/*
项目地址：https://github.com/yitd/ICP-API
*/

type AuthResponse struct {
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
	EndRow          int  `json:"endRow"`
	FirstPage       int  `json:"firstPage"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	IsFirstPage     bool `json:"isFirstPage"`
	IsLastPage      bool `json:"isLastPage"`
	LastPage        bool `json:"lastPage"`
	List            []struct {
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
	} `json:"list"`
	NavigatePages    int   `json:"navigatePages"`
	NavigatepageNums []int `json:"navigatepageNums"`
	NextPage         int   `json:"nextPage"`
	PageNum          int   `json:"pageNum"`
	PageSize         int   `json:"pageSize"`
	Pages            int   `json:"pages"`
	PrePage          int   `json:"prePage"`
	Size             int   `json:"size"`
	StartRow         int   `json:"startRow"`
	Total            int   `json:"total"`
}

type QueryRequest struct {
	PageNum  string `json:"pageNum"`
	PageSize string `json:"pageSize"`
	UnitName string `json:"unitName"`
}

func Query(domain string) {
	timestamp := time.Now().Unix()
	authKey := Md5(fmt.Sprintf("testtest%d", timestamp))
	authBody := fmt.Sprintf("authKey=%s&timeStamp=%d", authKey, timestamp)

	req, err := http.NewRequest(http.MethodPost, "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/auth", bytes.NewReader([]byte(authBody)))
	if err != nil {
		fmt.Println("auth:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Origin", "https://beian.miit.gov.cn/")
	req.Header.Set("Referer", "https://beian.miit.gov.cn/")
	req.Header.Set("token", "0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")

	ip := RandIp()
	req.Header.Set("CLIENT_IP", ip)
	req.Header.Set("X-FORWARDED-FOR", ip)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("auth: ", err)
		return
	}

	authRespBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("auth: ", err)
		return
	}

	// {"success":false,"code":500,"msg":"服务器异常"}
	/*
		{
			"code": 200,
			"msg": "操作成功",
			"params": {
				"bussiness": "eyJ0eXBlIjoxLCJ1IjoiMDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjYiLCJzIjoxNjMzOTUyOTI4NzIxLCJlIjoxNjMzOTUzNDA4NzIxfQ.9jJWrc2L1IwD4I_vs8p9O0oFFG6RUjIqda5Ubz2nZn4",
				"expire": 300000,
				"refresh": "eyJ0eXBlIjoyLCJ1IjoiMDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjYiLCJzIjoxNjMzOTUyOTI4NzIxLCJlIjoxNjMzOTUzNzA4NzIxfQ.r1vTT-MN3EquWVdshOlehr7caK4X2D59FAz3vjZjkNc"
			},
			"success": true
		}
		2021-10-11 19:48:48 s参数 1633952928721 2021-10-11 19:48:48
		token有效期：2021-10-11 19:56:48 e参数 1633953408721-1633952928721 8分钟有效期
	*/
	fmt.Println("body:", string(authRespBody))

	authResp := &AuthResponse{
		Params: &AuthParams{},
	}
	if err := json.Unmarshal(authRespBody, authResp); err != nil {
		fmt.Println("auth: ", err)
		return
	}

	if !authResp.Success {
		fmt.Println("auth: ", authResp.Msg)
		return
	}

	// ----
	queryRequest, err := json.Marshal(&QueryRequest{
		UnitName: domain,
	})
	if err != nil {
		fmt.Println("query: ", err)
		return
	}

	queryReq, err := http.NewRequest(http.MethodPost, "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/icpAbbreviateInfo/queryByCondition", bytes.NewReader(queryRequest))
	if err != nil {
		fmt.Println("query:", err)
		return
	}

	queryReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	queryReq.Header.Set("Origin", "https://beian.miit.gov.cn/")
	queryReq.Header.Set("Referer", "https://beian.miit.gov.cn/")
	queryReq.Header.Set("token", authResp.Params.(*AuthParams).Bussiness)
	queryReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")
	queryReq.Header.Set("CLIENT_IP", ip)
	queryReq.Header.Set("X-FORWARDED-FOR", ip)

	queryResp, err := http.DefaultClient.Do(queryReq)
	if err != nil {
		fmt.Println("query: ", err)
		return
	}

	queryRespBody, err := ioutil.ReadAll(queryResp.Body)
	if err != nil {
		fmt.Println("query: ", err)
		return
	}

	/*
		{
		    "code": 200,
		    "msg": "操作成功",
		    "params": {
		        "endRow": 0,
		        "firstPage": 1,
		        "hasNextPage": false,
		        "hasPreviousPage": false,
		        "isFirstPage": true,
		        "isLastPage": true,
		        "lastPage": 1,
		        "list": [
		            {
		                "contentTypeName": "",
		                "domain": "mi.cn",
		                "domainId": 10004219290,
		                "homeUrl": "www.mi.cn",
		                "leaderName": "",
		                "limitAccess": "否",
		                "mainId": 6504864,
		                "mainLicence": "京ICP备10046444号",
		                "natureName": "企业",
		                "serviceId": 10001392064,
		                "serviceLicence": "京ICP备10046444号-9",
		                "serviceName": "小米科技",
		                "unitName": "小米科技有限责任公司",
		                "updateRecordTime": "2021-08-16 13:55:56"
		            }
		        ],
		        "navigatePages": 8,
		        "navigatepageNums": [
		            1
		        ],
		        "nextPage": 1,
		        "pageNum": 1,
		        "pageSize": 10,
		        "pages": 1,
		        "prePage": 1,
		        "size": 1,
		        "startRow": 0,
		        "total": 1
		    },
		    "success": true
		}
	*/
	fmt.Println(string(queryRespBody))

	/*
		详细信息获得到的数据，列表信息里面都有
		详细信息 https://hlwicpfwc.miit.gov.cn/icpproject_query/api/icpAbbreviateInfo/queryDetailByServiceIdAndDomainId
		{
			"mainId": 6504864,
			"domainId": 10004219290,
			"serviceId": 10001392064
		}
		{
			"code": 200,
			"msg": "操作成功",
			"params": {
				"contentTypeName": "",
				"domain": "mi.cn",
				"domainId": 10004219290,
				"homeUrl": "www.mi.cn",
				"leaderName": "",
				"mainId": 6504864,
				"mainLicence": "京ICP备10046444号",
				"natureName": "企业",
				"serviceId": 10001392064,
				"serviceLicence": "京ICP备10046444号-9",
				"serviceName": "小米科技",
				"unitName": "小米科技有限责任公司",
				"updateRecordTime": "2021-08-16 13:55:56"
			},
			"success": true
		}
	*/
}

func PostAuth(url string, data io.Reader, contentType string, token string) {

}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
func RandIp() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("101.%d.%d.%d", 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254))
}

