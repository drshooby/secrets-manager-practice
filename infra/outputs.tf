output "ec2-ip" {
  value = aws_instance.webserver.public_ip
}

output "db-uri" {
  value = module.db.db_instance_address
}

output "db-port" {
  value = module.db.db_instance_port
}
