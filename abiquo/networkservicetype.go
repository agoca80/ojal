package abiquo

import "github.com/abiquo/opal/core"

type NetworkServiceType struct {
	Default bool   `json:"defaultNST"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	core.DTO
}

func newNetworkServiceType() core.Resource { return new(NetworkServiceType) }
