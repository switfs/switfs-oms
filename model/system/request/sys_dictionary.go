package request

import (
	"github.com/switfs/switfs-oms/model/common/request"
	"github.com/switfs/switfs-oms/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
