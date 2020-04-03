package appd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"appd_healthrule":   resourceHealthRule(),
			"appd_machineagent": resourceMachineAgent(),
		},
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return ControllerInfo(d.Get("protocol").(string), d.Get("host").(string), d.Get("port").(string), d.Get("user").(string), d.Get("password").(string), d.Get("account").(string))
	}
	return provider
}

// Controller information
type Controller struct {
	Host     string
	Port     string
	User     string
	Password string
	Account  string
	Protocol string
}

// Function to return a new instance of controller. To be used further for setting up clients and other configurations
func ControllerInfo(protocol string, controllerHost string, port string, username string, password string, account string) (*Controller, error) {

	c := &Controller{
		Protocol: protocol,
		Host:     controllerHost,
		Port:     port,
		User:     username,
		Password: password,
		Account:  account,
	}

	return c, nil
}
