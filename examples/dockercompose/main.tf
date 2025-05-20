terraform {
  required_providers {
    tfcoolify = {
      source = "local/tfcoolify"
    }
  }
}

provider "tfcoolify" {
  api_url   = "http://your-coolify-instance:8000"
  api_token = "your-api-token"
}

resource "tfcoolify_dockercompose_app" "example" {
  name              = "example-compose-app"
  project_uuid      = "your-project-uuid"
  server_uuid       = "your-server-uuid"
  docker_compose_raw = <<-EOT
    version: '3'
    services:
      web:
        image: nginx:latest
        ports:
          - "80:80"
  EOT
} 