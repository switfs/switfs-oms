package service

import (
	"github.com/switfs/switfs-oms/service/example"
	"github.com/switfs/switfs-oms/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
