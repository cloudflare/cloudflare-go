package cloudflare

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIShieldDiscoveryOperationID = "d3f614c0-7d73-4e4f-8d17-4215e7d78b77"

func TestListAPIShieldDiscoveryOperations(t *testing.T) {
	endpoint := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations", testZoneID)
	response := `{
		"success" : true,
		"errors": [],
		"messages": [],
		"result": [
			{
				"id": "9def2cb0-3ed0-4737-92ca-f09efa4718fd",
				"method": "POST",
				"host": "api.cloudflare.com",
				"endpoint": "/client/v4/zones",
				"last_updated": "2023-03-02T15:46:06.000000Z",
				"origin": ["ML"],
				"state": "review"
			}
		],
		"result_info": {
			"page": 3,
			"per_page": 20,
			"count": 1,
			"total_count": 2000
		}
	}`

	setup()
	t.Cleanup(teardown)
	handler := func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		require.Empty(t, r.URL.Query())
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, response)
	}

	mux.HandleFunc(endpoint, handler)

	actual, actualResultInfo, err := client.ListAPIShieldDiscoveryOperations(
		context.Background(),
		ZoneIdentifier(testZoneID),
		ListAPIShieldDiscoveryOperationsParams{},
	)

	lastUpdated := time.Date(2023, time.March, 2, 15, 46, 6, 0, time.UTC)
	expectedOps := []APIShieldDiscoveryOperation{
		{
			Method:      "POST",
			Host:        "api.cloudflare.com",
			Endpoint:    "/client/v4/zones",
			ID:          "9def2cb0-3ed0-4737-92ca-f09efa4718fd",
			LastUpdated: &lastUpdated,
			Origin:      []string{"ML"},
			State:       "review",
		},
	}

	expectedResultInfo := ResultInfo{
		Page:    3,
		PerPage: 20,
		Count:   1,
		Total:   2000,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, expectedOps, actual)
		assert.Equal(t, expectedResultInfo, actualResultInfo)
	}
}

func TestListAPIShieldDiscoveryOperationsWithParams(t *testing.T) {
	endpoint := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations", testZoneID)
	response := `{
		"success" : true,
		"errors": [],
		"messages": [],
		"result": [
			{
				"id": "9def2cb0-3ed0-4737-92ca-f09efa4718fd",
				"method": "POST",
				"host": "api.cloudflare.com",
				"endpoint": "/client/v4/zones",
				"last_updated": "2023-03-02T15:46:06.000000Z",
				"origin": ["ML"],
				"state": "review",
				"features": {
					"traffic_stats": {}
				}
			}
		],
		"result_info": {
			"page": 3,
			"per_page": 20,
			"count": 1,
			"total_count": 2000
		}
	}`

	tests := []struct {
		name           string
		params         ListAPIShieldDiscoveryOperationsParams
		expectedParams url.Values
	}{
		{
			name: "all params",
			params: ListAPIShieldDiscoveryOperationsParams{
				Direction: "desc",
				OrderBy:   "host",
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Hosts:    []string{"api.cloudflare.com", "developers.cloudflare.com"},
					Methods:  []string{"GET", "PUT"},
					Endpoint: "/client",
					Origin:   "ML",
					State:    "review",
					Diff:     true,
				},
				PaginationOptions: PaginationOptions{
					Page:    1,
					PerPage: 25,
				},
			},
			expectedParams: url.Values{
				"direction": []string{"desc"},
				"order":     []string{"host"},
				"host":      []string{"api.cloudflare.com", "developers.cloudflare.com"},
				"method":    []string{"GET", "PUT"},
				"endpoint":  []string{"/client"},
				"origin":    []string{"ML"},
				"state":     []string{"review"},
				"diff":      []string{"true"},
				"page":      []string{"1"},
				"per_page":  []string{"25"},
			},
		},
		{
			name: "direction only",
			params: ListAPIShieldDiscoveryOperationsParams{
				Direction: "desc",
			},
			expectedParams: url.Values{
				"direction": []string{"desc"},
			},
		},
		{
			name: "order only",
			params: ListAPIShieldDiscoveryOperationsParams{
				OrderBy: "host",
			},
			expectedParams: url.Values{
				"order": []string{"host"},
			},
		},
		{
			name: "hosts only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Hosts: []string{"api.cloudflare.com", "developers.cloudflare.com"},
				},
			},
			expectedParams: url.Values{
				"host": []string{"api.cloudflare.com", "developers.cloudflare.com"},
			},
		},
		{
			name: "methods only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Methods: []string{"GET", "PUT"},
				},
			},
			expectedParams: url.Values{
				"method": []string{"GET", "PUT"},
			},
		},
		{
			name: "endpoint only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Endpoint: "/client",
				},
			},
			expectedParams: url.Values{
				"endpoint": []string{"/client"},
			},
		},
		{
			name: "origin only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Origin: "ML",
				},
			},
			expectedParams: url.Values{
				"origin": []string{"ML"},
			},
		},
		{
			name: "state only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					State: "review",
				},
			},
			expectedParams: url.Values{
				"state": []string{"review"},
			},
		},
		{
			name: "diff only",
			params: ListAPIShieldDiscoveryOperationsParams{
				APIShieldListDiscoveryOperationsFilters: APIShieldListDiscoveryOperationsFilters{
					Diff: true,
				},
			},
			expectedParams: url.Values{
				"diff": []string{"true"},
			},
		},
		{
			name: "pagination only",
			params: ListAPIShieldDiscoveryOperationsParams{
				PaginationOptions: PaginationOptions{
					Page:    1,
					PerPage: 25,
				},
			},
			expectedParams: url.Values{
				"page":     []string{"1"},
				"per_page": []string{"25"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setup()
			t.Cleanup(teardown)
			handler := func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				require.Equal(t, test.expectedParams, r.URL.Query())
				w.Header().Set("content-type", "application/json")
				fmt.Fprint(w, response)
			}

			mux.HandleFunc(endpoint, handler)

			actual, _, err := client.ListAPIShieldDiscoveryOperations(
				context.Background(),
				ZoneIdentifier(testZoneID),
				test.params,
			)

			lastUpdated := time.Date(2023, time.March, 2, 15, 46, 6, 0, time.UTC)
			expected := []APIShieldDiscoveryOperation{
				{
					Method:      "POST",
					Host:        "api.cloudflare.com",
					Endpoint:    "/client/v4/zones",
					ID:          "9def2cb0-3ed0-4737-92ca-f09efa4718fd",
					LastUpdated: &lastUpdated,
					Origin:      []string{"ML"},
					State:       "review",
					Features: map[string]any{
						"traffic_stats": map[string]any{},
					},
				},
			}

			if assert.NoError(t, err) {
				assert.Equal(t, expected, actual)
			}
		})
	}
}

