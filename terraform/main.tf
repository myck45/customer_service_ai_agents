terraform {
  backend "s3" {
    bucket         = "terraform-state-restaurant-menu"
    key            = "terraform.tfstate"
    region         = "sa-east-1"
    dynamodb_table = "terraform-state-lock-restaurant-menu"
    encrypt        = true
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.61.0"
    }
  }
}

provider "aws" {
  region = "sa-east-1"
}
