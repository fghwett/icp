package abbreviateinfo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIcp_Query(t *testing.T) {
	icp := &Icp{}
	domainInfo, err := icp.Query("baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	info, _ := json.Marshal(domainInfo)
	fmt.Println(string(info))
}
