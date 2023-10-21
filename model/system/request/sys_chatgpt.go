package request

import (
	"github.com/switfs/switfs-oms/model/common/request"
	"github.com/switfs/switfs-oms/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
