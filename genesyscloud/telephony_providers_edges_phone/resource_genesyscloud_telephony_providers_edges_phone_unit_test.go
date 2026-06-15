package telephony_providers_edges_phone

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUnitGetLineProperties(t *testing.T) {
	tests := []struct {
		name         string
		resourceData *schema.ResourceData
		want         []linePropertyConfig
	}{
		{
			name:         "empty_resource_data",
			resourceData: schema.TestResourceDataRaw(t, linePropertiesTestSchema(), map[string]interface{}{}),
			want:         nil,
		},
		{
			name: "valid_line_properties",
			resourceData: schema.TestResourceDataRaw(t, linePropertiesTestSchema(), map[string]interface{}{
				"line_properties": []interface{}{
					map[string]interface{}{
						"line_id":      "guid-1",
						"line_address": "192.168.1.1",
					},
					map[string]interface{}{
						"line_id":        "guid-2",
						"remote_address": "10.0.0.1",
					},
				},
			}),
			want: []linePropertyConfig{
				{
					LineID:      "guid-1",
					LineAddress: "192.168.1.1",
				},
				{
					LineID:        "guid-2",
					RemoteAddress: "10.0.0.1",
				},
			},
		},
		{
			name: "empty_line_properties",
			resourceData: schema.TestResourceDataRaw(t, linePropertiesTestSchema(), map[string]interface{}{
				"line_properties": []interface{}{
					map[string]interface{}{},
				},
			}),
			want: []linePropertyConfig{
				{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLineProperties(tt.resourceData)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getLineProperties() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func linePropertiesTestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"line_properties": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 12,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"line_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"line_address": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"remote_address": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}
