package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetSecondaryDNSZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "269d8f4853475ca241c4e730be286b20",
				"name": "www.example.com.",
				"primaries": [
					"23ff594956f20c2a721606e94745a8aa",
					"00920f38ce07c2e2f4df50b1f61d4194"
				],
				"auto_refresh_seconds": 86400,
				"soa_serial": 2019102400,
				"created_time": "2019-10-24T17:09:42.883908+01:00",
				"checked_time": "2019-10-24T17:09:42.883908+01:00",
				"modified_time": "2019-10-24T17:09:42.883908+01:00"
			}
		}
		`)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/secondary_dns", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	modifiedOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	checkedOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	want := SecondaryDNSZone{
		ID:   "269d8f4853475ca241c4e730be286b20",
		Name: "www.example.com.",
		Primaries: []string{
			"23ff594956f20c2a721606e94745a8aa",
			"00920f38ce07c2e2f4df50b1f61d4194",
		},
		AutoRefreshSeconds: 86400,
		SoaSerial:          2019102400,
		CreatedTime:        createdOn,
		CheckedTime:        checkedOn,
		ModifiedTime:       modifiedOn,
	}

	actual, err := client.GetSecondaryDNSZone(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateSecondaryDNSZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "269d8f4853475ca241c4e730be286b20",
				"name": "www.example.com.",
				"primaries": [
					"23ff594956f20c2a721606e94745a8aa",
					"00920f38ce07c2e2f4df50b1f61d4194"
				],
				"auto_refresh_seconds": 86400,
				"soa_serial": 2019102400,
				"created_time": "2019-10-24T17:09:42.883908+01:00",
				"checked_time": "2019-10-24T17:09:42.883908+01:00",
				"modified_time": "2019-10-24T17:09:42.883908+01:00"
			}
		}
		`)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/secondary_dns", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	modifiedOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	checkedOn, _ := time.Parse(time.RFC3339, "2019-10-24T17:09:42.883908+01:00")
	want := SecondaryDNSZone{
		ID:   "269d8f4853475ca241c4e730be286b20",
		Name: "www.example.com.",
		Primaries: []string{
			"23ff594956f20c2a721606e94745a8aa",
			"00920f38ce07c2e2f4df50b1f61d4194",
		},
		AutoRefreshSeconds: 86400,
		SoaSerial:          2019102400,
		CreatedTime:        createdOn,
		CheckedTime:        checkedOn,
		ModifiedTime:       modifiedOn,
	}

	newSecondaryZone := SecondaryDNSZone{
		Name: "www.example.com.",
		Primaries: []string{
			"23ff594956f20c2a721606e94745a8aa",
			"00920f38ce07c2e2f4df50b1f61d4194",
		},
		AutoRefreshSeconds: 86400,
	}

	actual, err := client.CreateSecondaryDNSZone(context.Background(), "01a7362d577a6c3019a474fd6f485823", newSecondaryZone)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

