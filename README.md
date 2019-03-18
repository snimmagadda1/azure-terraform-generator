# Azure Terraform Generator
[![Go Report Card](https://goreportcard.com/badge/github.com/snimmagadda1/azure-terraform-generator)](https://goreportcard.com/report/github.com/snimmagadda1/azure-terraform-generator)


Azure API -> Terraform Resources

## Overview
----
This is a work in progress initiative to create a CLI tool to generate Terraform definitions for Azure resources. The idea for this came about when teams I've worked with had existing Azure resources but were lacking terraform definitions. As a result, this slowed down their time to production.

## Usage

As a prerequisite to using the azure apis, credentials must be provided. Currently a file-based credential method is supported. As this project is built out, more authorization methods will be supported.

The simplest way to obtain this file is throughthe creation of an Azure service principal. Either download the credentials of an existing service principal to a file my.auth in the root directory or create a new service principal as follows: 

```
az ad sp create-for-rbac --sdk-auth > my.auth
```

## TODO
This project is just getting started, so stay tuned for the intitial implementation!

## Authors

* **Sai Nimmagadda** - *Initial work* - [snimmagadda1](https://github.com/snimmagadda1)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
