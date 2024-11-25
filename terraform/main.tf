terraform {
  backend "s3" {
    bucket         = "terraform-state-res-menu"
    key            = "terraform.tfstate"
    region         = "sa-east-1"
    dynamodb_table = "terraform-state-lock-res-menu"
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
