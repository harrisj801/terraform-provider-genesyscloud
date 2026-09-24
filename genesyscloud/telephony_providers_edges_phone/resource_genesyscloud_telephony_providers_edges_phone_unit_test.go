package telephony_providers_edges_phone

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Tests the GetLineProperties function to ensure that the NIL values are checked*/
func TestUnitIsWebRtcPhoneAlreadyAssignedError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "matching api error",
			err:      fmt.Errorf("API Error: 400 - A web rtc phone has already been assigned to this user. (4692d3e3-0c6f-48ef-bdd9-0ec58991c5b5)"),
			expected: true,
		},
		{
			name:     "unrelated error",
			err:      fmt.Errorf("API Error: 400 - invalid request"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWebRtcPhoneAlreadyAssignedError(tt.err); got != tt.expected {
				t.Errorf("isWebRtcPhoneAlreadyAssignedError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

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
						"line_name":    "name-1",
						"line_address": "192.168.1.1",
					},
					map[string]interface{}{
						"line_name":      "name-1",
						"remote_address": "10.0.0.1",
					},
				},
			}),
			want: []linePropertyConfig{
				{
					LineName:    "name-1",
					LineAddress: "192.168.1.1",
				},
				{
					LineName:      "name-2",
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
					"line_name": {
						Type:     schema.TypeString,
						Optional: true,
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
