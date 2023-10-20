package common

import "github.com/nexient-llc/tf-caf-terratest-common/types"

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	Vpc_id      string `json:"vpc_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
