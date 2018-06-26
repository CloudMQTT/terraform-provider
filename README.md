# Terraform provider for CloudMQTT

Setup your CloudMQTT cluster with Terraform

## Install

```sh
git clone https://github.com/cloudmqtt/terraform-provider.git
cd terraform-provider
make depupdate
make init
```

Now the provider is installed in the terraform plugins folder and ready to be used.

## Example

```hcl
provider "cloudmqtt" {}

resource "cloudmqtt_instance" "mqtt_koala" {
  name   = "terraform-provider-test"
  plan   = "koala"
  region = "amazon-web-services::us-east-1"
}

output "mqtt_url" {
  value = "${cloudmqtt_instance.mqtt_koala.url}"
}
```



