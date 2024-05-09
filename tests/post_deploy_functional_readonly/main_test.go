package main

import (
	"testing"

	testimpl "github.com/launchbynttdata/tf-aws-module_primitive-private_dns_namespace/tests/testimpl"

	"github.com/launchbynttdata/lcaf-component-terratest/lib"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
)

const (
	testConfigsExamplesFolderDefault = "../../examples"
	infraTFVarFileNameDefault        = "test.tfvars"
)

func TestPvtDNSNamespaceModule(t *testing.T) {
	//to implement non destructive tests for this TF module first
	ctx := types.CreateTestContextBuilder().
		SetTestConfig(&testimpl.ThisTFModuleConfig{}).
		SetTestConfigFolderName(testConfigsExamplesFolderDefault).
		SetTestConfigFileName(infraTFVarFileNameDefault).
		SetTestSpecificFlags(map[string]types.TestFlags{
			"complete": {
				"IS_TERRAFORM_IDEMPOTENT_APPLY": true,
			},
		}).
		Build()
	lib.RunNonDestructiveTest(t, *ctx, testimpl.TestComposableComplete)
}
