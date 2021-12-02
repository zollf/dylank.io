resource "aws_ecr_repository" "go" {
  name = "${var.project}_go_image"
}

resource "aws_ecr_repository" "nginx" {
  name = "${var.project}_nginx_image"
}

resource "aws_ecr_repository" "node" {
  name = "${var.project}_node_image"
}
