# Azure Terraform Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/snimmagadda1/azure-terraform-generator)](https://goreportcard.com/report/github.com/snimmagadda1/azure-terraform-generator)

Azure API -> Terraform Resources

## Overview

---

This is a work in progress initiative to create a CLI tool to generate Terraform definitions for Azure resources. The idea for this came about when teams I've worked with had existing Azure resources but were lacking terraform definitions, subsequently slowing down migrations & time to production.

## Usage
Simply clone the repository or install with `go get` & install to expose the cli:
```
git clone https://github.com/snimmagadda1/azure-terraform-generator.git
go install
```

As a prerequisite to using the azure apis, credentials must be provided. Currently a file-based credential method is supported. As this project is built out, more authorization methods will be supported.

The simplest way to obtain this file is throughthe creation of an Azure service principal. Either download the credentials of an existing service principal to a file my.auth in the root directory or create a new service principal as follows:

```
az ad sp create-for-rbac --sdk-auth > my.auth
```

Next set an environment variable `AZURE_AUTH_LOCATION` to the location of my.auth:

```
export AZURE_AUTH_LOCATION=$HOME/my.auth
```

### Commands

```
Usage:
  azure-terraform-generator [command]

Available Commands:
  group       Generate terraform resource for an Azure ResourceGroup
  help        Help about any command

Flags:
  -h, --help   help for azure-terraform-generator

Use "azure-terraform-generator [command] --help" for more information about a command.
```

Currently supported Azure Objects -> Terraform resources:

1. [ResourceGroup](https://www.terraform.io/docs/providers/azurerm/d/resource_group.html)

## TODO

This project is just getting started, so stay tuned!

## Authors

- **Sai Nimmagadda** - _Initial implementation_ - [snimmagadda1](https://github.com/snimmagadda1)

## Packages 
* [spf13/cobra](https://github.com/spf13/cobra#flags)
* [Azure/azure-sdk-for-go](https://github.com/Azure/azure-sdk-for-go)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
