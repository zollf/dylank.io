resource "aws_s3_bucket" "assets_bucket" {
  bucket = "${var.project}-assets-bucket"
}

resource "aws_s3_bucket_public_access_block" "assets_bucket_public_access" {
  bucket = aws_s3_bucket.assets_bucket.id

  block_public_acls   = true
  block_public_policy = true
}