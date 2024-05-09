# Cloud Map private Namespace setup example
This module will provision a Cloud Map private Namespace. It will provision all dependent resources as shown below
- VPC

## Pre-requisites
- Need to create a `provider.tf` with below contents
    ```
    provider "aws" {
      profile = "<profile-name>"
      region  = "<aws-region>"
    }
    ```
- Command to execute the module
  `terraform apply -var-file test.tfvars`
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0, <= 1.5.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.48.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform-aws-modules/vpc/aws | ~> 5.1.1 |
| <a name="module_namespace"></a> [namespace](#module\_namespace) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_service_discovery_service.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/service_discovery_service) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_vpc_name"></a> [vpc\_name](#input\_vpc\_name) | n/a | `string` | `"test-vpc-015935234"` | no |
| <a name="input_vpc_cidr"></a> [vpc\_cidr](#input\_vpc\_cidr) | n/a | `string` | `"10.1.0.0/16"` | no |
| <a name="input_private_subnets"></a> [private\_subnets](#input\_private\_subnets) | List of private subnet cidrs | `list(string)` | <pre>[<br>  "10.1.1.0/24",<br>  "10.1.2.0/24",<br>  "10.1.3.0/24"<br>]</pre> | no |
| <a name="input_availability_zones"></a> [availability\_zones](#input\_availability\_zones) | List of availability zones for the VPC | `list(string)` | <pre>[<br>  "us-east-2a",<br>  "us-east-2b",<br>  "us-east-2c"<br>]</pre> | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the namespace. For example example.local | `string` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | Description about the namespace | `string` | `""` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | ID of the namespace |
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the namespace |
| <a name="output_hosted_zone"></a> [hosted\_zone](#output\_hosted\_zone) | Hosted Zone for the namespace |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | n/a |
| <a name="output_service_name"></a> [service\_name](#output\_service\_name) | n/a |
| <a name="output_service_id"></a> [service\_id](#output\_service\_id) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