func TestPatchAPIShieldDiscoveryOperation(t *testing.T) {
	setup()
	t.Cleanup(teardown)

	endpoint := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations/%s", testZoneID, testAPIShieldDiscoveryOperationID)
	response := `{
		"success" : true,
		"errors": [],
		"messages": [],
		"result": {
			"state": "ignored"
		}
	}`

	handler := func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		require.Empty(t, r.URL.Query())

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, `{"state":"ignored"}`, string(body))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, response)
	}

	mux.HandleFunc(endpoint, handler)

	actual, err := client.PatchAPIShieldDiscoveryOperation(
		context.Background(),
		ZoneIdentifier(testZoneID),
		PatchAPIShieldDiscoveryOperationParams{
			OperationID: testAPIShieldDiscoveryOperationID,
			PatchAPIShieldDiscoveryOperation: PatchAPIShieldDiscoveryOperation{
				State: "ignored",
			},
		},
	)

	// patch result is a cut down representation of the schema
	// so metadata like created date is not populated
	expected := &PatchAPIShieldDiscoveryOperation{
		State: "ignored",
	}

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

	assert.NoError(t, err)
}

func TestPatchAPIShieldDiscoveryOperations(t *testing.T) {
	setup()
	t.Cleanup(teardown)

	endpoint := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations", testZoneID)
	response := `{
		"success" : true,
		"errors": [],
		"messages": [],
		"result": {
			"9b16ce22-d1bf-425d-869f-a11f8240fafb": { "state": "ignored" },
			"c51c2ea1-a690-48fd-8e3f-7fc79b269947": { "state": "review" }
		}
	}`

	handler := func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		require.Empty(t, r.URL.Query())

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, `{"9b16ce22-d1bf-425d-869f-a11f8240fafb":{"state":"ignored"},"c51c2ea1-a690-48fd-8e3f-7fc79b269947":{"state":"review"}}`, string(body))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, response)
	}

	mux.HandleFunc(endpoint, handler)

	actual, err := client.PatchAPIShieldDiscoveryOperations(
		context.Background(),
		ZoneIdentifier(testZoneID),
		PatchAPIShieldDiscoveryOperationsParams{
			"9b16ce22-d1bf-425d-869f-a11f8240fafb": PatchAPIShieldDiscoveryOperation{State: "ignored"},
			"c51c2ea1-a690-48fd-8e3f-7fc79b269947": PatchAPIShieldDiscoveryOperation{State: "review"},
		},
	)

	expected := &PatchAPIShieldDiscoveryOperationsParams{
		"9b16ce22-d1bf-425d-869f-a11f8240fafb": PatchAPIShieldDiscoveryOperation{State: "ignored"},
		"c51c2ea1-a690-48fd-8e3f-7fc79b269947": PatchAPIShieldDiscoveryOperation{State: "review"},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

	assert.NoError(t, err)
}

func TestMustProvideDiscoveryOperationID(t *testing.T) {
	setup()
	t.Cleanup(teardown)

	_, err := client.PatchAPIShieldDiscoveryOperation(context.Background(), ZoneIdentifier(testZoneID), PatchAPIShieldDiscoveryOperationParams{})
	require.ErrorContains(t, err, "params.OperationID must be provided")
}
