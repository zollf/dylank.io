resource "aws_s3_bucket" "s3" {
  bucket = "dylank.io"
  acl    = "public-read"

  website {
    index_document = "index.html"
  }
}

data "aws_iam_policy_document" "public_user_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.s3.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.cloudfront_identity.iam_arn]
    }
  }
}

resource "aws_s3_bucket_policy" "s3_policy" {
  bucket = aws_s3_bucket.s3.id
  policy = data.aws_iam_policy_document.public_user_policy.json
}