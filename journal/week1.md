# Terraform-beginner-bootcamp-2023 - Week 1
## Root Module structure 
our Root module structure is as follows:
```
PROJECT_ROOT
├── main.tf              # everything else
├── variables.tf         #stores structure of input variables
├── terraform.tfvars     #the data of variations we load into terraform project
├── providers.tf         #defined required providers and their configuration
├── outputs.tf           # stores our outputs
└── README.md            #required for root mod
```
[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Input Variables

### Terraform Cloud Variables

In terraform we can set two kind of variables:
- Enviroment Variables - these are variables that are set in your bash terminal eg. AWS credentials
- Terraform Variables - these are variables that you would set in your tfvars file eg uuid 

We can set Terraform Cloud variables to be sensitive so they are not shown visibliy in the UI.
so i set my aws credentials in my terraform cloud as an environment variable and set them to sensitive.. i noticed i could not do a terraform destroy without this 
and i also set my uuid as a terraform variable prior to this i kept getting an error because my terraform files where in tf cloud so it couldnt read the uuid i set locally 

### Loading Terraform Input Variables

[Terraform Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)

### var flag
We can use the `-var` flag to set an input variable or override a variable in the tfvars file eg. 
```
terraform -var user_uuid="my-user_id"
```

### var-file flag

- The -var-file flag is used to pass Input Variable values into Terraform plan and apply commands using a file that contains the values.
- This allows you to save the Input Variable values in a file with a .tfvars
- e.g
```
  terraform plan \
-var-file 'production.tfvars' \
-var-file 'secrets.tfvars'
```

### terraform.tvfars

This is the default file to load in terraform variables eg `user_uuid="my user uuid"

### auto.tfvars

The "auto.tfvars" file is intended to store variable values that you want to set automatically without having to pass them as command-line arguments or interactively during the terraform apply or terraform plan commands. It's a convenient way to store sensitive or environment-specific variable values without exposing them in your Terraform configuration or on the command line
You can use this file to set variables like API keys, secret values, or any values that you don't want to hardcode in your main Terraform configuration files.

For example, if you have a variable in your Terraform configuration called aws_access_key and you don't want to hardcode it in your main .tf files, you can define it in "auto.tfvars" eg.` aws_access_key = "your-access-key-value"
`

### order of terraform variables

- Terraform uses a specific order of precedence when determining the value of a variable. If the same variable is assigned multiple values, Terraform will use the value of highest precedence, overriding any other values. Below is the precedence order starting from the highest priority to the lowest.

- Environment variables (TF_VAR_variable_name)
- The terraform.tfvars file
- The terraform.tfvars.json file
- Any .auto.tfvars or .auto.tfvars.json files, processed in lexical order of their filenames.
- Any -var and -var-file options on the command line, in the order they are provided.
- Variable defaults
  
  [blog post by env0 on terraform](https://www.env0.com/blog/terraform-variables)

  ## Dealing With Configuration Drift

## What happens if we lose our state file?

If you lose your statefile, you most likley have to tear down all your cloud infrastructure manually.

You can use terraform port but it won't for all cloud resources. You need check the terraform providers documentation for which resources support import.

### Fix Missing Resources with Terraform Import

`terraform import aws_s3_bucket.bucket bucket-name`

[Terraform Import](https://developer.hashicorp.com/terraform/cli/import)
[AWS S3 Bucket Import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)

### Fix Manual Configuration

If someone goes and delete or modifies cloud resource manually through ClickOps. 

If we run Terraform plan is with attempt to put our infrstraucture back into the expected state fixing Configuration Drift

### Fix using Terraform Refresh

```sh
terraform apply -refresh-only -auto-approve
```

## Terraform Modules

### Terraform Module Structure

It is recommend to place modules in a `modules` directory when locally developing modules but you can name it whatever you like.

### Passing Input Variables

We can pass input variables to our module.
The module has to declare the terraform variables in its own variables.tf

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
```

### Modules Sources

Using the source we can import the module from various places eg:
- locally
- Github
- Terraform Registry

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}
```


[Modules Sources](https://developer.hashicorp.com/terraform/language/modules/sources)

## considerations when using chatgpt to wrte trraform 

A large language model (LLM) such as chatgpt may not be trained on the latest documentation or information about terraform or IACs. it may likely produce examples from older documentations that could be deprecated often affecting  providers
## WORKING WITH FILES IN TERRAFORM
### Fileexists function
Built in terraform function to check existence of a file 
```
condition = fileexists(var.error_html_filepath)
```
https://developer.hashicorp.com/terraform/language/functions/fileexists

### Filemd5

https://developer.hashicorp.com/terraform/language/functions/filemd5

### Working with path in terraform

In terraform there is a special variable called `path` that allows us to ref local path

- path.module = get path for the module
- path.root = get the path for the root module of the configuration 
[Special path variable reference](https://developer.hashicorp.com/terraform/language/expressions/references)


```
resource "aws_s3_object" "error_html" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "error.html"
  source = "${path.root/public/error.html}"
 etag = filemd5("${path.root/public/error.html}")
}
```
terraform console - an interactive window to debug stuff 