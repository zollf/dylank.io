resource "aws_s3_bucket" "s3" {
  bucket = "dylank.io"
  acl    = "public-read"
}