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

func TestTeamsAccountGetLoggingConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {"settings_by_rule_type":{"dns":{"log_all":false,"log_blocks":true}},"redact_pii":true}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/logging", handler)

	actual, err := client.TeamsAccountLoggingConfiguration(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual, TeamsLoggingSettings{
			RedactPii: true,
			LoggingSettingsByRuleType: map[TeamsRuleType]TeamsAccountLoggingConfiguration{
				TeamsDNSRuleType: {LogAll: false, LogBlocks: true},
			},
		})
	}
}

func TestTeamsAccountUpdateLoggingConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {"settings_by_rule_type":{"dns":{"log_all":false,"log_blocks":true}, "http":{"log_all":true,"log_blocks":false}, "l4": {"log_all": false, "log_blocks": true}},"redact_pii":true}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/gateway/logging", handler)

	actual, err := client.TeamsAccountUpdateLoggingConfiguration(context.Background(), testAccountID, TeamsLoggingSettings{
		RedactPii: true,
		LoggingSettingsByRuleType: map[TeamsRuleType]TeamsAccountLoggingConfiguration{
			TeamsDNSRuleType: {
				LogAll:    false,
				LogBlocks: true,
			},
			TeamsHTTPRuleType: {
				LogAll: true,
			},
			TeamsL4RuleType: {
				LogBlocks: true,
			},
		},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, actual, TeamsLoggingSettings{
			RedactPii: true,
			LoggingSettingsByRuleType: map[TeamsRuleType]TeamsAccountLoggingConfiguration{
				TeamsDNSRuleType:  {LogAll: false, LogBlocks: true},
				TeamsHTTPRuleType: {LogAll: true, LogBlocks: false},
				TeamsL4RuleType:   {LogAll: false, LogBlocks: true},
			},
		})
	}
}

func TestTeamsAccountGetDeviceConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {"gateway_proxy_enabled": true,"gateway_udp_proxy_enabled":false}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/settings", handler)

	actual, err := client.TeamsAccountDeviceConfiguration(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual, TeamsDeviceSettings{
			GatewayProxyEnabled:    true,
			GatewayProxyUDPEnabled: false,
		})
	}
}

func TestTeamsAccountUpdateDeviceConfiguration(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {"gateway_proxy_enabled": true,"gateway_udp_proxy_enabled":true}
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/settings", handler)

	actual, err := client.TeamsAccountDeviceUpdateConfiguration(context.Background(), testAccountID, TeamsDeviceSettings{
		GatewayProxyUDPEnabled: true,
		GatewayProxyEnabled:    true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, actual, TeamsDeviceSettings{
			GatewayProxyEnabled:    true,
			GatewayProxyUDPEnabled: true,
		})
	}
}
