terraform {
cloud {
    organization = "MARY"

    workspaces {
      name = "terra-house-blue"
    }
  }

  required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"

    }
  }

}
provider "terratowns" {
  endpoint = var.terratowns_endpoint 
  user_uuid = var.teacherseat_user_uuid
  token = var.terratowns_access_token
  
}

module "football_home" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  content_version = var.football.content_version
  public_path = var.football.public_path
 
  }


resource "terratowns_home" "home" {
  name = "TIPS TO PLAY THE GAME "
  description= <<DESCRIPTION
 1.PLAY IN PAIRS
 2.ALWAYS ABIDE BY THE RULES
  
DESCRIPTION
  domain_name = module.football_home.domain_name
  #domain_name = "3fdqg3gz.cloudfront.net"
  town = "missingo"
  content_version = var.football.content_version
   
 }
 

 module "home_cakes" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  content_version = var.cakes.content_version
  public_path = var.cakes.public_path
 
  }



 resource "terratowns_home" "home_cake" {
  name = "How to bake cakes "
  description= <<DESCRIPTION
I love baking goods and cakes are something i will definetly try baking

DESCRIPTION
  domain_name = module.home_cakes.domain_name
  #domain_name = "3fdqg3gz.cloudfront.net"
  town = "missingo"
  content_version = "1"
   
 }

