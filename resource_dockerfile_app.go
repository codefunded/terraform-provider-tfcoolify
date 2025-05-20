package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DockerfileApp struct {
	ProjectUUID           string `json:"project_uuid"`
	ServerUUID           string `json:"server_uuid"`
	EnvironmentName      string `json:"environment_name"`
	Dockerfile           string `json:"dockerfile"`
	Name                 string `json:"name"`
	Domains              string `json:"domains"`
	PortsExposes         string `json:"ports_exposes"`
	HealthCheckEnabled   bool   `json:"health_check_enabled"`
	HealthCheckPath      string `json:"health_check_path"`
	HealthCheckPort      string `json:"health_check_port"`
	HealthCheckHost      string `json:"health_check_host"`
	HealthCheckMethod    string `json:"health_check_method"`
	HealthCheckScheme    string `json:"health_check_scheme"`
	HealthCheckReturnCode int    `json:"health_check_return_code"`
	HealthCheckInterval  int    `json:"health_check_interval"`
	InstantDeploy        bool   `json:"instant_deploy"`
}

func resourceDockerfileApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceDockerfileAppCreate,
		Read:   resourceDockerfileAppRead,
		Delete: resourceDockerfileAppDelete,

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
			"dockerfile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domains": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ports_exposes": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"health_check_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"health_check_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
				ForceNew: true,
			},
			"health_check_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "80",
				ForceNew: true,
			},
			"health_check_host": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
				ForceNew: true,
			},
			"health_check_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET",
				ForceNew: true,
			},
			"health_check_scheme": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http",
				ForceNew: true,
			},
			"health_check_return_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  200,
				ForceNew: true,
			},
			"health_check_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
				ForceNew: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDockerfileAppCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	app := DockerfileApp{
		ProjectUUID:           d.Get("project_uuid").(string),
		ServerUUID:           d.Get("server_uuid").(string),
		EnvironmentName:      "production",
		Name:                 d.Get("name").(string),
		Dockerfile:           d.Get("dockerfile").(string),
		Domains:              d.Get("domains").(string),
		PortsExposes:         d.Get("ports_exposes").(string),
		HealthCheckEnabled:   d.Get("health_check_enabled").(bool),
		HealthCheckPath:      d.Get("health_check_path").(string),
		HealthCheckPort:      d.Get("health_check_port").(string),
		HealthCheckHost:      d.Get("health_check_host").(string),
		HealthCheckMethod:    d.Get("health_check_method").(string),
		HealthCheckScheme:    d.Get("health_check_scheme").(string),
		HealthCheckReturnCode: d.Get("health_check_return_code").(int),
		HealthCheckInterval:  d.Get("health_check_interval").(int),
		InstantDeploy:        true,
	}

	jsonData, err := json.Marshal(app)
	if err != nil {
		return fmt.Errorf("error marshaling app data: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/applications/dockerfile", client.apiURL), bytes.NewBuffer(jsonData))
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

func resourceDockerfileAppRead(d *schema.ResourceData, m interface{}) error {
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

func resourceDockerfileAppDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/applications/%s", client.apiURL, d.Id()), nil)
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