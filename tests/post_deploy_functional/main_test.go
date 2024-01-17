package main

import (
	"testing"

	"github.com/nexient-llc/lcaf-component-terratest-common/lib"
	"github.com/nexient-llc/lcaf-component-terratest-common/types"
	testimpl "github.com/nexient-llc/tf-aws-module-private_dns_namespace/tests/testimpl"
)

const (
	testConfigsExamplesFolderDefault = "../../examples"
	infraTFVarFileNameDefault        = "test.tfvars"
)

func TestPvtDNSNamespaceModule(t *testing.T) {

	ctx := types.TestContext{
		TestConfig: &testimpl.ThisTFModuleConfig{},
	}
	lib.RunSetupTestTeardown(t, testConfigsExamplesFolderDefault, infraTFVarFileNameDefault, ctx,
		testimpl.TestNonComposableComplete,
		testimpl.TestComposableComplete)
}
