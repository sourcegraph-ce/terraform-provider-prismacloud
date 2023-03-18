package prismacloud

import (
	log "github.com/sourcegraph-ce/logrus"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/integration"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIntegrations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIntegrationsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("all integrations"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all integrations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"integration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration type",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status",
						},
						"valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Valid",
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	items, err := integration.List(client, "")
	if err != nil {
		return err
	}

	d.SetId("integrations")
	d.Set("total", len(items))

	listing := make([]interface{}, 0, len(items))
	for _, o := range items {
		listing = append(listing, map[string]interface{}{
			"integration_id":   o.Id,
			"name":             o.Name,
			"description":      o.Description,
			"integration_type": o.IntegrationType,
			"enabled":          o.Enabled,
			"status":           o.Status,
			"valid":            o.Valid,
		})
	}

	if err := d.Set("listing", listing); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
