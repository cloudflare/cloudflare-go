package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloudConnectorRules(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
		{
			"id": "some_id_1",
			"provider": "aws_s3",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"parameters": {
				"host": "aws.s3.bucket"
			}
	  	},
		{
			"id": "some_id_2",
			"provider": "r2",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"parameters": {
				"host": "r2.hostname"
			}
	  	}
	  ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}
	mux.HandleFunc("/zones/"+testZoneID+"/cloud_connector/rules", handler)

	want := []CloudConnectorRule{
		{
			ID:          "some_id_1",
			Provider:    "aws_s3",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "aws.s3.bucket",
			},
		},
		{
			ID:          "some_id_2",
			Provider:    "r2",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "r2.hostname",
			},
		},
	}

	zoneActual, err := client.ListZoneCloudConnectorRules(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}

func TestUpdateCloudConnectorRules(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
      "result": [
		{
			"id": "some_id_1",
			"provider": "aws_s3",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"parameters": {
				"host": "aws.s3.bucket"
			}
	  	},
		{
			"id": "some_id_2",
			"provider": "r2",
			"expression": "true",
			"enabled": true,
			"description": "some description",
			"parameters": {
				"host": "r2.hostname"
			}
	  	}
	  ],
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/cloud_connector/rules", handler)
	toUpdate := []CloudConnectorRule{
		{
			Provider:    "aws_s3",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "aws.s3.bucket",
			},
		},
		{
			Provider:    "r2",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "r2.hostname",
			},
		},
	}
	want := []CloudConnectorRule{
		{
			ID:          "some_id_1",
			Provider:    "aws_s3",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "aws.s3.bucket",
			},
		},
		{
			ID:          "some_id_2",
			Provider:    "r2",
			Expression:  "true",
			Enabled:     BoolPtr(true),
			Description: "some description",
			Parameters: CloudConnectorRuleParameters{
				Host: "r2.hostname",
			},
		},
	}
	zoneActual, err := client.UpdateZoneCloudConnectorRules(context.Background(), ZoneIdentifier(testZoneID), toUpdate)
	if assert.NoError(t, err) {
		assert.Equal(t, want, zoneActual)
	}
}
