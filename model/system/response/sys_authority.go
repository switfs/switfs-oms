package response

import "github.com/switfs/switfs-oms/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId"` // 旧角色ID
}
