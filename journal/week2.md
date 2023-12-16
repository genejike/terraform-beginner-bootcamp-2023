# Terraform Beginner Bootcamp 2023 - Week 2
# Table of contents

  - [GITOPS with terraform cloud](#gitops-with-terraform-cloud)
  - [Working with Ruby](#working-with-ruby)
  - [Terratowns Mock Server](#terratowns-mock-server)

## GITOPS with terraform cloud 

## Working with Ruby

### Bundler

Bundler is a package manager for runy.
It is the primary way to install ruby packages (known as gems) for ruby.

#### Install Gems

You need to create a Gemfile and define your gems in that file.

```rb
source "https://rubygems.org"

gem 'sinatra'
gem 'rake'
gem 'pry'
gem 'puma'
gem 'activerecord'
```

Then you need to run the `bundle install` command

This will install the gems on the system globally (unlike nodejs which install packages in place in a folder called node_modules)

A Gemfile.lock will be created to lock down the gem versions used in this project.

#### Executing ruby scripts in the context of bundler

We  use `bundle exec` to tell future ruby scripts to use the gems we installed. This is the way we set context.

### Sinatra

Sinatra is a micro web-framework for ruby to build web-apps.

Its great for mock or development servers or for very simple projects.

You can create a web-server in a single file.

https://sinatrarb.com/

## Terratowns Mock Server

### Running the web server

We can run the web server by executing the following commands:

```rb
bundle install
bundle exec ruby server.rb
```

All of the code for our server is stored in the `server.rb` file.

## Error encoutered while migrating state files back to tf cloud 
* I ran the tf login to connect to my terraform.io .but i wasnt still viewing my state files in cloud 
* I fixed this by providing the block 
```
cloud {
    organization = "MARY"

    workspaces {
      name = "terra-house-blue"
    }
  }

```
in the `/workspace/terraform-beginner-bootcamp-2023/main.tf` as i previously had it in this file `/workspace/terraform-beginner-bootcamp-2023/modules/terrahouse/main.tf`.

* I rotated out my token  in the terraform login prompt and the $TERRAFORM_CLOUD_TOKEN env 

## terrahome_aws
```tf
module "home_cakes" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  content_version = var.content_version
  public_path = var.cakes_public_path
 
  }

```
In this module the public directory expects the following 
- index.html
- error.html
- assets 

all top level files in the assets will be copied, but not any subdirectories

## configuration drift between local state files and terraform cloud 

* i encountered a drift because i ran the `./bin/build_provider`

* i ran `tf init` and `tf state pull` and everything went in place 