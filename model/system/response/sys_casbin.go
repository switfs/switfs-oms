package response

import (
	"github.com/switfs/switfs-oms/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
