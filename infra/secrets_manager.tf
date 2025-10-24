#BTW this still puts the secret in the state file...

resource "aws_secretsmanager_secret" "db_user_secret" {
  name = "DB_USER"
}

resource "aws_secretsmanager_secret_version" "db_user_sv" {
  secret_id     = aws_secretsmanager_secret.db_user_secret.id
  secret_string = var.db_username
}

resource "aws_secretsmanager_secret" "db_password_secret" {
  name = "DB_PASSWORD"
}

resource "aws_secretsmanager_secret_version" "db_password_sv" {
  secret_id     = aws_secretsmanager_secret.db_password_secret.id
  secret_string = var.db_password
}
