resource "aws_ecs_task_definition" "task" {
  family                   = "${var.project}_task"
  container_definitions    = <<EOL
  ${jsonencode([
    {
      "name": "${var.project}_elixir",
      "image": "${aws_ecr_repository.elixir.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 8080,
          "hostPort": 8080,
          "protocol": "tcp"
        }
      ],
      "memory": 256,
      "cpu": 96,
      "secrets": [for secret in local.environment_secrets : {
        "name": "${secret}",
        "valueFrom": "${local.ssm_prefix}/${secret}"
      }],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.logs.name}",
          "awslogs-region": "ap-southeast-2",
          "awslogs-stream-prefix": "ecs"
        }
      }
    },
    {
      "name": "${var.project}_nginx",
      "image": "${aws_ecr_repository.nginx.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 80,
          "hostPort": 80,
          "protocol": "tcp"
        }
      ],
      "memory": 128,
      "cpu": 96,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.logs.name}",
          "awslogs-region": "ap-southeast-2",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "dependsOn": [
        {
          "containerName": "${var.project}_elixir",
          "condition": "START"
        },
        {
          "containerName": "${var.project}_node",
          "condition": "START"
        }
      ]
    },
    {
      "name": "${var.project}_node",
      "image": "${aws_ecr_repository.node.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 3000,
          "hostPort": 3000,
          "protocol": "tcp"
        }
      ],
      "memory": 128,
      "cpu": 64,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.logs.name}",
          "awslogs-region": "ap-southeast-2",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ])}
  EOL 
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = "arn:aws:iam::703161335764:role/main-cluster-ecs-task-execution-role"
}
