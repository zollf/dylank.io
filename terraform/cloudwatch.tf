resource "aws_cloudwatch_log_group" "logs" {
  name              = "/ecs/service/${var.project}"
  retention_in_days = 7
}