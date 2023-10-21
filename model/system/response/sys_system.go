package response

import "github.com/switfs/switfs-oms/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
