package abbreviateinfo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fghwett/icp/model"
	"github.com/fghwett/icp/util"
)

var IcpNotForRecord = errors.New("域名未备案")

type Icp struct {
	token string
}

func (i *Icp) Query(domain string) (*model.DomainInfo, error) {
	if err := i.auth(); err != nil {
		return nil, err
	}

	return i.query(domain)
}

func (i *Icp) query(domain string) (*model.DomainInfo, error) {
	queryRequest, _ := json.Marshal(&model.QueryRequest{
		UnitName: domain,
	})

	result := &model.IcpResponse{Params: &model.QueryParams{}}
	err := i.post("icpAbbreviateInfo/queryByCondition", bytes.NewReader(queryRequest), "application/json;charset=UTF-8", i.token, result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf("查询：%s", result.Msg)
	}

	queryParams := result.Params.(*model.QueryParams)
	if len(queryParams.List) == 0 {
		return nil, IcpNotForRecord
	}

	return queryParams.List[0], nil
}

func (i *Icp) auth() error {
	timestamp := time.Now().Unix()
	authKey := util.Md5(fmt.Sprintf("testtest%d", timestamp))
	authBody := fmt.Sprintf("authKey=%s&timeStamp=%d", authKey, timestamp)

	result := &model.IcpResponse{Params: &model.AuthParams{}}
	err := i.post("auth", bytes.NewReader([]byte(authBody)), "application/x-www-form-urlencoded;charset=UTF-8", "0", result)
	if err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf("获取token失败：%s", result.Msg)
	}

	authParams := result.Params.(*model.AuthParams)
	i.token = authParams.Bussiness

	return nil
}

func (i *Icp) post(url string, data io.Reader, contentType string, token string, result interface{}) error {
	postUrl := fmt.Sprintf("https://hlwicpfwc.miit.gov.cn/icpproject_query/api/%s", url)
	queryReq, err := http.NewRequest(http.MethodPost, postUrl, data)
	if err != nil {
		return err
	}

	queryReq.Header.Set("Content-Type", contentType)
	queryReq.Header.Set("Origin", "https://beian.miit.gov.cn/")
	queryReq.Header.Set("Referer", "https://beian.miit.gov.cn/")
	queryReq.Header.Set("token", token)
	queryReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")

	ip := util.RandIp()
	queryReq.Header.Set("CLIENT_IP", ip)
	queryReq.Header.Set("X-FORWARDED-FOR", ip)

	resp, err := http.DefaultClient.Do(queryReq)
	return util.GetHTTPResponse(resp, postUrl, err, result)
}
