output "bucket_name" {
 description = "Bucket name for our static website hosting"
 value = module.football_home.bucket_name
}

output "s3_website_endpoint" {
 description = "S3 static website hosting endpoint"
 value = module.football_home.website_endpoint
}
output "domain_name" {
    description = "Cloudfront distribution Domain name"
  value = module.football_home.domain_name
}
