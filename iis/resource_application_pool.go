package iis

import (
	"context"
	"github.com/nrgribeiro/microsoft-iis-administration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const NameKey = "name"
const ManagedRuntimeKey = "runtime"
const StatusKey = "status"

func resourceApplicationPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationPoolCreate,
		ReadContext:   resourceApplicationPoolRead,
		UpdateContext: resourceApplicationPoolUpdate,
		DeleteContext: resourceApplicationPoolDelete,

		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			ManagedRuntimeKey: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			StatusKey: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "started",
			},
		},
	}
}

func resourceApplicationPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	name := d.Get(NameKey).(string)
	runtime := d.Get(ManagedRuntimeKey).(string)
	pool, err := client.CreateAppPool(ctx, name, runtime)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(pool.ID)
	return nil
}

func resourceApplicationPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	id := d.Id()
	appPool, err := client.ReadAppPool(ctx, id)
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	if err = d.Set(NameKey, appPool.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(ManagedRuntimeKey, appPool.ManagedRuntimeVersion); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(StatusKey, appPool.Status); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceApplicationPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	if d.HasChange(NameKey) || d.HasChange(ManagedRuntimeKey) {
		applicationPool, err := client.UpdateAppPool(ctx, d.Id(), d.Get(NameKey).(string), d.Get(ManagedRuntimeKey).(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(applicationPool.ID)
	}
	return nil
}

func resourceApplicationPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*iis.Client)
	id := d.Id()
	err := client.DeleteAppPool(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
