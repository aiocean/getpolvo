output "service_address" {
  value = google_dns_record_set.resource_recordset.name
}
