---
page_title: "Resource: tfcoolify_dockercompose_app"
description: |-
  Manages a Docker Compose-based application in Coolify.
---

# Resource: tfcoolify_dockercompose_app

This resource allows you to create and manage Docker Compose-based applications in Coolify.

## Example Usage

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

## Argument Reference

The following arguments are supported:

* `project_uuid` - (Required) UUID of the Coolify project.
* `server_uuid` - (Required) UUID of the Coolify server.
* `name` - (Required) Name of the application.
* `docker_compose_raw` - (Required) Docker Compose file content.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` - The UUID of the created application.

## Import

Docker Compose applications can be imported using their UUID:

```bash
terraform import tfcoolify_dockercompose_app.example <uuid>
``` 