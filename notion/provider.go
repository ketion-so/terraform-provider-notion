package notion

import (
	"context"

	"github.com/ketion-so/go-notion/notion"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Description: "Access token for the Notion API",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NOTION_API_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"notion_user": datasourceNotionUser(),
		},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(_ context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	accessToken := data.Get("access_token").(string)
	c := notion.NewClient(accessToken)
	return c, diags
}
