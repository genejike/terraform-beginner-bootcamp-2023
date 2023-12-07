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
variable "index_html_filepath" {
  type        = string
}

variable "error_html_filepath" {
  type        = string
}

variable "content_version" {
  type    = number 
}

variable "assets_path" {
  type        = string
  
}