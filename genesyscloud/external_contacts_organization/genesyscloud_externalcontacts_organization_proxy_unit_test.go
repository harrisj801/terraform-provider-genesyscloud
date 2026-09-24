package external_contacts_organization

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mypurecloud/platform-client-sdk-go/v195/platformclientv2"
	"github.com/mypurecloud/terraform-provider-genesyscloud/genesyscloud/resource_cache"
)

func TestUnitGetAllExternalContactsOrganizationsUsesDivisionViews(t *testing.T) {
	var cursors []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/externalcontacts/scan/organizations/divisionviews/all" {
			t.Errorf("unexpected scan endpoint: %s", r.URL.Path)
			http.Error(w, "wrong endpoint", http.StatusNotFound)
			return
		}
		if r.URL.Query().Get("limit") != "200" {
			t.Errorf("unexpected limit: %s", r.URL.Query().Get("limit"))
		}
		cursor := r.URL.Query().Get("cursor")
		cursors = append(cursors, cursor)
		w.Header().Set("Content-Type", "application/json")
		switch cursor {
		case "":
			fmt.Fprint(w, `{"entities":[{"id":"org-1","name":"First"}],"cursors":{"after":"next"}}`)
		case "next":
			fmt.Fprint(w, `{"entities":[{"id":"org-2","name":"Second"}],"cursors":{}}`)
		default:
			http.Error(w, "unexpected cursor", http.StatusBadRequest)
		}
	}))
	defer server.Close()

	config := platformclientv2.NewConfigurationWithConfigFile("", false)
	config.BasePath = server.URL
	proxy := newExternalContactsOrganizationProxy(config)
	proxy.externalOrganizationCache = resource_cache.NewResourceCache[platformclientv2.Externalorganization]()

	organizations, _, err := getAllExternalContactsOrganizationFn(context.Background(), proxy)
	if err != nil {
		t.Fatal(err)
	}
	if organizations == nil || len(*organizations) != 2 {
		t.Fatalf("expected organizations from both pages, got %v", organizations)
	}
	if len(cursors) != 2 || cursors[0] != "" || cursors[1] != "next" {
		t.Fatalf("unexpected scan cursors: %v", cursors)
	}
}
