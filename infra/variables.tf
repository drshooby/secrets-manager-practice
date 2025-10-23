variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "test_secret" {
  type      = string
  sensitive = true
}
