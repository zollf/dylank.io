resource "aws_route53_record" "dns_record" {
  name    = ""
  type    = "A"
  zone_id = "Z055331928IOJE42QZAS8"

  alias {
    name                   = aws_cloudfront_distribution.dist.domain_name
    zone_id                = aws_cloudfront_distribution.dist.hosted_zone_id
    evaluate_target_health = false
  }
}