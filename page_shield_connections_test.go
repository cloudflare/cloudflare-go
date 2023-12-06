package cloudflare

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock data for PageShieldConnections
var mockPageShieldConnections = []PageShieldConnection{
	{
		AddedAt:                 "2021-08-18T10:51:10.09615Z",
		DomainReportedMalicious: false,
		FirstPageURL:            "blog.cloudflare.com/page",
		FirstSeenAt:             "2021-08-18T10:51:08Z",
		Host:                    "blog.cloudflare.com",
		ID:                      "c9ef84a6bf5e47138c75d95e2f933e8f",
		LastSeenAt:              "2021-09-02T09:57:54Z",
		PageURLs:                []string{"blog.cloudflare.com/page1", "blog.cloudflare.com/page2"},
		URL:                     "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
		URLContainsCdnCgiPath:   false,
	},
	{
		AddedAt:                 "2021-09-18T10:51:10.09615Z",
		DomainReportedMalicious: false,
		FirstPageURL:            "blog.cloudflare.com/page02",
		FirstSeenAt:             "2021-08-18T10:51:08Z",
		Host:                    "blog.cloudflare.com",
		ID:                      "c9ef84a6bf5e47138c75d95e2f933e8f",
		LastSeenAt:              "2021-09-02T09:57:54Z",
		PageURLs:                []string{"blog.cloudflare.com/page1", "blog.cloudflare.com/page2"},
		URL:                     "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
		URLContainsCdnCgiPath:   false,
	},
	{
		AddedAt:                 "2021-10-18T10:51:10.09615Z",
		DomainReportedMalicious: false,
		FirstPageURL:            "blog.cloudflare.com/page03",
		FirstSeenAt:             "2021-08-18T10:51:08Z",
		Host:                    "blog.cloudflare.com",
		ID:                      "c9ef84a6bf5e47138c75d95e2f933e8f",
		LastSeenAt:              "2021-09-02T09:57:54Z",
		PageURLs:                []string{"blog.cloudflare.com/page1", "blog.cloudflare.com/page2"},
		URL:                     "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
		URLContainsCdnCgiPath:   false,
	},
}

func TestListPageShieldConnections(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/testzone/page_shield/connections", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		response := ListPageShieldConnectionsResponse{
			Result: mockPageShieldConnections,
		}
		json.NewEncoder(w).Encode(response)
	})

	result, _, err := client.ListPageShieldConnections(context.Background(), &ResourceContainer{Identifier: "testzone"}, ListPageShieldConnectionsParams{})
	assert.NoError(t, err)
	assert.Equal(t, mockPageShieldConnections, result)
}

func TestGetPageShieldConnection(t *testing.T) {
	setup()
	defer teardown()

	connectionID := "c9ef84a6bf5e47138c75d95e2f933e8f"
	mux.HandleFunc(fmt.Sprintf("/zones/testzone/page_shield/connections/%s", connectionID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		response := mockPageShieldConnections[0] // Assuming it's the first mock connection
		json.NewEncoder(w).Encode(response)
	})

	result, err := client.GetPageShieldConnection(context.Background(), &ResourceContainer{Identifier: "testzone"}, connectionID)
	assert.NoError(t, err)
	assert.Equal(t, &mockPageShieldConnections[0], result)
}
