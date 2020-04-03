package appd

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// To be used later when testing registerting MA with a new application
// var application_id string

// type Data struct {
// 	XMLName      xml.Name      `xml:"data" json:"-"`
// 	Applications []Application `xml:"person" json:"people"`
// }
// type Application struct {
// 	Id          string
// 	Name        string
// 	accountGuid string
// }

func resourceMachineAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceMachineAgentCreate,
		Read:   resourceMachineAgentRead,
		Update: resourceMachineAgentUpdate,
		Delete: resourceMachineAgentDelete,

		Schema: map[string]*schema.Schema{
			"path": {
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
				Required: true,
			},
			"sim_enabled": {
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

func resourceMachineAgentCreate(d *schema.ResourceData, m interface{}) error {

	host := m.(*Controller).Host
	port := m.(*Controller).Port
	account := m.(*Controller).Account

	path := d.Get("path").(string)
	accountAccessKey := d.Get("account_access_key").(string)
	simEnabled := d.Get("sim_enabled").(string)
	sslEnabled := d.Get("ssl_enabled").(string)

	startup_params := "java -Dappdynamics.controller.hostName=" + host + " -Dappdynamics.controller.port=" + port + " -Dappdynamics.agent.accountName=" + account + " -Dappdynamics.agent.accountAccessKey=" + accountAccessKey + " -Dappdynamics.sim.enabled=" + simEnabled + " -Dappdynamics.controller.ssl.enabled=" + sslEnabled
	ioutil.WriteFile("outputParams.txt", []byte(startup_params+" -jar "+path+"machineagent.jar "), 777)

	cmd := exec.Command("bash", "-c", startup_params+" -jar "+path+"machineagent.jar &")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	d.SetId(string(seededRand.Int()))
	time.Sleep(15 * time.Second)
	return nil
}

/**
* Fetch all applications to get the appliction id of the one created above
**/
func resourceMachineAgentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMachineAgentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMachineAgentRead(d, m)
}

func resourceMachineAgentDelete(d *schema.ResourceData, m interface{}) error {

	_, err := exec.Command("sh", "-c", " kill `ps | grep machineagent`").Output()

	if err != nil {
		ioutil.WriteFile("deleteError.txt", []byte(fmt.Sprint(err.Error())), 777)
	}
	return nil
}
