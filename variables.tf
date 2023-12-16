variable "terratowns_access_token" {
  description = "The accesstoken of the user"
  type        = string
 
}
variable "teacherseat_user_uuid" {
  description = "The UUID of the user"
  type        = string
 
}


variable "terratowns_endpoint" {
  type        = string
}

/*
variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string

 
}

*/
variable "football" {
  type        = object({
    public_path= string
    content_version = number
    
  })
}
variable "cakes" {
  type        = object({
    public_path= string
    content_version = number
    
  })
}
/*

variable "content_version" {
  type    = number 
}


variable "error_html_filepath" {
  type        = string
}


variable "assets_path" {
  type        = string
  
}
*/