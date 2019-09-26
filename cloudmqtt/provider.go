package cloudmqtt

import (
	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDMQTT_APIKEY", nil),
				Description: "Key used to authentication to the CloudMQTT API",
			},
			"baseurl": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "https://customer.cloudmqtt.com",
				Optional:    true,
				Description: "Base URL to CloudMQTT website",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudmqtt_instance": resourceInstance(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return api.New(d.Get("baseurl").(string), d.Get("apikey").(string)), nil
}
