resource "aws_route53_record" "dns_record" {
  name    = ""
  type    = "A"
  zone_id = "Z055331928IOJE42QZAS8"

  alias {
    name    = aws_alb.application_load_balancer.dns_name
    zone_id = aws_alb.application_load_balancer.zone_id

    evaluate_target_health = false
  }
}