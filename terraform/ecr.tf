resource "aws_ecr_repository" "go" {
  name = "${var.project}_go_image"
}
