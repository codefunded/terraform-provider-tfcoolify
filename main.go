package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Coolify API URL",
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The Coolify API token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"tfcoolify_dockerfile_app": resourceDockerfileApp(),
			"tfcoolify_dockercompose_app": resourceDockerComposeApp(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		APIURL:   d.Get("api_url").(string),
		APIToken: d.Get("api_token").(string),
	}
	client, err := config.Client()
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return client, nil
} 