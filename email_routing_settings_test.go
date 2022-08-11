package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func createTestEmailRoutingSettings() EmailRoutingSettings {
	created, _ := time.Parse(time.RFC3339, "2014-01-02T02:20:00Z")
	modified, _ := time.Parse(time.RFC3339, "2014-01-02T02:20:00Z")
	return EmailRoutingSettings{
		Tag:        "75610dab9e69410a82cf7e400a09ecec",
		Name:       "example.net",
		Enabled:    true,
		Created:    &created,
		Modified:   &modified,
		SkipWizard: true,
		Status:     "read",
	}
}

func TestEmailRouting_GetSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "75610dab9e69410a82cf7e400a09ecec",
    "name": "example.net",
    "enabled": true,
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z",
    "skip_wizard": true,
    "status": "read"
  }
}`)
	})

	_, err := client.EmailRoutingGetSettings(context.Background(), AccountIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := createTestEmailRoutingSettings()

	res, err := client.EmailRoutingGetSettings(context.Background(), AccountIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_Enable(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/enabled", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "75610dab9e69410a82cf7e400a09ecec",
    "name": "example.net",
    "enabled": true,
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z",
    "skip_wizard": true,
    "status": "read"
  }
}`)
	})

	_, err := client.EmailRoutingEnableRouting(context.Background(), AccountIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := createTestEmailRoutingSettings()

	res, err := client.EmailRoutingEnableRouting(context.Background(), AccountIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_Disabled(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/disable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "75610dab9e69410a82cf7e400a09ecec",
    "name": "example.net",
    "enabled": true,
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z",
    "skip_wizard": true,
    "status": "read"
  }
}`)
	})

	_, err := client.EmailRoutingDisableRouting(context.Background(), AccountIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := createTestEmailRoutingSettings()

	res, err := client.EmailRoutingDisableRouting(context.Background(), AccountIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_DNSSettings(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/dns", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "type": "A",
      "name": "example.com",
      "content": "127.0.0.1",
      "ttl": 3600,
      "priority": 10
    }
  ]
}`)
	})

	_, err := client.EmailRoutingDNSSettings(context.Background(), AccountIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := []DNSRecord{
		{
			Type:     "A",
			Name:     "example.com",
			Content:  "127.0.0.1",
			TTL:      3600,
			Priority: Uint16Ptr(10),
		},
	}

	res, err := client.EmailRoutingDNSSettings(context.Background(), AccountIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Len(t, res, 1)
		assert.Equal(t, want, res)
	}
}
