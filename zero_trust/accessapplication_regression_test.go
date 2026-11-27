package zero_trust_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

const realisticSelfHostedAppJSON = `{
	"id": "699d98642c564d2e855e9661899b7252",
	"uid": "699d98642c564d2e855e9661899b7252",
	"aud": "737646a56ab1df6d91c7c852d637c1317e88392467c4c9799d806c7a7c66c8e0",
	"created_at": "2024-01-01T00:00:00Z",
	"updated_at": "2024-01-02T00:00:00Z",
	"name": "Admin Site",
	"domain": "app.example.com",
	"type": "self_hosted",
	"session_duration": "24h",
	"allowed_idps": ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"],
	"app_launcher_visible": true,
	"auto_redirect_to_identity": false,
	"enable_binding_cookie": false,
	"http_only_cookie_attribute": true,
	"same_site_cookie_attribute": "lax",
	"skip_interstitial": false,
	"self_hosted_domains": ["app.example.com"],
	"cors_headers": {
		"allow_all_headers": false,
		"allow_credentials": true,
		"allowed_origins": ["https://app.example.com"]
	},
	"mfa_config": {
		"allowed_authenticators": ["totp"],
		"mfa_disabled": false,
		"session_duration": "12h"
	},
	"policies": [
		{
			"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
			"name": "Allow engineers",
			"decision": "allow",
			"precedence": 1,
			"include": [{"email_domain": {"domain": "example.com"}}],
			"exclude": [],
			"require": []
		}
	]
}`

func TestAccessApplicationSelfHostedDecode(t *testing.T) {
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			var status int
			var body string
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/client/v4/accounts/account_id/access/apps":
				status = http.StatusOK
				body = `{
					"success": true,
					"errors": [],
					"messages": [],
					"result": [` + realisticSelfHostedAppJSON + `],
					"result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 1}
				}`
			case r.Method == http.MethodGet && r.URL.Path == "/client/v4/accounts/account_id/access/apps/699d98642c564d2e855e9661899b7252":
				status = http.StatusOK
				body = `{
					"success": true,
					"errors": [],
					"messages": [],
					"result": ` + realisticSelfHostedAppJSON + `
				}`
			default:
				status = http.StatusNotFound
				body = `{"success":false,"errors":[{"code":1000,"message":"not found: ` + r.URL.Path + `"}]}`
			}
			return &http.Response{
				StatusCode: status,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       io.NopCloser(strings.NewReader(body)),
				Request:    r,
			}, nil
		}),
	}

	client := cloudflare.NewClient(
		option.WithHTTPClient(httpClient),
		option.WithAPIToken("test-token"),
	)

	t.Run("List", func(t *testing.T) {
		page, err := client.ZeroTrust.Access.Applications.List(context.Background(), zero_trust.AccessApplicationListParams{
			AccountID: cloudflare.F("account_id"),
		})
		if err != nil {
			t.Fatalf("List failed: %v", err)
		}
		if len(page.Result) != 1 {
			t.Fatalf("got %d results, want 1", len(page.Result))
		}
		app := page.Result[0]
		if app.Type != zero_trust.AccessApplicationListResponseTypeSelfHosted {
			t.Errorf("app.Type = %q, want %q", app.Type, zero_trust.AccessApplicationListResponseTypeSelfHosted)
		}
		if app.SessionDuration != "24h" {
			t.Errorf("app.SessionDuration = %q, want %q", app.SessionDuration, "24h")
		}
		if !app.AppLauncherVisible {
			t.Errorf("app.AppLauncherVisible = false, want true")
		}
		if !app.HTTPOnlyCookieAttribute {
			t.Errorf("app.HTTPOnlyCookieAttribute = false, want true")
		}
		if app.SameSiteCookieAttribute != "lax" {
			t.Errorf("app.SameSiteCookieAttribute = %q, want %q", app.SameSiteCookieAttribute, "lax")
		}
		selfHosted, ok := app.AsUnion().(zero_trust.AccessApplicationListResponseSelfHostedApplication)
		if !ok {
			t.Fatalf("app.AsUnion() has type %T, want zero_trust.AccessApplicationListResponseSelfHostedApplication", app.AsUnion())
		}
		if selfHosted.SessionDuration != "24h" {
			t.Errorf("selfHosted.SessionDuration = %q, want %q", selfHosted.SessionDuration, "24h")
		}
		if selfHosted.Domain != "app.example.com" {
			t.Errorf("selfHosted.Domain = %q, want %q", selfHosted.Domain, "app.example.com")
		}
		if selfHosted.MfaConfig.SessionDuration != "12h" {
			t.Errorf("selfHosted.MfaConfig.SessionDuration = %q, want %q", selfHosted.MfaConfig.SessionDuration, "12h")
		}
		if !selfHosted.CORSHeaders.AllowCredentials {
			t.Errorf("selfHosted.CORSHeaders.AllowCredentials = false, want true")
		}
	})

	t.Run("Get", func(t *testing.T) {
		app, err := client.ZeroTrust.Access.Applications.Get(
			context.Background(),
			"699d98642c564d2e855e9661899b7252",
			zero_trust.AccessApplicationGetParams{
				AccountID: cloudflare.F("account_id"),
			},
		)
		if err != nil {
			t.Fatalf("Get failed: %v", err)
		}
		if app.Type != zero_trust.AccessApplicationGetResponseTypeSelfHosted {
			t.Errorf("app.Type = %q, want %q", app.Type, zero_trust.AccessApplicationGetResponseTypeSelfHosted)
		}
		if app.SessionDuration != "24h" {
			t.Errorf("app.SessionDuration = %q, want %q", app.SessionDuration, "24h")
		}
		selfHosted, ok := app.AsUnion().(zero_trust.AccessApplicationGetResponseSelfHostedApplication)
		if !ok {
			t.Fatalf("app.AsUnion() has type %T, want zero_trust.AccessApplicationGetResponseSelfHostedApplication", app.AsUnion())
		}
		if selfHosted.SessionDuration != "24h" {
			t.Errorf("selfHosted.SessionDuration = %q, want %q", selfHosted.SessionDuration, "24h")
		}
	})

	t.Run("NewResponse", func(t *testing.T) {
		var app zero_trust.AccessApplicationNewResponse
		if err := json.Unmarshal([]byte(realisticSelfHostedAppJSON), &app); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if app.SessionDuration != "24h" {
			t.Errorf("app.SessionDuration = %q, want %q", app.SessionDuration, "24h")
		}
		if _, ok := app.AsUnion().(zero_trust.AccessApplicationNewResponseSelfHostedApplication); !ok {
			t.Fatalf("app.AsUnion() has type %T, want zero_trust.AccessApplicationNewResponseSelfHostedApplication", app.AsUnion())
		}
	})

	t.Run("UpdateResponse", func(t *testing.T) {
		var app zero_trust.AccessApplicationUpdateResponse
		if err := json.Unmarshal([]byte(realisticSelfHostedAppJSON), &app); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if app.SessionDuration != "24h" {
			t.Errorf("app.SessionDuration = %q, want %q", app.SessionDuration, "24h")
		}
		if _, ok := app.AsUnion().(zero_trust.AccessApplicationUpdateResponseSelfHostedApplication); !ok {
			t.Fatalf("app.AsUnion() has type %T, want zero_trust.AccessApplicationUpdateResponseSelfHostedApplication", app.AsUnion())
		}
	})
}

