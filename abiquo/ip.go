package abiquo

import "github.com/abiquo/opal/core"

// IP represents an Abiquo API IP DTO
type IP struct {
	Available  bool   `json:"available"`
	ProviderID string `json:"providerId"`
	ID         int    `json:"id,omitempty"`
	IPv6       bool   `json:"ipv6"`
	IP         string `json:"ip"`
	core.DTO
}

func newIP() core.Resource { return new(IP) }
