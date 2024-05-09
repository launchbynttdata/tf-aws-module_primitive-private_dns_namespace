package common

import "github.com/launchbynttdata/lcaf-component-terratest/types"

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	Vpc_id      string `json:"vpc_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
