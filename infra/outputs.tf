output "ec2-ip" {
  value = aws_instance.webserver.public_ip
}
