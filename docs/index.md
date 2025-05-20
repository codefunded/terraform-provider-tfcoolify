---
page_title: "Provider: Coolify"
description: |-
  The Coolify provider is used to interact with Coolify resources.
---

# Coolify Provider

The Coolify provider is used to interact with Coolify resources. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
terraform {
  required_providers {
    tfcoolify = {
      source = "codefunded/tfcoolify"
      version = "~> 1.0"
    }
  }
}

provider "tfcoolify" {
  api_url   = "https://your-coolify-instance.com"
  api_token = "your-api-token"
}
```

## Authentication

The Coolify provider offers a flexible means of providing credentials for authentication. The following methods are supported:

### API Token

You can provide your Coolify API token directly in the provider configuration:

```hcl
provider "tfcoolify" {
  api_url   = "https://your-coolify-instance.com"
  api_token = "your-api-token"
}
```

## Configuration Reference

The following arguments are supported:

* `api_url` - (Required) The URL of your Coolify instance. This can also be specified using the `TFC_COOLIFY_API_URL` environment variable.
* `api_token` - (Required) Your Coolify API token. This can also be specified using the `TFC_COOLIFY_API_TOKEN` environment variable.

## Environment Variables

The following environment variables can be used to configure the provider:

* `TFC_COOLIFY_API_URL` - The URL of your Coolify instance
* `TFC_COOLIFY_API_TOKEN` - Your Coolify API token 