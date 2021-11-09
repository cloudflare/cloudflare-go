package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamsAccount(t *testing.T) {
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
				"id": "%s",
				"provider_name": "cf",
				"gateway_tag": "1234"
			}
		}
		`, testAccountID)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway", handler)

	actual, err := client.TeamsAccount(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual.ProviderName, "cf")
		assert.Equal(t, actual.GatewayTag, "1234")
		assert.Equal(t, actual.ID, testAccountID)
	}
}

func TestTeamsAccountConfiguration(t *testing.T) {
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
				"settings": {
					"antivirus": {
						"enabled_download_phase": true
					},
					"tls_decrypt": {
						"enabled": true
					},
					"fips": {
						"tls": true
					},
					"activity_log": {
						"enabled": true
					}
				}
			}
		}
		`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/configuration", handler)

	actual, err := client.TeamsAccountConfiguration(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual.Settings, TeamsAccountSettings{
			Antivirus:   &TeamsAntivirus{EnabledDownloadPhase: true},
			ActivityLog: &TeamsActivityLog{Enabled: true},
			TLSDecrypt:  &TeamsTLSDecrypt{Enabled: true},
			FIPS:        &TeamsFIPS{TLS: true},
		})
	}
}

func TestTeamsAccountUpdateConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'put', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"settings": {
					"antivirus": {
						"enabled_download_phase": false
					},
					"tls_decrypt": {
						"enabled": true
					},
					"activity_log": {
						"enabled": true
					}
				}
			}
		}
		`)
	}

	settings := TeamsAccountSettings{
		Antivirus:   &TeamsAntivirus{EnabledDownloadPhase: false},
		ActivityLog: &TeamsActivityLog{Enabled: true},
		TLSDecrypt:  &TeamsTLSDecrypt{Enabled: true},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/configuration", handler)

	configuration := TeamsConfiguration{
		Settings: settings,
	}
	actual, err := client.TeamsAccountUpdateConfiguration(context.Background(), testAccountID, configuration)

	if assert.NoError(t, err) {

		assert.Equal(t, actual, configuration)
	}
}
