package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListRegions(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": [
				{
				  "key": "ca",
				  "label": "Canada"
				},
				{
				  "key": "eu",
				  "label": "Europe"
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/addressing/regional_hostnames/regions", handler)

	want := []Region{
		{
			Key:   "ca",
			Label: "Canada",
		},
		{
			Key:   "eu",
			Label: "Europe",
		},
	}

	actual, err := client.ListDataLocalizationRegions(context.Background(), AccountIdentifier(testAccountID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListRegionalHostnames(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": [
			{
			  "hostname": "ca.regional.ipam.rocks",
			  "region_key": "ca",
			  "created_on": "2023-01-14T00:47:57.060267Z"
			}
		  ],
		  "success": true,
		  "errors": [],
		  "messages": []
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/addressing/regional_hostnames", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2023-01-14T00:47:57.060267Z")
	want := []RegionalHostname{
		{
			Hostname:  "ca.regional.ipam.rocks",
			RegionKey: "ca",
			CreatedOn: &createdOn,
		},
	}

	actual, err := client.ListDataLocalizationRegionalHostnames(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateRegionalHostname(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": {
			  "hostname": "ca.regional.ipam.rocks",
			  "region_key": "ca",
			  "created_on": "2023-01-14T00:47:57.060267Z"
		  },
		  "success": true,
		  "errors": [],
		  "messages": []
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/addressing/regional_hostnames", handler)

	want := RegionalHostname{
		RegionKey: "ca",
		Hostname:  "ca.regional.ipam.rocks",
	}

	actual, err := client.CreateDataLocalizationRegionalHostname(context.Background(), ZoneIdentifier(testZoneID), want)
	createdOn, _ := time.Parse(time.RFC3339, "2023-01-14T00:47:57.060267Z")
	want.CreatedOn = &createdOn
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetRegionalHostname(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": {
			  "hostname": "ca.regional.ipam.rocks",
			  "region_key": "ca",
			  "created_on": "2023-01-14T00:47:57.060267Z"
		  },
		  "success": true,
		  "errors": [],
		  "messages": []
		}`)
	}

	testHostname := "ca.regional.ipam.rocks"
	mux.HandleFunc("/zones/"+testZoneID+"/addressing/regional_hostnames/"+testHostname, handler)

	actual, err := client.GetDataLocalizationRegionalHostname(context.Background(), ZoneIdentifier(testZoneID), testHostname)
	createdOn, _ := time.Parse(time.RFC3339, "2023-01-14T00:47:57.060267Z")
	want := RegionalHostname{
		RegionKey: "ca",
		Hostname:  "ca.regional.ipam.rocks",
		CreatedOn: &createdOn,
	}
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateRegionalHostname(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "result": {
			  "hostname": "ca.regional.ipam.rocks",
			  "region_key": "eu",
			  "created_on": "2023-01-14T00:47:57.060267Z"
		  },
		  "success": true,
		  "errors": [],
		  "messages": []
		}`)
	}

	testHostname := "ca.regional.ipam.rocks"
	want := RegionalHostname{
		RegionKey: "eu",
		Hostname:  testHostname,
	}
	mux.HandleFunc("/zones/"+testZoneID+"/addressing/regional_hostnames/"+testHostname, handler)

	actual, err := client.UpdateDataLocalizationRegionalHostname(context.Background(), ZoneIdentifier(testZoneID), want)
	createdOn, _ := time.Parse(time.RFC3339, "2023-01-14T00:47:57.060267Z")
	want.CreatedOn = &createdOn
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteRegionalHostname(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
	}

	testHostname := "ca.regional.ipam.rocks"
	mux.HandleFunc("/zones/"+testZoneID+"/addressing/regional_hostnames/"+testHostname, handler)

	err := client.DeleteDataLocalizationRegionalHostname(context.Background(), ZoneIdentifier(testZoneID), testHostname)
	assert.NoError(t, err)
}
