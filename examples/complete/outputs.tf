output "id" {
  description = "ID of the namespace"
  value       = module.namespace.id
}

output "arn" {
  description = "ARN of the namespace"
  value       = module.namespace.arn
}

output "hosted_zone" {
  description = "Hosted Zone for the namespace"
  value       = module.namespace.hosted_zone
}

output "vpc_id" {
  value = module.vpc.vpc_id
}

output "service_name" {
  value = aws_service_discovery_service.example.name
}
output "service_id" {
  value = aws_service_discovery_service.example.id
}
