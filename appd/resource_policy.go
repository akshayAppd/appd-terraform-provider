package appd

import (
	"github.com/buger/jsonparser"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"json_file": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {

	application_id := d.Get("application_id").(string)
	json_file := d.Get("json_file").(string)

	parsed_response := CreateResource(m.(*Controller), application_id, "policies", json_file)
	val, err := jsonparser.GetUnsafeString([]byte(parsed_response), "id")

	if err != nil {
		log.Error("Failed to parse create policy response with error: " + err.Error())
	}
	d.SetId(val)

	return resourcePolicyRead(d, m)
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {

	application_id := d.Get("application_id").(string)
	d.Set("data", GetResource(m.(*Controller), application_id, "policies", d.Id()))

	return nil
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePolicyRead(d, m)
}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	application_id := d.Get("application_id").(string)
	val := DeleteResource(m.(*Controller), application_id, "policies", d.Id())

	log.Info("Policy deleted: " + val)

	return nil
}
