terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

  backend "s3" {
    bucket = "dylank.io-remote-state"
    key    = "terraform.tfstate"
    region = "ap-southeast-2"
  }
}

data "terraform_remote_state" "remote" {
  backend = "s3"
  config = {
    bucket = "dylank.io-remote-state"
    key    = "terraform.tfstate"
    region = "ap-southeast-2"
  }
}

provider "aws" {
  region = "ap-southeast-2"
}