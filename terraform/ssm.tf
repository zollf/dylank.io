locals {
  ssm_prefix = "/ecs/${var.project}"
  environment_secrets = [
    "SECRET_KEY_BASE",
    "MYSQL_HOSTNAME",
    "MYSQL_USER",
    "MYSQL_PASSWORD",
    "MYSQL_DATABASE"
  ]
}

resource "aws_ssm_parameter" "secrets" {
  for_each = toset(local.environment_secrets)
  name        = "${local.ssm_prefix}/${each.key}"
  type        = "SecureString"
  value       = "default"
  overwrite   = true

  tags = {
    terraform = "True"
  }

  lifecycle {
    ignore_changes = [value]
  }
}