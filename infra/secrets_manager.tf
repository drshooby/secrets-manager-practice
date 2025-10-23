# BTW this still puts the secret in the state file...

resource "aws_secretsmanager_secret" "test_secret" {
  name = "test"
}

resource "aws_secretsmanager_secret_version" "example" {
  secret_id     = aws_secretsmanager_secret.test_secret.id
  secret_string = var.test_secret
}
