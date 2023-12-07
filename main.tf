terraform {
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

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.teacherseat_user_uuid
  content_version = var.content_version
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  assets_path = var.assets_path
 
  }


resource "terratowns_home" "home" {
  name = "TIPS TO PLAY THE GAME "
  description= <<DESCRIPTION
 1.PLAY IN PAIRS
 2.ALWAYS ABIDE BY THE RULES
  
DESCRIPTION
  domain_name = module.terrahouse_aws.cloudfront_url
  #domain_name = "3fdqg3gz.cloudfront.net"
  town = "missingo"
  content_version = "1"
   
 }

