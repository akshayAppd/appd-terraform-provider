package appd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Optional: true,
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

	host := m.(*Controller).Host
	user := m.(*Controller).User
	password := m.(*Controller).Password
	port := m.(*Controller).Port
	account := m.(*Controller).Account
	application_id := d.Get("application_id").(string)
	auth := user + "@" + account
	url := "https://" + host + ":" + port + "/controller/alerting/rest/v1/applications/" + application_id + "/health-rules"

	data, err := ioutil.ReadFile("SERVERS_MATCHING_PATTERN.json")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.SetBasicAuth(auth, password)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Test case failed with error %s", err.Error())
		ioutil.WriteFile("error.txt", []byte(fmt.Sprint(err)), 777)
	}
	body, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("responseBody.txt", body, 777)

	val, err1 := jsonparser.GetUnsafeString(body, "id")
	ioutil.WriteFile("output.txt", []byte(val), 777)
	if err1 != nil {
		fmt.Printf("Test case failed with error %s", err1.Error())
	}

	d.SetId(val)

	defer resp.Body.Close()
	return resourceHealthRuleRead(d, m)
}

func resourceHealthRuleRead(d *schema.ResourceData, m interface{}) error {
	host := m.(*Controller).Host
	user := m.(*Controller).User
	password := m.(*Controller).Password
	port := m.(*Controller).Port
	account := m.(*Controller).Account
	application_id := d.Get("application_id").(string)
	auth := user + "@" + account
	url := "https://" + host + ":" + string(port) + "/controller/alerting/rest/v1/applications/" + application_id + "/health-rules/" + d.Id()

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(auth, password)

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Test case failed with error %s", err.Error())
		ioutil.WriteFile("readError.txt", []byte(fmt.Sprint(err)), 777)
	}
	fmt.Printf("HR read: " + string(body))
	d.Set("data", string(body))

	defer res.Body.Close()

	return nil
}

func resourceHealthRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceHealthRuleRead(d, m)
}

func resourceHealthRuleDelete(d *schema.ResourceData, m interface{}) error {
	host := m.(*Controller).Host
	user := m.(*Controller).User
	password := m.(*Controller).Password
	port := m.(*Controller).Port
	account := m.(*Controller).Account
	application_id := d.Get("application_id").(string)
	auth := user + "@" + account
	url := "https://" + host + ":" + string(port) + "/controller/alerting/rest/v1/applications/" + application_id + "/health-rules/" + d.Id()

	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(auth, password)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		ioutil.WriteFile("deleteError.txt", []byte(fmt.Sprint(err.Error())), 777)
	}

	defer res.Body.Close()
	return nil
}
