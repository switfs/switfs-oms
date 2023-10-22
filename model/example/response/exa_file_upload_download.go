package response

import "github.com/switfs/switfs-oms/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
