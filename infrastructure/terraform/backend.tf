terraform {
  backend "gcs" {
    bucket = "REPLACE_WITH_-backend-config"
    prefix = "REPLACE_WITH_-backend-config"
  }
}
