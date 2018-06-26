provider "cloudmqtt" {}

resource "cloudmqtt_instance" "mqtt_koala" {
  name   = "terraform-provider-test"
  plan   = "koala"
  region = "amazon-web-services::us-east-1"
}

output "mqtt_url" {
  value = "${cloudmqtt_instance.mqtt_koala.url}"
}
