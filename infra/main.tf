resource "google_cloud_run_v2_service" "default" {
  name     = "cloudrun-service"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  project = "esc-api-384517"

  template {
    containers {
      image = "gcr.io/esc-api-384517/cloud-run:latest"
      env {
        name = "TOKEN"
        value = var.TOKEN
      }
    }
  }
}

variable "TOKEN" {
  type = string
  description = "API Token"
}
