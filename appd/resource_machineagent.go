package appd

import (
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
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
			"unique_host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

/**
* Starts Machine Agent with params specified in main.tf
**/
func resourceMachineAgentCreate(d *schema.ResourceData, m interface{}) error {

	host := m.(*Controller).Host
	port := m.(*Controller).Port
	account := m.(*Controller).Account

	path := d.Get("path").(string)
	accountAccessKey := d.Get("account_access_key").(string)
	simEnabled := d.Get("sim_enabled").(string)
	sslEnabled := d.Get("ssl_enabled").(string)
	uniqueHostId := d.Get("unique_host_id").(string)
	application_name := d.Get("application_name").(string)
	node_name := d.Get("node_name").(string)
	tier_name := d.Get("tier_name").(string)

	startup_params := ""

	if simEnabled != "" {
		startup_params = "java -Dappdynamics.agent.uniqueHostId=" + uniqueHostId + " -Dappdynamics.controller.hostName=" + host + " -Dappdynamics.controller.port=" + port + " -Dappdynamics.agent.accountName=" + account + " -Dappdynamics.agent.accountAccessKey=" + accountAccessKey + " -Dappdynamics.sim.enabled=" + simEnabled + " -Dappdynamics.controller.ssl.enabled=" + sslEnabled
	} else {
		startup_params = "java -Dappdynamics.agent.uniqueHostId=" + uniqueHostId + " -Dappdynamics.controller.hostName=" + host + " -Dappdynamics.controller.port=" + port + " -Dappdynamics.agent.accountName=" + account + " -Dappdynamics.agent.accountAccessKey=" + accountAccessKey + " -Dappdynamics.agent.applicationName" + application_name + " -Dappdynamics.agent.nodeName" + node_name + " -Dappdynamics.agent.tierName" + tier_name + " -Dappdynamics.controller.ssl.enabled=" + sslEnabled
	}
	log.Info("Machine Agent startup command: " + startup_params + " -jar " + path + "machineagent.jar ")

	cmd := exec.Command("bash", "-c", startup_params+" -jar "+path+"machineagent.jar &")
	err := cmd.Start()
	if err != nil {
		log.Error(err.Error())
	}

	u1 := uuid.Must(uuid.NewV4(), nil)
	d.SetId(u1.String())
	time.Sleep(15 * time.Second)
	return nil
}

/**
* Fetch all applications to get the appliction id of the one created above
**/
func resourceMachineAgentRead(d *schema.ResourceData, m interface{}) error {
	// host := m.(*Controller).Host
	// //host := controller.Host
	// user := m.(*Controller).User
	// password := m.(*Controller).Password
	// port := m.(*Controller).Port
	// account := m.(*Controller).Account
	// auth := user + "@" + account
	// application_name := d.Get("application_name").(string)

	// ioutil.WriteFile("urlMA.txt", []byte("https://"+host+":"+string(port)+"/controller/rest/applications/"), 0644)

	// url := "https://" + host + ":" + string(port) + "/controller/rest/applications/"

	// req, err := http.NewRequest("GET", url, nil)
	// req.SetBasicAuth(auth, password)

	// res, err := http.DefaultClient.Do(req)
	// body, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Printf("Test case failed with error %s", err.Error())
	// 	ioutil.WriteFile("readErrorMA.txt", []byte(fmt.Sprint(err)), 777)
	// }

	// defer res.Body.Close()

	// var data Data
	// xml.Unmarshal(body, &data)
	// for i := 0; i < len(data.Applications); i++ {
	// 	if data.Applications[i].Name == application_name {
	// 		val := data.Applications[i].Id
	// 		application_id = val
	// 		ioutil.WriteFile("readOutputMA.txt", []byte(val), 777)
	// 	}
	// }
	return nil
}

func resourceMachineAgentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMachineAgentRead(d, m)
}

func resourceMachineAgentDelete(d *schema.ResourceData, m interface{}) error {

	_, err := exec.Command("sh", "-c", " kill `ps | grep machineagent`").Output()

	if err != nil {
		log.Error("Error stopping machine agent" + err.Error())
	}
	return nil
}
