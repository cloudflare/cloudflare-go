package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamsCertificate(t *testing.T) {
	setup()
	defer teardown()

	testCertID := "80c8a54e-d55c-46c6-86bb-e8a3c90472f4"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"enabled": false,
				"id": "80c8a54e-d55c-46c6-86bb-e8a3c90472f4",
				"binding_status": "inactive",
				"qs_pack_id": "00000000-0000-0000-0000-000000000000",
				"type": "gateway_managed",
				"updated_at": "0001-01-01T00:00:00Z",
				"uploaded_on": "0001-01-01T00:00:00Z",
				"created_at": "2024-06-19T08:28:34.235742Z",
				"expires_on": "2029-06-19T08:24:00Z"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates/"+testCertID, handler)

	actual, err := client.TeamsCertificate(context.Background(), testAccountID, testCertID)

	if assert.NoError(t, err) {
		assert.Equal(t, *actual.Enabled, false)
		assert.Equal(t, actual.ID, testCertID)
		assert.Equal(t, actual.BindingStatus, "inactive")
		assert.Equal(t, actual.Type, "gateway_managed")
	}
}

func TestTeamsCertificatesList(t *testing.T) {
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
					"enabled": false,
					"id": "43a36083-987b-4321-95ab-4052771c4e6f",
					"binding_status": "inactive",
					"qs_pack_id": "00000000-0000-0000-0000-000000000000",
					"type": "custom",
					"updated_at": "0001-01-01T00:00:00Z",
					"uploaded_on": "2022-12-08T22:58:13.646596Z",
					"created_at": "0001-01-01T00:00:00Z",
					"expires_on": "2122-10-29T16:59:47Z"
				},
				{
					"enabled": false,
					"id": "4a9d6ecf-0fdd-4676-818a-ee6b45f17f9b",
					"binding_status": "inactive",
					"qs_pack_id": "00000000-0000-0000-0000-000000000000",
					"type": "gateway_managed",
					"updated_at": "0001-01-01T00:00:00Z",
					"uploaded_on": "0001-01-01T00:00:00Z",
					"created_at": "2024-05-29T01:32:21.133597Z",
					"expires_on": "2029-05-29T01:27:00Z"
				}
			]
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates", handler)

	actual, err := client.TeamsCertificates(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, len(actual), 2)
		assert.Equal(t, actual[0].ID, "43a36083-987b-4321-95ab-4052771c4e6f")
		assert.Equal(t, actual[1].ID, "4a9d6ecf-0fdd-4676-818a-ee6b45f17f9b")
	}
}

func TestTeamsAccountGenerateCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'post', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"enabled": false,
				"id": "80c8a54e-d55c-46c6-86bb-e8a3c90472f4",
				"binding_status": "inactive",
				"qs_pack_id": "00000000-0000-0000-0000-000000000000",
				"type": "gateway_managed",
				"updated_at": "0001-01-01T00:00:00Z",
				"uploaded_on": "0001-01-01T00:00:00Z",
				"created_at": "2024-06-19T08:28:34.235742Z",
				"expires_on": "2029-06-19T08:24:00Z"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates", handler)

	createRequest := TeamsCertificateCreateRequest{
		ValidityPeriodDays: 1826,
	}
	actual, err := client.TeamsGenerateCertificate(context.Background(), testAccountID, createRequest)

	if assert.NoError(t, err) {
		assert.Equal(t, actual.ID, "80c8a54e-d55c-46c6-86bb-e8a3c90472f4")
	}
}

func TestTeamsActivateCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"enabled": false,
				"id": "80c8a54e-d55c-46c6-86bb-e8a3c90472f4",
				"binding_status": "pending_deployment",
				"qs_pack_id": "00000000-0000-0000-0000-000000000000",
				"type": "gateway_managed",
				"updated_at": "0001-01-01T00:00:00Z",
				"uploaded_on": "0001-01-01T00:00:00Z",
				"created_at": "2024-06-19T08:28:34.235742Z",
				"expires_on": "2029-06-19T08:24:00Z"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates/80c8a54e-d55c-46c6-86bb-e8a3c90472f4/activate", handler)

	cert, err := client.TeamsActivateCertificate(context.Background(), testAccountID, "80c8a54e-d55c-46c6-86bb-e8a3c90472f4")

	assert.NoError(t, err)
	assert.Equal(t, cert.BindingStatus, "pending_deployment")
}

func TestTeamsDeactivateCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"enabled": false,
				"id": "80c8a54e-d55c-46c6-86bb-e8a3c90472f4",
				"binding_status": "pending_deletion",
				"qs_pack_id": "00000000-0000-0000-0000-000000000000",
				"type": "gateway_managed",
				"updated_at": "0001-01-01T00:00:00Z",
				"uploaded_on": "0001-01-01T00:00:00Z",
				"created_at": "2024-06-19T08:28:34.235742Z",
				"expires_on": "2029-06-19T08:24:00Z"
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates/80c8a54e-d55c-46c6-86bb-e8a3c90472f4/deactivate", handler)

	cert, err := client.TeamsDeactivateCertificate(context.Background(), testAccountID, "80c8a54e-d55c-46c6-86bb-e8a3c90472f4")

	assert.NoError(t, err)
	assert.Equal(t, cert.BindingStatus, "pending_deletion")
}

func TestTeamsDeleteCertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'Delete', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": nil
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/certificates/80c8a54e-d55c-46c6-86bb-e8a3c90472f4", handler)

	err := client.TeamsDeleteCertificate(context.Background(), testAccountID, "80c8a54e-d55c-46c6-86bb-e8a3c90472f4")

	assert.NoError(t, err)
}
