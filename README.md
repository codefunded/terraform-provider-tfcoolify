# Terraform Provider for Coolify

This Terraform provider allows you to manage Coolify applications using Terraform. It supports creating and managing both Dockerfile-based and Docker Compose-based applications.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

## Installation

### Using Terraform Registry

```hcl
terraform {
  required_providers {
    tfcoolify = {
      source = "codefunded/tfcoolify"
      version = "~> 1.0"
    }
  }
}
```

### Manual Installation

1. Download the latest release for your platform from the [releases page](https://github.com/codefunded/terraform-provider-tfcoolify/releases)
2. Extract the binary to your Terraform plugins directory:
   - Linux: `~/.terraform.d/plugins/registry.terraform.io/codefunded/tfcoolify/1.0.0/linux_amd64/`
   - macOS: `~/.terraform.d/plugins/registry.terraform.io/codefunded/tfcoolify/1.0.0/darwin_amd64/`
   - Windows: `%APPDATA%\terraform.d\plugins\registry.terraform.io\codefunded\tfcoolify\1.0.0\windows_amd64\`

## Provider Configuration

```hcl
provider "tfcoolify" {
  api_url   = "https://your-coolify-instance.com"
  api_token = "your-api-token"
}
```

### Configuration Reference

| Name | Type | Required | Description |
|------|------|----------|-------------|
| api_url | string | Yes | The URL of your Coolify instance |
| api_token | string | Yes | Your Coolify API token |

## Resources

### Dockerfile Application

Creates a Dockerfile-based application in Coolify.

```hcl
resource "tfcoolify_dockerfile_app" "example" {
  project_uuid = "your-project-uuid"
  server_uuid  = "your-server-uuid"
  name         = "my-app"
  dockerfile   = <<-EOT
    FROM node:18
    WORKDIR /app
    COPY . .
    RUN npm install
    CMD ["npm", "start"]
  EOT
  domains      = "example.com,www.example.com"
  ports_exposes = "3000"
  
  health_check_enabled = true
  health_check_path    = "/health"
  health_check_port    = "3000"
  health_check_host    = "0.0.0.0"
  health_check_method  = "GET"
  health_check_scheme  = "http"
  health_check_return_code = 200
  health_check_interval    = 10
}
```

#### Arguments Reference

| Name | Type | Required | Description |
|------|------|----------|-------------|
| project_uuid | string | Yes | UUID of the Coolify project |
| server_uuid | string | Yes | UUID of the Coolify server |
| name | string | Yes | Name of the application |
| dockerfile | string | Yes | Dockerfile content |
| domains | string | Yes | Comma-separated list of domains |
| ports_exposes | string | Yes | Comma-separated list of ports to expose |
| health_check_enabled | bool | No | Enable health checks (default: false) |
| health_check_path | string | No | Health check path (default: "/") |
| health_check_port | string | No | Health check port (default: "80") |
| health_check_host | string | No | Health check host (default: "0.0.0.0") |
| health_check_method | string | No | Health check HTTP method (default: "GET") |
| health_check_scheme | string | No | Health check scheme (default: "http") |
| health_check_return_code | int | No | Expected HTTP return code (default: 200) |
| health_check_interval | int | No | Health check interval in seconds (default: 10) |

#### Attributes Reference

| Name | Type | Description |
|------|------|-------------|
| uuid | string | The UUID of the created application |

### Docker Compose Application

Creates a Docker Compose-based application in Coolify.

```hcl
resource "tfcoolify_dockercompose_app" "example" {
  project_uuid = "your-project-uuid"
  server_uuid  = "your-server-uuid"
  name         = "my-compose-app"
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

#### Arguments Reference

| Name | Type | Required | Description |
|------|------|----------|-------------|
| project_uuid | string | Yes | UUID of the Coolify project |
| server_uuid | string | Yes | UUID of the Coolify server |
| name | string | Yes | Name of the application |
| docker_compose_raw | string | Yes | Docker Compose file content |

#### Attributes Reference

| Name | Type | Description |
|------|------|-------------|
| uuid | string | The UUID of the created application |

## Examples

### Basic Dockerfile Application

```hcl
provider "tfcoolify" {
  api_url   = "https://coolify.example.com"
  api_token = "your-api-token"
}

resource "tfcoolify_dockerfile_app" "web" {
  project_uuid = "proj-123"
  server_uuid  = "srv-456"
  name         = "web-app"
  dockerfile   = file("${path.module}/Dockerfile")
  domains      = "web.example.com"
  ports_exposes = "8080"
}
```

### Basic Docker Compose Application

```hcl
provider "tfcoolify" {
  api_url   = "https://coolify.example.com"
  api_token = "your-api-token"
}

resource "tfcoolify_dockercompose_app" "stack" {
  project_uuid = "proj-123"
  server_uuid  = "srv-456"
  name         = "full-stack"
  docker_compose_raw = file("${path.module}/docker-compose.yml")
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 