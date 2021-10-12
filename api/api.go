package api

import (
	"net/http"

	"github.com/fghwett/icp/abbreviateinfo"
)

func Query(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")

	icp := &abbreviateinfo.Icp{}
	domainInfo, err := icp.Query(domain)
	if err == abbreviateinfo.IcpNotForRecord {
		Json(w, &BaseResponse{Code: CodeOk, Msg: "Success", Data: &QueryResponse{IsRecorded: false}})
		return
	} else if err != nil {
		Json(w, &BaseResponse{Code: CodeErr, Msg: err.Error()})
		return
	}
	Json(w, &BaseResponse{Code: CodeOk, Msg: "Success", Data: &QueryResponse{IsRecorded: true, DomainInfo: domainInfo}})
}
