package main

import (
	"testing"

	testimpl "github.com/nexient-llc/tf-aws-module-private_dns_namespace/tests/testimpl"
	"github.com/nexient-llc/tf-caf-terratest-common/lib"
	"github.com/nexient-llc/tf-caf-terratest-common/types"
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
