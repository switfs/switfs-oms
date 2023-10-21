package v1

import (
	"github.com/switfs/switfs-oms/api/v1/example"
	"github.com/switfs/switfs-oms/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
