variable "my_ip" {
  type        = string
  default     = "xxx.xxx.xxx/32"
  description = "My IP."
}


variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "VPC CIDR block."
}

variable "vpc_name" {
  type        = string
  default     = "a14-vpc"
  description = "VPC name."
}

variable "aws_azs" {
  description = "List of az in the specified region"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}

variable "public_subnet_cidrs" {
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
  description = "Public subnet CIDR blocks."
}

variable "private_subnet_cidrs" {
  type        = list(string)
  default     = ["10.0.101.0/24", "10.0.102.0/24"]
  description = "Private subnet CIDR blocks."
}

variable "db_username" {
  type = string
}

variable "db_password" {
  sensitive = true
  type      = string
}
