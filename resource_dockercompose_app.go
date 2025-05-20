package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DockerComposeApp struct {
	ProjectUUID      string `json:"project_uuid"`
	ServerUUID       string `json:"server_uuid"`
	EnvironmentName  string `json:"environment_name"`
	DockerComposeRaw string `json:"docker_compose_raw"`
	Name             string `json:"name"`
	InstantDeploy    bool   `json:"instant_deploy"`
}

func resourceDockerComposeApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceDockerComposeAppCreate,
		Read:   resourceDockerComposeAppRead,
		Delete: resourceDockerComposeAppDelete,

		Schema: map[string]*schema.Schema{
			"project_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"server_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"docker_compose_raw": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDockerComposeAppCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	app := DockerComposeApp{
		ProjectUUID:      d.Get("project_uuid").(string),
		ServerUUID:       d.Get("server_uuid").(string),
		EnvironmentName:  "production",
		Name:             d.Get("name").(string),
		DockerComposeRaw: d.Get("docker_compose_raw").(string),
		InstantDeploy:    true,
	}

	jsonData, err := json.Marshal(app)
	if err != nil {
		return fmt.Errorf("error marshaling app data: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/applications/dockercompose", client.apiURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	uuid, ok := result["uuid"].(string)
	if !ok {
		return fmt.Errorf("uuid not found in response")
	}

	d.SetId(uuid)
	d.Set("uuid", uuid)

	return nil
}

func resourceDockerComposeAppRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/applications/%s", client.apiURL, d.Id()), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiToken))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func resourceDockerComposeAppDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/services/%s", client.apiURL, d.Id()), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiToken))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
} 