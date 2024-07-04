package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListGatewayCategories(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"beta": false,
					"class": "premium",
					"description": "Sites related to educational content.",
					"id": 0,
					"name": "Education",
					"subcategories": [
						{
							"beta": true,
							"class": "premium",
							"description": "Sites related to educational content.",
							"id": 0,
							"name": "Education"
						}
					]
				}
			],
			"result_info": {
				"count": 1,
				"page": 1,
				"per_page": 20,
				"total_count": 2000
			}
		}`)
	}

	want := []GatewayCategory{
		{
			Beta:        BoolPtr(false),
			Class:       "premium",
			Description: "Sites related to educational content.",
			ID:          0,
			Name:        "Education",
			Subcategories: []GatewayCategory{
				{
					Beta:        BoolPtr(true),
					Class:       "premium",
					Description: "Sites related to educational content.",
					ID:          0,
					Name:        "Education",
				},
			},
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/categories", handler)

	actual, _, err := client.ListGatewayCategories(context.Background(), AccountIdentifier(testAccountID), ListGatewayCategoriesParams{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
