# Terraform-beginner-bootcamp-2023 - Week 1

# Table of contents
  - [Root Module structure](#root-module-structure)
  - [Terraform and Input Variables](#terraform-and-input-variables)
  - [Dealing With Configuration Drift](#dealing-with-configuration-drift)
  - [What happens if we lose our state file?](#what-happens-if-we-lose-our-state-file)
  - [Terraform Modules](#terraform-modules)
  - [considerations when using chatgpt to wrte trraform](#considerations-when-using-chatgpt-to-wrte-trraform)
  - [WORKING WITH FILES IN TERRAFORM](#working-with-files-in-terraform)
  - [Provisioners](#provisioners)
  - [FOR EACH EXPRESSIONS](#for-each-expressions)
  
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

### how to delete remote tags 
[How to delete a tag](https://devconnected.com/how-to-delete-local-and-remote-tags-on-git/)
localy delete a tag 

```
 git tag -d <tag_name>
```
remotely delete a tag

```
 git push --delete origin tagname
```

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
### Terraform Locals 
[local values](https://developer.hashicorp.com/terraform/language/values/locals)
```tf
locals {
  service_name = "forum"
  owner        = "Community Team"
}

```
### Terraform Data Sources
This allows us to source data from cloud resources
this is useful when we want to reference cloud resources without importing them 
[data sources](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)
```tf
data "aws_caller_identity" "current" {}

output "account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "caller_arn" {
  value = data.aws_caller_identity.current.arn
}

output "caller_user" {
  value = data.aws_caller_identity.current.user_id
}
```

### working with json
```tf
 jsonencode({"hello"="world"})
{"hello":"world"}

```
we used the json policy to create the json bucket policy inline with the hcl 
[jsonencode](https://developer.hashicorp.com/terraform/language/functions/jsonencode)

when defining my bucket policy i used this 
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowCloudFrontServicePrincipalReadOnly",
            "Effect": "Allow",
            "Principal": {
                "Service": "cloudfront.amazonaws.com"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::DOC-EXAMPLE-BUCKET/*",
            "Condition": {
                "StringEquals": {
                    "AWS:SourceArn": "arn:aws:cloudfront::ACCOUNT_ID:distribution/DISTRIBUTION_ID"
                }
            }
        },
    ]
}
```
for some reason i kept getting this error 
![Alt text](image.png)
i reverted back to an older version and it works fine 
```
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Sid": "AllowCloudFrontServicePrincipal",
            "Effect": "Allow",
            "Principal": {
                "Service": "cloudfront.amazonaws.com"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::DOC-EXAMPLE-BUCKET/*",
            "Condition": {
                "StringEquals": {
                    "AWS:SourceArn": "arn:aws:cloudfront::ACCOUNT_ID:distribution/DISTRIBUTION_ID"
                }
            }
        },
    ]
}
```

### changing the lifecycle of resources 
[lifecycle of resources meta arguements](https://developer.hashicorp.com/terraform/tutorials/state/resource-lifecycle)

### Terraform data 
Plain data values such as Local Values and Input Variables don't have any side-effects to plan against and so they aren't valid in replace_triggered_by. You can use terraform_data's behavior of planning an action each time input changes to indirectly use a plain value to trigger replacement.
[Terraform_data](https://developer.hashicorp.com/terraform/language/resources/terraform-data)
```
variable "revision" {
  default = 1
}

resource "terraform_data" "replacement" {
  input = var.revision
}

# This resource has no convenient attribute which forces replacement,
# but can now be replaced by any change to the revision variable value.
resource "example_database" "test" {
  lifecycle {
    replace_triggered_by = [terraform_data.replacement]
  }
}

```
## Provisioners 
 Provisioners allow to execute commands on compute instances eg. a AWS CLI Command 

 they are not recommended for use by Hashicorp because configuration management tools are better to use but the functionality exits.
 [Provisioners](https://developer.hashicorp.com/terraform/language/resources/provisioners/syntax)

### Local-exec 
In Terraform, the local-exec provisioner is used in a resource block to execute a command on the local machine where Terraform is running, typically for tasks that cannot be accomplished using native Terraform providers. This provisioner is often used for running local scripts, commands, or other actions outside of Terraform's declarative infrastructure management.

```
resource "some_resource" "example" {
  # Resource configuration...

  provisioner "local-exec" {
    command = "echo 'This is a local command'"
  }
}

```
[local-exec](https://developer.hashicorp.com/terraform/language/resources/provisioners/local-exec)
### Remote-exec 
The remote-exec provisioner in Terraform allows you to execute a script or command on a remote resource after it is created. This can be useful for running configuration management tools, bootstrapping a cluster, or other tasks.
This will execute commands on a machine which you target.you will need to provide credentials such as ssh to get into the machine 
[remote _exec ](https://developer.hashicorp.com/terraform/language/resources/provisioners/remote-exec)

## FOR EACH EXPRESSIONS
For each allows us to enumerate ver complex data types 
this is mostly useful when you are creating multiples and u want to reduce the amount of repetitive terraform code 
[For each expressions ](https://developer.hashicorp.com/terraform/language/meta-arguments/for_each)