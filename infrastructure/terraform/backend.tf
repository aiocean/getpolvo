terraform {
  backend "gcs" {
    bucket = "REPLACE_WITH_-backend-config"
    prefix = "getpolvo/develop"
  }
}
