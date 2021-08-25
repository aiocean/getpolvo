variable "service_id" {
  type        = string
  default     = "getpolvo"
  description = "The id of the service. Should be a single word"
}

variable "service_base_domain" {
  type        = string
  default     = "truepro.fit"
  description = "Base do main for the service. It should be already exists"
}

variable "project_id" {
  type        = string
  default     = "trueprofit-frontend"
  description = "The project name"
}

variable "region" {
  type        = string
  default     = "asia-southeast1"
  description = "The region"
}

variable "env" {
  type        = string
  default     = "develop"
  description = "The environment"
}

variable "docker_image_url" {
  type        = string
  description = "The image digest"
}

variable "managed_dns_zone" {
  type        = string
  default     = "trueprofit"
  description = "The managed zone to create domain"
}

variable "polvo_service_address" {
  type        = string
  default     = "polvo.truepro.fit:443"
  description = "Polvo service address"
}

locals {
  service_domain    = "${var.service_id}.${var.service_base_domain}"
  service_full_name = "${var.service_id}-service"
}
