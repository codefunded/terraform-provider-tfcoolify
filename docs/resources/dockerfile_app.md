---
page_title: "Resource: tfcoolify_dockerfile_app"
description: |-
  Manages a Dockerfile-based application in Coolify.
---

# Resource: tfcoolify_dockerfile_app

This resource allows you to create and manage Dockerfile-based applications in Coolify.

## Example Usage

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

## Argument Reference

The following arguments are supported:

* `project_uuid` - (Required) UUID of the Coolify project.
* `server_uuid` - (Required) UUID of the Coolify server.
* `name` - (Required) Name of the application.
* `dockerfile` - (Required) Dockerfile content.
* `domains` - (Required) Comma-separated list of domains.
* `ports_exposes` - (Required) Comma-separated list of ports to expose.
* `health_check_enabled` - (Optional) Enable health checks. Defaults to `false`.
* `health_check_path` - (Optional) Health check path. Defaults to `"/"`.
* `health_check_port` - (Optional) Health check port. Defaults to `"80"`.
* `health_check_host` - (Optional) Health check host. Defaults to `"0.0.0.0"`.
* `health_check_method` - (Optional) Health check HTTP method. Defaults to `"GET"`.
* `health_check_scheme` - (Optional) Health check scheme. Defaults to `"http"`.
* `health_check_return_code` - (Optional) Expected HTTP return code. Defaults to `200`.
* `health_check_interval` - (Optional) Health check interval in seconds. Defaults to `10`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` - The UUID of the created application.

## Import

Dockerfile applications can be imported using their UUID:

```bash
terraform import tfcoolify_dockerfile_app.example <uuid>
``` 