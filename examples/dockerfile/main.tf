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

resource "tfcoolify_dockerfile_app" "example" {
  name = "test-app"
  project_uuid      = "your-project-uuid"
  server_uuid       = "your-server-uuid"
  dockerfile = "FROM nginx:latest"
  domains = "example.com"
  ports_exposes = ["80"]
}