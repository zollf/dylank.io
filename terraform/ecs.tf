// My main cluster
resource "aws_ecs_cluster" "cluster" {
  name = var.cluster
}

resource "aws_ecs_task_definition" "task" {
  family                   = "${var.project}_task"
  container_definitions    = <<DEFINITION
  [
    {
      "name": "${var.project}_go",
      "image": "${aws_ecr_repository.go.repository_url}",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 8080,
          "hostPort": 8080,
          "protocol": "tcp"
        }
      ],
      "memory": 171,
      "cpu": 85,
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
      "memory": 170,
      "cpu": 85,
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
      "memory": 171,
      "cpu": 86,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.logs.name}",
          "awslogs-region": "ap-southeast-2",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
}

resource "aws_ecs_service" "ecs_service" {
  name                              = "${var.project}_service"
  cluster                           = aws_ecs_cluster.cluster.id
  task_definition                   = aws_ecs_task_definition.task.arn
  launch_type                       = "FARGATE"
  desired_count                     = 1
  health_check_grace_period_seconds = 30

  network_configuration {
    subnets = [
      "${aws_default_subnet.default_subnet_a.id}",
      "${aws_default_subnet.default_subnet_b.id}",
      "${aws_default_subnet.default_subnet_c.id}"
    ]
    assign_public_ip = true
    security_groups  = ["${aws_security_group.service_security_group.id}"]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn
    container_name   = "${var.project}_nginx"
    container_port   = 80
  }
}
