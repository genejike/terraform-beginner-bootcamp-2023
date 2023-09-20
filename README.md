# Terraform Beginner Bootcamp 2023


## semantic versioning :mage:,:cloud:
This PROJECT IS GOING TO UTILISE SEMANTIC VERSIONING 
[semver.org](https://semver.org/)

Given a version number **MAJOR.MINOR.PATCH ** eg . `1.0.0`, increment the:

* MAJOR version when you make incompatible API changes
* MINOR version when you add functionality in a backward compatible manner
* PATCH version when you make backward compatible bug fixes

Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.

## install Terraform cli
### Terraform cli changes
the terraform cli installations instructions have changed due to the gpg keyring changes so we needed to refer to the latest installation instructions and change the terraform scripting for the install
[install terraform cli](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli#install-terraform)

### Thoughts for linux distributions 
check for the linux distribution to identify any code change according to your distribution needs 
[how to check for linux distribution ](https://www.ionos.com/digitalguide/server/know-how/how-to-check-your-linux-version/)

```
cat /etc/os-release
```

### Restructuring the Bash scripts 
while fixing the installation issues ,the newer codes is a considerable about of code so we decided to use a bash script to install the terraform cli
The bash script:[./bin/install_terraform_cli](./bin/install_terraform_cli)

- The [gitpod.yml](.gitpod.yml) is tidy
- For easier debug sessions and manuel installation of terraform cli 
- this will make it more portable 
### shebang
 shebang tells the bashscript what program that will interpret the script eq `#!/bin/bash`
 we used the `#!/usr/bin/env bash` recommended by chatgpt for portability for diffeerent os distribution
[Shebang](https://en.wikipedia.org/wiki/Shebang_(Unix))

- for portability
- will search the users path for executables 

## Execution Consideration
when executing the bash scripts we can use `./`

eg,`./bin/install_terraform_cli`

if it doesnt work to use the `source ./ `

eg,` source ./bin/install_terraform_cli`

then check the files execute permissions with `ls-la` and run with chmod command to fix it 

```sh
chmod u+x ./bin/install_terraform_cli 
```
[chmod commands](https://en.wikipedia.org/wiki/Chmod)


## GITHUB LIFECYCLE (Before,init,command)
We need need to be careful when using the init cause it wunt restart an existing workspace so we use the `before` instead 
[gitpod lifecycle](https://www.gitpod.io/docs/configure/workspaces/tasks)
