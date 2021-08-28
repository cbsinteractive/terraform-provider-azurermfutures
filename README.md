# Terraform Provider for Microsoft Azure (Futures)

This Terraform provider aims to provide support for _some_ Microsoft Azure resources and configurations not yet fully supported by the official Terraform provider for Azure. Any use of this provider should be considered experimental, at worst, or temporary, at best. Please use the official provider if and when it meets your needs.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.0.3+

## Acceptance Tests

Running tests involves **TODO**.

Example:

```shell
make testacc
```

You can specify an individual test to run using the `TESTARGS` argument.

Example:

```shell
make testacc TEST=./azurermfutures TESTARGS='-run=TestAccStreamAnalyticsJob_basic'
```
