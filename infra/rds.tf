module "db" {
  source = "terraform-aws-modules/rds/aws"

  identifier = "a14db"

  engine               = "mysql"
  engine_version       = "8.0"
  major_engine_version = "8.0"
  family               = "mysql8.0"
  instance_class       = "db.t3.micro"
  allocated_storage    = 5

  db_name                     = "a14db"
  username                    = var.db_username
  password                    = var.db_password
  manage_master_user_password = false
  port                        = "3306"

  iam_database_authentication_enabled = true

  vpc_security_group_ids = [
    aws_security_group.db_sg.id
  ]

  create_db_subnet_group = true
  subnet_ids             = module.vpc.private_subnets

  # Database Deletion Protection
  deletion_protection = false
  skip_final_snapshot = true
}
