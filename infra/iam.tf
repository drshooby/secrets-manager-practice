# IAM role that can be assumed
resource "aws_iam_role" "app_role" {
  name = "app-secrets-rds-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "secrets_manager_policy" {
  name        = "secrets-manager-access"
  description = "Allow access to Secrets Manager"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_policy" "rds_policy" {
  name        = "rds-access"
  description = "Allow access to RDS"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "rds:DescribeDBInstances",
        ]
        Resource = module.db.db_instance_arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "secrets_attach" {
  role       = aws_iam_role.app_role.name
  policy_arn = aws_iam_policy.secrets_manager_policy.arn
}

resource "aws_iam_role_policy_attachment" "rds_attach" {
  role       = aws_iam_role.app_role.name
  policy_arn = aws_iam_policy.rds_policy.arn
}

resource "aws_iam_instance_profile" "app_profile" {
  name = "app-instance-profile"
  role = aws_iam_role.app_role.name
}
