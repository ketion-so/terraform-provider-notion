package notion

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceNotionUser() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a Spinnaker application resource",
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"people": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: getUserPeopleSchema(),
				},
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bot": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: getUserBotSchema(),
				},
			},
		},
		ReadContext: resourceNotionUserRead,
		Importer: &schema.ResourceImporter{
			StateContext: resourceNotionUserImport,
		},
	}
}

func getUserPeopleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email": {
			Type:        schema.TypeString,
			Description: "Email of the people",
			Computed:    true,
		},
	}
}

func getUserBotSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
