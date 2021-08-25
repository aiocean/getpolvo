terraform {
  backend "gcs" {
    bucket = "getpolvo-tfstate"
    prefix = "env/develop"
  }
}
