package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_helper_servicediscovery "github.com/launchbynttdata/lcaf-component-terratest/lib/aws/servicediscovery"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNonComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestIsPrivateServiceDiscoverable", func(t *testing.T) {
		test_service_id := terraform.Output(t, ctx.TerratestTerraformOptions(), "service_id")
		test_service_name := terraform.Output(t, ctx.TerratestTerraformOptions(), "service_name")

		awsServiceDiscoveryClient := test_helper_servicediscovery.GetAwsServiceDiscoveryClient(t)

		svcOut, err2 := awsServiceDiscoveryClient.GetService(context.TODO(), &servicediscovery.GetServiceInput{
			Id: aws.String(test_service_id),
		})

		require.NoError(t, err2, "retrieve AWS example service instance")
		assert.Equal(t, test_service_name, *svcOut.Service.Name)

	})
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestIsNamespaceExist", func(t *testing.T) {
		namespace_id := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")

		awsServiceDiscoveryClient := test_helper_servicediscovery.GetAwsServiceDiscoveryClient(t)
		namespaceOut, err := awsServiceDiscoveryClient.GetNamespace(context.TODO(), &servicediscovery.GetNamespaceInput{
			Id: aws.String(namespace_id),
		})

		require.NoError(t, err, "retrieve AWS namespace")
		assert.Equal(t, ctx.TestConfig().(*ThisTFModuleConfig).Name, *namespaceOut.Namespace.Name)

	})
}
