// What the site can access
// E.g., the site can upload to s3

resource "aws_iam_user" "iam_user" {
  name = "${var.project}-site-user"
}

data "aws_iam_policy_document" "iam_site_policy" {
  statement {
    actions = ["s3:GetObject", "s3:PutObject", "s3:PutObjectAcl"]
    resources = [
      aws_s3_bucket.assets_bucket.arn,
      "${aws_s3_bucket.assets_bucket.arn}/*"
    ]
    effect = "Allow"
  }
}

resource "aws_iam_user_policy" "iam_policy" {
  name = "${var.project}_site_policy"
  user = aws_iam_user.iam_user.name

  policy = data.aws_iam_policy_document.iam_site_policy.json
}
