package router

import (
	"github.com/switfs/switfs-oms/router/example"
	"github.com/switfs/switfs-oms/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
