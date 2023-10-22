package response

import "github.com/switfs/switfs-oms/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