func TestAccessApplicationUnionDiscriminatorVariants(t *testing.T) {
	cases := []struct {
		appType  string
		wantType reflect.Type
	}{
		{"end_user", reflect.TypeOf(zero_trust.AccessApplicationListResponseEndUserApplication{})},
		{"self_hosted", reflect.TypeOf(zero_trust.AccessApplicationListResponseSelfHostedApplication{})},
		{"saas", reflect.TypeOf(zero_trust.AccessApplicationListResponseSaaSApplication{})},
		{"ssh", reflect.TypeOf(zero_trust.AccessApplicationListResponseBrowserSSHApplication{})},
		{"vnc", reflect.TypeOf(zero_trust.AccessApplicationListResponseBrowserVNCApplication{})},
		{"app_launcher", reflect.TypeOf(zero_trust.AccessApplicationListResponseAppLauncherApplication{})},
		{"warp", reflect.TypeOf(zero_trust.AccessApplicationListResponseDeviceEnrollmentPermissionsApplication{})},
		{"biso", reflect.TypeOf(zero_trust.AccessApplicationListResponseBrowserIsolationPermissionsApplication{})},
		{"proxy_endpoint", reflect.TypeOf(zero_trust.AccessApplicationListResponseGatewayIdentityProxyEndpointApplication{})},
		{"bookmark", reflect.TypeOf(zero_trust.AccessApplicationListResponseBookmarkApplication{})},
		{"infrastructure", reflect.TypeOf(zero_trust.AccessApplicationListResponseInfrastructureApplication{})},
		{"rdp", reflect.TypeOf(zero_trust.AccessApplicationListResponseBrowserRDPApplication{})},
		{"mcp", reflect.TypeOf(zero_trust.AccessApplicationListResponseMcpServerApplication{})},
		{"mcp_portal", reflect.TypeOf(zero_trust.AccessApplicationListResponseMcpServerPortalApplication{})},
	}

	for _, tc := range cases {
		t.Run(tc.appType, func(t *testing.T) {
			payload := []byte(`{"id":"app-1","name":"App","domain":"app.example.com","type":"` + tc.appType + `","session_duration":"24h","created_at":"2024-01-01T00:00:00Z"}`)
			var app zero_trust.AccessApplicationListResponse
			if err := json.Unmarshal(payload, &app); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			gotType := reflect.TypeOf(app.AsUnion())
			if gotType != tc.wantType {
				t.Errorf("type=%q decoded as %v, want %v", tc.appType, gotType, tc.wantType)
			}
		})
	}

	t.Run("unmapped_type_falls_back_cleanly", func(t *testing.T) {
		payload := []byte(`{"id":"app-1","name":"Dash SSO","domain":"dash.example.com","type":"dash_sso"}`)
		var app zero_trust.AccessApplicationListResponse
		if err := json.Unmarshal(payload, &app); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if app.AsUnion() == nil {
			t.Fatal("expected non-nil union fallback for dash_sso")
		}
		if app.Type != zero_trust.AccessApplicationListResponseTypeDashSSO {
			t.Errorf("app.Type = %q, want %q", app.Type, zero_trust.AccessApplicationListResponseTypeDashSSO)
		}
	})
}
