package request

import (
	"github.com/switfs/switfs-oms/model/common/request"
	"github.com/switfs/switfs-oms/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
