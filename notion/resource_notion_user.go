package notion

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ketion-so/go-notion/notion"
)

func resourceNotionUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*notion.Client)
	var diags diag.Diagnostics

	userID := d.Get("id").(string)

	user, err := client.Users.Get(ctx, userID)
	if err != nil {
		return diag.FromErr(err)
	}

	if user == nil {
		d.SetId("")
		return diags
	}

	if err := d.Set("name", user.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("avatar_url", user.AvatarURL); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("person", flatterUserPerson(user.Person)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("bot", flatterUserBot(user.Bot)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNotionUserImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if diags := resourceNotionUserRead(context.Background(), d, meta); diags.HasError() {
		return nil, fmt.Errorf("failed to read spinnaker application")
	}
	return []*schema.ResourceData{d}, nil
}

func flatterUserPerson(person *notion.People) []interface{} {
	p := map[string]interface{}{
		"email": person.Email,
	}

	return []interface{}{p}
}

func flatterUserBot(bot *notion.Bot) []interface{} {
	b := map[string]interface{}{}
	return []interface{}{b}
}
