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
