package api

import (
	"encoding/json"
	"net/http"

	"github.com/fghwett/icp/model"
)

const (
	CodeOk  = 0
	CodeErr = -1
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type QueryResponse struct {
	IsRecorded bool `json:"isRecorded"`
	*model.DomainInfo
}

func Json(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(o)
	w.Write(b)
}
