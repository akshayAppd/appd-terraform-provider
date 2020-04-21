package appd

import (
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func resourceJavaAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceJavaAgentCreate,
		Read:   resourceJavaAgentRead,
		Update: resourceJavaAgentUpdate,
		Delete: resourceJavaAgentDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jdk_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tier_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

/**
* Starts Java Agent with params specified in main.tf
**/
func resourceJavaAgentCreate(d *schema.ResourceData, m interface{}) error {

	path := d.Get("path").(string)
	jdkPath := d.Get("jdk_path").(string)
	pid := d.Get("pid").(string)

	startup_params := "java -Xbootclasspath/a:" + jdkPath + " -jar " + path + " " + pid
	ioutil.WriteFile("urlJA.txt", []byte(startup_params), 0644)
	cmd := exec.Command("bash", "-c", startup_params)
	err := cmd.Start()
	if err != nil {
		log.Error(err.Error())
	}

	u1 := uuid.Must(uuid.NewV4(), nil)
	d.SetId(u1.String())
	time.Sleep(60 * time.Second)
	return nil
}

/**
* Fetch all applications to get the appliction id of the one created above
**/
func resourceJavaAgentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceJavaAgentUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceJavaAgentDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
