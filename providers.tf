terraform {
  
    cloud {
    organization = "MARY"

    workspaces {
      name = "terra-house-blue"
    }
  }

  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.20.0"
    }
  }
}

provider "aws" {
  # Configuration options

}

provider "random" {
  # Configuration options
}