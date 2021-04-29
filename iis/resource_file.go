package iis

import (
	"context"
	"github.com/nrgribeiro/microsoft-iis-administration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const fileNameKey = "name"
const filePhysicalPathKey = "physical_path"
const typeKey = "type"
const parentKey = "parent"

func resourceFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileCreate,
		ReadContext:   resourceFileRead,
		UpdateContext: resourceFileUpdate,
		DeleteContext: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			fileNameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			filePhysicalPathKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			typeKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			parentKey: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	request := createFileRequest(d)
	site, err := client.CreateFile(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(site.ID)
	return nil
}

func resourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	site, err := client.ReadFile(ctx, d.Id())
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}
	if err = d.Set(nameKey, site.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(physicalPathKey, site.PhysicalPath); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(typeKey, site.Type); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	id := d.Id()
	err := client.DeleteFile(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func createFileRequest(d *schema.ResourceData) iis.CreateFileRequest {
	name := d.Get(nameKey).(string)
	physicalPath := d.Get(physicalPathKey).(string)
	parentId := d.Get(parentKey).(string)
	typeName := d.Get(typeKey).(string)
	request := iis.CreateFileRequest{
		Name:         name,
		PhysicalPath: physicalPath,
		Parent: iis.FileParent{
			Id: parentId,
		},
		Type: typeName,
	}

	return request
}
