resource "aws_iam_user" "iam_user" {
  name = "${var.project}-site-user"
}

data "aws_iam_policy_document" "iam_site_policy" {
  statement {
    actions = ["s3:GetObject", "s3:PutObject"]
    resources = [
      "arn:aws:s3:::${data.terraform_remote_state.remote.config.bucket}",
      "arn:aws:s3:::${data.terraform_remote_state.remote.config.bucket}/*"
    ]
    effect = "Allow"
  }
  statement {
    actions = [
      # Users
      "iam:GetUserPolicy",
      "iam:ListGroupsForUser",
      "iam:ListAttachedUserPolicies",
      "iam:ListUserPolicies",
      "iam:GetUser",
      "iam:GetGroupPolicy",
      "iam:GetPolicyVersion",
      "iam:GetPolicy",
      "iam:ListAttachedGroupPolicies",
      "iam:ListGroupPolicies",
      "iam:ListPolicyVersions",
      "iam:ListPolicies",
      "iam:ListUsers",
      # Roles
      "iam:GetRole",
      "iam:ListRolePolicies",
      "iam:ListAttachedRolePolicies",
    ]
    resources = [aws_iam_user.iam_user.arn, aws_iam_role.ecs_task_execution_role.arn]
    effect    = "Allow"
  }

  statement {
    actions = ["ecr:*"]
    effect  = "Allow"
    resources = [
      aws_ecr_repository.go.arn,
      aws_ecr_repository.node.arn,
      aws_ecr_repository.nginx.arn
    ]
  }

  statement {
    actions = [
      "ecr:GetAuthorizationToken"
    ]
    effect    = "Allow"
    resources = ["*"]
  }

  statement {
    actions   = ["ec2:*"]
    effect    = "Allow"
    resources = ["*"]
  }

  statement {
    actions   = ["ecs:*"]
    effect    = "Allow"
    resources = ["*"]
  }

  statement {
    actions   = ["elasticloadbalancing:*"]
    effect    = "Allow"
    resources = ["*"]
  }

  statement {
    actions   = ["route53:*"]
    effect    = "Allow"
    resources = ["*"]
  }
}

resource "aws_iam_user_policy" "iam_policy" {
  name = "${var.project}_site_policy"
  user = aws_iam_user.iam_user.name

  policy = data.aws_iam_policy_document.iam_site_policy.json
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "${var.project}_ecs_execution_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}