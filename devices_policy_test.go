package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	deviceSettingsPolicyID         = "a842fa8a-a583-482e-9cd9-eb43362949fd"
	deviceSettingsPolicyMatch      = "identity.email == \"test@example.com\""
	deviceSettingsPolicyPrecedence = 10

	defaultDeviceSettingsPolicy = DeviceSettingsPolicy{
		ServiceModeV2: &ServiceModeV2{
			Mode: "warp",
		},
		DisableAutoFallback: BoolPtr(false),
		FallbackDomains: &[]FallbackDomain{
			{Suffix: "invalid"},
			{Suffix: "test"},
		},
		Exclude: &[]SplitTunnel{
			{Address: "10.0.0.0/8"},
			{Address: "100.64.0.0/10"},
		},
		GatewayUniqueID: StringPtr("t1235"),
		SupportURL:      StringPtr(""),
		CaptivePortal:   IntPtr(180),
		AllowModeSwitch: BoolPtr(false),
		SwitchLocked:    BoolPtr(false),
		AllowUpdates:    BoolPtr(false),
		AutoConnect:     IntPtr(0),
		AllowedToLeave:  BoolPtr(true),
		Enabled:         BoolPtr(true),
		PolicyID:        nil,
		Name:            nil,
		Match:           nil,
		Precedence:      nil,
		Default:         true,
	}

	nonDefaultDeviceSettingsPolicy = DeviceSettingsPolicy{
		ServiceModeV2: &ServiceModeV2{
			Mode: "warp",
		},
		DisableAutoFallback: BoolPtr(false),
		FallbackDomains: &[]FallbackDomain{
			{Suffix: "invalid"},
			{Suffix: "test"},
		},
		Exclude: &[]SplitTunnel{
			{Address: "10.0.0.0/8"},
			{Address: "100.64.0.0/10"},
		},
		GatewayUniqueID: StringPtr("t1235"),
		SupportURL:      StringPtr(""),
		CaptivePortal:   IntPtr(180),
		AllowModeSwitch: BoolPtr(false),
		SwitchLocked:    BoolPtr(false),
		AllowUpdates:    BoolPtr(false),
		AutoConnect:     IntPtr(0),
		AllowedToLeave:  BoolPtr(true),
		PolicyID:        &deviceSettingsPolicyID,
		Enabled:         BoolPtr(true),
		Name:            StringPtr("test"),
		Match:           &deviceSettingsPolicyMatch,
		Precedence:      &deviceSettingsPolicyPrecedence,
		Default:         false,
	}

	defaultDeviceSettingsPolicyJson = `{
		"service_mode_v2": {
			"mode": "warp"
		},
		"disable_auto_fallback": false,
		"fallback_domains": [
			{
				"suffix": "invalid"
			},
			{
				"suffix": "test"
			}
		],
		"exclude": [
			{
				"address": "10.0.0.0/8"
			},
			{
				"address": "100.64.0.0/10"
			}
		],
		"gateway_unique_id": "t1235",
		"support_url": "",
		"captive_portal": 180,
		"allow_mode_switch": false,
		"switch_locked": false,
		"allow_updates": false,
		"auto_connect": 0,
		"allowed_to_leave": true,
		"enabled": true,
		"default": true
	}`

	nonDefaultDeviceSettingsPolicyJson = fmt.Sprintf(`{
		"service_mode_v2": {
			"mode": "warp"
		},
		"disable_auto_fallback": false,
		"fallback_domains": [
			{
				"suffix": "invalid"
			},
			{
				"suffix": "test"
			}
		],
		"exclude": [
			{
				"address": "10.0.0.0/8"
			},
			{
				"address": "100.64.0.0/10"
			}
		],
		"gateway_unique_id": "t1235",
		"support_url": "",
		"captive_portal": 180,
		"allow_mode_switch": false,
		"switch_locked": false,
		"allow_updates": false,
		"auto_connect": 0,
		"allowed_to_leave": true,
		"policy_id": "%s",
		"enabled": true,
		"name": "test",
		"match": %#v,
		"precedence": 10,
		"default": false
	}`, deviceSettingsPolicyID, deviceSettingsPolicyMatch)
)

func TestUpdateDeviceClientCertificatesZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": {"enabled": true}
		}`)
	}

	want := DeviceClientCertificatesZone{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: Enabled{true},
	}

	mux.HandleFunc("/zones/"+testZoneID+"/devices/policy/certificates", handler)

	actual, err := client.UpdateDeviceClientCertificatesZone(context.Background(), testZoneID, true)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetDeviceClientCertificatesZone(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": {"enabled": false}
		}`)
	}

	want := DeviceClientCertificatesZone{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: Enabled{false},
	}

	mux.HandleFunc("/zones/"+testZoneID+"/devices/policy/certificates", handler)

	actual, err := client.GetDeviceClientCertificatesZone(context.Background(), testZoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateDeviceSettingsPolicy(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": %s
		}`, nonDefaultDeviceSettingsPolicyJson)
	}

	want := DeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: nonDefaultDeviceSettingsPolicy,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy", handler)

	actual, err := client.CreateDeviceSettingsPolicy(context.Background(), testAccountID, DeviceSettingsPolicyRequest{
		Precedence: IntPtr(10),
		Match:      &deviceSettingsPolicyMatch,
		Name:       StringPtr("test"),
	})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateDefaultDeviceSettingsPolicy(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": %s
		}`, defaultDeviceSettingsPolicyJson)
	}

	want := DeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: defaultDeviceSettingsPolicy,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy", handler)

	actual, err := client.UpdateDefaultDeviceSettingsPolicy(context.Background(), testAccountID, DeviceSettingsPolicyRequest{})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateDeviceSettingsPolicy(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": %s
		}`, nonDefaultDeviceSettingsPolicyJson)
	}

	precedence := 10
	want := DeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: nonDefaultDeviceSettingsPolicy,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/"+deviceSettingsPolicyID, handler)

	actual, err := client.UpdateDeviceSettingsPolicy(context.Background(), testAccountID, deviceSettingsPolicyID, DeviceSettingsPolicyRequest{
		Precedence: &precedence,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteDeviceSettingsPolicy(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": [ %s ]
		}`, defaultDeviceSettingsPolicyJson)
	}

	want := DeleteDeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: []DeviceSettingsPolicy{defaultDeviceSettingsPolicy},
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/"+deviceSettingsPolicyID, handler)

	actual, err := client.DeleteDeviceSettingsPolicy(context.Background(), testAccountID, deviceSettingsPolicyID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetDefaultDeviceSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": %s
		}`, defaultDeviceSettingsPolicyJson)
	}

	want := DeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: defaultDeviceSettingsPolicy,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy", handler)

	actual, err := client.GetDefaultDeviceSettingsPolicy(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetDeviceSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": null,
			"messages": null,
			"result": %s
		}`, nonDefaultDeviceSettingsPolicyJson)
	}

	want := DeviceSettingsPolicyResponse{
		Response: Response{
			Success:  true,
			Errors:   nil,
			Messages: nil,
		},
		Result: nonDefaultDeviceSettingsPolicy,
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/devices/policy/"+deviceSettingsPolicyID, handler)

	actual, err := client.GetDeviceSettingsPolicy(context.Background(), testAccountID, deviceSettingsPolicyID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
