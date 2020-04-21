package appd

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Generates a http Request for GET/DELETE
func generateRequest(c *Controller, application_id string, end_point string, resource_id string, method_type string) *http.Request {
	protocol := c.Protocol
	host := c.Host
	user := c.User
	password := c.Password
	port := c.Port
	account := c.Account
	auth := user + "@" + account
	url := protocol + "://" + host + ":" + string(port) + "/controller/alerting/rest/v1/applications/" + application_id + "/" + end_point + "/" + resource_id
	log.Debug("URL to read resource: " + url)

	req, err := http.NewRequest(method_type, url, nil)
	req.SetBasicAuth(auth, password)
	if err != nil {
		log.Error("Failed to create request with error: " + err.Error())
	}
	return req
}

func GetResource(c *Controller, application_id string, end_point string, resource_id string) string {

	req := generateRequest(c, application_id, end_point, resource_id, "GET")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Failed to get response from API with error: " + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error("Failed to parse API response with error: " + err.Error())
	}
	log.Info("Resource found with data" + string(body))
	return string(body)
}

func CreateResource(c *Controller, application_id string, end_point string, file_path string) string {
	protocol := c.Protocol
	host := c.Host
	user := c.User
	password := c.Password
	port := c.Port
	account := c.Account
	auth := user + "@" + account
	url := protocol + "://" + host + ":" + string(port) + "/controller/alerting/rest/v1/applications/" + application_id + "/" + end_point

	log.Debug("URL generated to create resource " + url)

	data, err := ioutil.ReadFile(file_path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.SetBasicAuth(auth, password)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error("Resource creation failed: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Info("Response received on create API" + string(body))

	if err != nil {
		log.Error("Create API response read failed: " + err.Error())
	}

	return string(body)
}

func DeleteResource(c *Controller, application_id string, end_point string, resource_id string) string {
	req := generateRequest(c, application_id, end_point, resource_id, "DELETE")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error("Error received from delete API" + err.Error())
	}
	defer res.Body.Close()

	return resource_id
}
