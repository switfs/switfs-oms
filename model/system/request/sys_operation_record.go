package request

import (
	"github.com/switfs/switfs-oms/model/common/request"
	"github.com/switfs/switfs-oms/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
