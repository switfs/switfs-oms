package initialize

import (
	_ "github.com/switfs/switfs-oms/source/example"
	_ "github.com/switfs/switfs-oms/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
