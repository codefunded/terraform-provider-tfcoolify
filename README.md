# Terraform Provider for Coolify

This is a custom Terraform provider for managing Coolify applications.

## Building the Provider

1. Clone the repository
2. Enter the provider directory
3. Run `go mod tidy` to download dependencies
4. Run `go build -o terraform-provider-tfcoolify` to build the provider

## Using the Provider

1. Create a `.terraformrc` file in your home directory with the following content:

```hcl
provider_installation {
  dev_overrides {
    "yourusername/coolify" = "/path/to/terraform-provider-tfcoolify"
  }
  direct {}
}
```

2. Create a Terraform configuration file:

```hcl
terraform {
  required_providers {
    coolify = {
      source = "yourusername/coolify"
    }
  }
}

provider "coolify" {
  api_url   = "https://your-coolify-instance.com"
  api_token = "your-api-token"
}

resource "coolify_dockerfile_app" "example" {
  project_uuid = "your-project-uuid"
  server_uuid  = "your-server-uuid"
  name         = "example-app"
  dockerfile   = "FROM nginx:latest"
  domains      = "example.com"
  ports_exposes = "80"
  
  health_check_enabled = true
  health_check_path    = "/"
  health_check_port    = "80"
}

resource "coolify_dockercompose_app" "example" {
  project_uuid      = "your-project-uuid"
  server_uuid       = "your-server-uuid"
  name              = "compose-app"
  docker_compose_raw = <<-EOT
    version: '3'
    services:
      web:
        image: nginx:latest
        ports:
          - "80:80"
  EOT
}
```

## Features

- Create Dockerfile-based applications
- Create Docker Compose applications
- Configure health checks
- Manage application domains and ports
- Automatic cleanup on destroy

## Requirements

- Go 1.21 or later
- Terraform 1.0 or later 