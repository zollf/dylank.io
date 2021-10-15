resource "aws_iam_user" "iam_user" {
  name = "dylank.io-site"
}

data "aws_iam_policy_document" "user_policy" {
  statement {
    actions = ["s3:*"]
    resources = [
      aws_s3_bucket.s3.arn,
      "${aws_s3_bucket.s3.arn}/*"
    ]
    effect = "Allow"
  }
  statement {
    actions = ["s3:GetObject", "s3:PutObject"]
    resources = [
      "arn:aws:s3:::${data.terraform_remote_state.remote.config.bucket}",
      "arn:aws:s3:::${data.terraform_remote_state.remote.config.bucket}/*"
    ]
    effect = "Allow"
  }
  statement {
    actions   = ["iam:GetUser", "iam:GetUserPolicy"]
    resources = ["arn:aws:iam::*:user/${aws_iam_user.iam_user.name}"]
    effect    = "Allow"
  }
}

resource "aws_iam_user_policy" "iam_policy" {
  name = "dylank.io-site-policy"
  user = aws_iam_user.iam_user.name

  policy = data.aws_iam_policy_document.user_policy.json
}