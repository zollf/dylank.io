resource "aws_ecr_repository" "nginx" {
  name = "${var.project}_nginx_image"
}

resource "aws_ecr_repository" "node" {
  name = "${var.project}_node_image"
}

resource "aws_ecr_repository" "elixir" {
  name = "${var.project}_elixir_image"
}
