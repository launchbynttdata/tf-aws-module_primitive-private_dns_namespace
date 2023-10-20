### VPC related variables

variable "vpc_name" {
  default = "test-vpc-015935234"
}

variable "vpc_cidr" {
  default = "10.1.0.0/16"
}

variable "private_subnets" {
  description = "List of private subnet cidrs"
  default     = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
}

variable "availability_zones" {
  description = "List of availability zones for the VPC"
  default     = ["us-east-2a", "us-east-2b", "us-east-2c"]
}

### Namespace related variables
variable "name" {
  description = "Name of the namespace. For example example.local"
  type        = string
}

variable "description" {
  description = "Description about the namespace"
  default     = ""
  type        = string
}

variable "tags" {
  default = {}
}
