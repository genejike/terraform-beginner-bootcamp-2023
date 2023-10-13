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
