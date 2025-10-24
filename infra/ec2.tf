resource "aws_instance" "webserver" {
  ami                    = "ami-052064a798f08f0d3" # Amazon Linux 2
  instance_type          = "t2.micro"
  subnet_id              = module.vpc.public_subnets[0]
  key_name               = "cloud-computing-kp"
  vpc_security_group_ids = [aws_security_group.webserver_sg.id]

  associate_public_ip_address = true

  iam_instance_profile = aws_iam_instance_profile.app_profile.name

  tags = {
    Name = "webserver"
  }
}
