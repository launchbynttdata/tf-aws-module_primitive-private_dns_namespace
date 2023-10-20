module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.1.1"

  name                 = var.vpc_name
  cidr                 = var.vpc_cidr
  private_subnets      = var.private_subnets
  azs                  = var.availability_zones
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = var.tags
}

module "namespace" {
  source = "../.."

  vpc_id      = module.vpc.vpc_id
  name        = var.name
  description = var.description
  tags        = var.tags
}

resource "aws_service_discovery_service" "example" {
  name = "example"

  dns_config {
    namespace_id = module.namespace.id

    dns_records {
      ttl  = 10
      type = "A"
    }

    routing_policy = "MULTIVALUE"
  }

  health_check_custom_config {
    failure_threshold = 1
  }
}
