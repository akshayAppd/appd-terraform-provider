package appd

import (
	"github.com/buger/jsonparser"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
)

func resourceHealthRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthRuleCreate,
		Read:   resourceHealthRuleRead,
		Update: resourceHealthRuleUpdate,
		Delete: resourceHealthRuleDelete,

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

func resourceHealthRuleCreate(d *schema.ResourceData, m interface{}) error {

	application_id := d.Get("application_id").(string)
	json_file := d.Get("json_file").(string)

	parsed_response := CreateResource(m.(*Controller), application_id, "health-rules", json_file)
	val, err := jsonparser.GetUnsafeString([]byte(parsed_response), "id")
	if err != nil {
		log.Error("Error received on health rule create API" + err.Error())
	}
	d.SetId(val)

	return resourceHealthRuleRead(d, m)
}

func resourceHealthRuleRead(d *schema.ResourceData, m interface{}) error {

	application_id := d.Get("application_id").(string)
	d.Set("data", GetResource(m.(*Controller), application_id, "health-rules", d.Id()))

	return nil
}

func resourceHealthRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceHealthRuleRead(d, m)
}

func resourceHealthRuleDelete(d *schema.ResourceData, m interface{}) error {
	application_id := d.Get("application_id").(string)

	val := DeleteResource(m.(*Controller), application_id, "health-rules", d.Id())

	log.Info("HealthRule deleted: " + val)

	return nil
}
