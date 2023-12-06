package cloudflare

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock PageShieldSettings data
var mockPageShieldSettings = PageShieldSettings{
	PageShield: PageShield{
		Enabled:                        true,
		UseCloudflareReportingEndpoint: true,
		UseConnectionURLPath:           true,
	},
	UpdatedAt: "2022-10-12T17:56:52.083582+01:00",
}

func TestGetPageShieldSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/testzone/page_shield", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		response := PageShieldSettingsResponse{
			PageShield: mockPageShieldSettings,
		}
		json.NewEncoder(w).Encode(response)
	})

	result, err := client.GetPageShieldSettings(context.Background(), &ResourceContainer{Identifier: "testzone"})
	assert.NoError(t, err)
	assert.Equal(t, &PageShieldSettingsResponse{PageShield: mockPageShieldSettings}, result)
}

func TestUpdatePageShieldSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/testzone/page_shield", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		var params PageShield
		err := json.NewDecoder(r.Body).Decode(&params)
		assert.NoError(t, err)

		response := PageShieldSettingsResponse{
			PageShield: PageShieldSettings{
				PageShield: params,
				UpdatedAt:  "2022-10-13T10:00:00.000Z",
			},
		}
		json.NewEncoder(w).Encode(response)
	})

	newSettings := PageShield{
		Enabled:                        false,
		UseCloudflareReportingEndpoint: false,
		UseConnectionURLPath:           false,
	}
	result, err := client.UpdatePageShieldSettings(context.Background(), &ResourceContainer{Identifier: "testzone"}, newSettings)
	assert.NoError(t, err)
	assert.Equal(t, false, result.PageShield.Enabled)
	assert.Equal(t, false, result.PageShield.UseCloudflareReportingEndpoint)
	assert.Equal(t, false, result.PageShield.UseConnectionURLPath)
}
