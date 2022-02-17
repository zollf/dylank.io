resource "aws_s3_bucket" "assets_bucket" {
  bucket = "${var.project}-assets-bucket"
}

resource "aws_s3_bucket_public_access_block" "assets_bucket_public_access" {
  bucket = aws_s3_bucket.assets_bucket.id

  block_public_acls   = false
  block_public_policy = false
}

resource "aws_s3_bucket_policy" "public_access_policy" {
  bucket = aws_s3_bucket.assets_bucket.id
  policy = data.aws_iam_policy_document.public_access_document.json
}

data "aws_iam_policy_document" "public_access_document" {
  statement {
    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:GetObject",
    ]

    resources = [
      aws_s3_bucket.assets_bucket.arn,
      "${aws_s3_bucket.assets_bucket.arn}/*",
    ]
  }
}