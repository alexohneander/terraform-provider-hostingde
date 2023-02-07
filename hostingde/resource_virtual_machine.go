package hostingde

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineCreate,
		Read:   resourceVirtualMachineRead,
		Update: resourceVirtualMachineUpdate,
		Delete: resourceVirtualMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"productCode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

// TODO: Funktion muss fertiggestellt werden
func resourceVirtualMachineCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	productCode := d.Get("type").(string)

	if productCode == "" {
		productCode = "machine-virtualmachine-small-v1-1m"
	}
	req := VirtualMachineCreateRequest{
		BaseRequest: &BaseRequest{},
		Name:        name,
		ProductCode: productCode,
		Description: description,
	}
	resp, err := c.createVirtualMachine(req)
	if err != nil {
		return err
	}

	// d.SetId(resp.Response.ZoneConfig.ID)
	// return resourceZoneRead(d, m)
}

func resourceVirtualMachineRead(d *schema.ResourceData, m interface{}) error {

}

func resourceVirtualMachineUpdate(d *schema.ResourceData, m interface{}) error {

}

func resourceVirtualMachineDelete(d *schema.ResourceData, m interface{}) error {

}
