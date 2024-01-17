package common

import "github.com/nexient-llc/lcaf-component-terratest-common/types"

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	Vpc_id      string `json:"vpc_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
