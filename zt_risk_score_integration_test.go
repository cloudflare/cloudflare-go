package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/goccy/go-json"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRiskScoreIntegrations(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `
		{
			"result": [
				{
					"id": "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
					"account_tag": "f772f54476db1178347e5bfba3c9ac83",
					"integration_type": "Okta",
					"reference_id": "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
					"tenant_url": "https://my-tenant.oktapreview.com",
					"well_known_url": "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
					"active": true,
					"created_at": "2024-06-13T18:17:22Z"
				},
				{
					"id": "0d5e7916-0b18-4a22-839f-b82e33a87843",
					"account_tag": "f772f54476db1178347e5bfba3c9ac83",
					"integration_type": "Okta",
					"reference_id": "d77c9506-00bc-43aa-8e95-3223351ea330",
					"tenant_url": "https://my-tenant-2.oktapreview.com",
					"well_known_url": "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/0d5e7916-0b18-4a22-839f-b82e33a87843",
					"active": false,
					"created_at": "2024-06-13T18:17:22Z"
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}
		`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2024-06-13T18:17:22Z")

	want := []RiskScoreIntegration{
		{
			ID:              "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
			AccountTag:      "f772f54476db1178347e5bfba3c9ac83",
			IntegrationType: "Okta",
			ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
			TenantUrl:       "https://my-tenant.oktapreview.com",
			WellKnownUrl:    "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
			Active:          BoolPtr(true),
			CreatedAt:       &createdAt,
		},
		{
			ID:              "0d5e7916-0b18-4a22-839f-b82e33a87843",
			AccountTag:      "f772f54476db1178347e5bfba3c9ac83",
			IntegrationType: "Okta",
			ReferenceID:     "d77c9506-00bc-43aa-8e95-3223351ea330",
			TenantUrl:       "https://my-tenant-2.oktapreview.com",
			WellKnownUrl:    "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/0d5e7916-0b18-4a22-839f-b82e33a87843",
			Active:          BoolPtr(false),
			CreatedAt:       &createdAt,
		},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/zt_risk_scoring/integrations", handler)

	actual, err := client.ListRiskScoreIntegrations(context.Background(), AccountIdentifier(testAccountID), ListRiskScoreIntegrationParams{})
	require.NoError(t, err)
	require.Equal(t, want, actual)
}

func TestGetRiskScoreIntegration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"account_tag": "f772f54476db1178347e5bfba3c9ac83",
				"integration_type": "Okta",
				"reference_id": "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
				"tenant_url": "https://my-tenant.oktapreview.com",
				"well_known_url": "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"active": true,
				"created_at": "2024-06-13T18:17:22Z"
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2024-06-13T18:17:22Z")

	want := RiskScoreIntegration{
		ID:              "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		AccountTag:      "f772f54476db1178347e5bfba3c9ac83",
		IntegrationType: "Okta",
		ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
		TenantUrl:       "https://my-tenant.oktapreview.com",
		WellKnownUrl:    "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		Active:          BoolPtr(true),
		CreatedAt:       &createdAt,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/zt_risk_scoring/integrations/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9", handler)

	actual, err := client.GetRiskScoreIntegration(context.Background(), AccountIdentifier(testAccountID), "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9")
	require.NoError(t, err)
	require.Equal(t, want, actual)
}

func TestCreateRiskScoreIntegration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		var reqBody RiskScoreIntegrationCreateRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.Nil(t, err)
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"account_tag": "f772f54476db1178347e5bfba3c9ac83",
				"integration_type": "`+reqBody.IntegrationType+`",
				"reference_id": "`+reqBody.ReferenceID+`",
				"tenant_url": "`+reqBody.TenantUrl+`",
				"well_known_url": "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"active": true,
				"created_at": "2024-06-13T18:17:22Z"
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2024-06-13T18:17:22Z")

	want := RiskScoreIntegration{
		ID:              "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		AccountTag:      "f772f54476db1178347e5bfba3c9ac83",
		IntegrationType: "Okta",
		ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
		TenantUrl:       "https://my-tenant.oktapreview.com",
		WellKnownUrl:    "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		Active:          BoolPtr(true),
		CreatedAt:       &createdAt,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/zt_risk_scoring/integrations", handler)

	integration := RiskScoreIntegrationCreateRequest{
		IntegrationType: "Okta",
		TenantUrl:       "https://my-tenant.oktapreview.com",
		ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
	}
	actual, err := client.CreateRiskScoreIntegration(context.Background(), AccountIdentifier(testAccountID), integration)
	require.NoError(t, err)
	require.Equal(t, want, actual)
}

func TestUpdateRiskScoreIntegration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		var reqBody RiskScoreIntegrationUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.Nil(t, err)

		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"id": "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"account_tag": "f772f54476db1178347e5bfba3c9ac83",
				"integration_type": "`+reqBody.IntegrationType+`",
				"reference_id": "`+reqBody.ReferenceID+`",
				"tenant_url": "`+reqBody.TenantUrl+`",
				"well_known_url": "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
				"active": `+strconv.FormatBool(*reqBody.Active)+`,
				"created_at": "2024-06-13T18:17:22Z"
			}
		}`)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2024-06-13T18:17:22Z")

	want := RiskScoreIntegration{
		ID:              "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		AccountTag:      "f772f54476db1178347e5bfba3c9ac83",
		IntegrationType: "Okta",
		ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
		TenantUrl:       "https://my-tenant.oktapreview.com",
		WellKnownUrl:    "https://ssf.risk-score.teams.cloudflare.com/.well-known/sse-configuration/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9",
		Active:          BoolPtr(false),
		CreatedAt:       &createdAt,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/zt_risk_scoring/integrations/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9", handler)

	integration := RiskScoreIntegrationUpdateRequest{
		IntegrationType: "Okta",
		TenantUrl:       "https://my-tenant.oktapreview.com",
		ReferenceID:     "72f826ea-8744-4dcb-a00e-8cc00bb6d34b",
		Active:          BoolPtr(false),
	}

	actual, err := client.UpdateRiskScoreIntegration(context.Background(), AccountIdentifier(testAccountID), "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9", integration)
	require.NoError(t, err)
	require.Equal(t, want, actual)
}

func TestDeleteRiskScoreIntegration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": null,
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/zt_risk_scoring/integrations/91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9", handler)

	err := client.DeleteRiskScoreIntegration(context.Background(), AccountIdentifier(testAccountID), "91dbf25d-2b51-415e-a6df-5a6d2ff4f7e9")
	require.NoError(t, err)
}
