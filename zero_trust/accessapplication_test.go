// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestAccessApplicationNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.New(context.TODO(), zero_trust.AccessApplicationNewParams{
		Body: zero_trust.AccessApplicationNewParamsBodySelfHostedApplication{
			Domain:                   cloudflare.F("test.example.com/admin"),
			Type:                     cloudflare.F("self_hosted"),
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AllowedIdPs:              cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F(true),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CORSHeaders: cloudflare.F(zero_trust.CORSHeadersParam{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]zero_trust.AllowedHeadersParam{"string"}),
				AllowedMethods:   cloudflare.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
				AllowedOrigins:   cloudflare.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
				MaxAge:           cloudflare.F(-1.000000),
			}),
			CustomDenyMessage:        cloudflare.F("custom_deny_message"),
			CustomDenyURL:            cloudflare.F("custom_deny_url"),
			CustomNonIdentityDenyURL: cloudflare.F("custom_non_identity_deny_url"),
			CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252"}),
			Destinations: cloudflare.F([]zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestination{{
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestinationsTypePublic),
				URI:  cloudflare.F("test.example.com/admin"),
			}, {
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestinationsTypePublic),
				URI:  cloudflare.F("test.anotherexample.com/staff"),
			}, {
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestinationsTypePublic),
				URI:  cloudflare.F("10.5.0.2"),
			}, {
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestinationsTypePublic),
				URI:  cloudflare.F("10.5.0.3/32:1234-4321"),
			}, {
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationDestinationsTypePublic),
				URI:  cloudflare.F("private-sni.example.com"),
			}}),
			EnableBindingCookie:     cloudflare.F(true),
			HTTPOnlyCookieAttribute: cloudflare.F(true),
			LogoURL:                 cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                    cloudflare.F("Admin Site"),
			OptionsPreflightBypass:  cloudflare.F(true),
			PathCookieAttribute:     cloudflare.F(true),
			Policies: cloudflare.F([]zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
				ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				Precedence: cloudflare.F(int64(0)),
			}}),
			SameSiteCookieAttribute: cloudflare.F("strict"),
			SCIMConfig: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig{
				IdPUID:    cloudflare.F("idp_uid"),
				RemoteURI: cloudflare.F("remote_uri"),
				Authentication: cloudflare.F[zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
					Password: cloudflare.F("password"),
					Scheme:   cloudflare.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
					User:     cloudflare.F("user"),
				}),
				DeactivateOnDelete: cloudflare.F(true),
				Enabled:            cloudflare.F(true),
				Mappings: cloudflare.F([]zero_trust.SCIMConfigMappingParam{{
					Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
					Enabled: cloudflare.F(true),
					Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
					Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
						Create: cloudflare.F(true),
						Delete: cloudflare.F(true),
						Update: cloudflare.F(true),
					}),
					Strictness:       cloudflare.F(zero_trust.SCIMConfigMappingStrictnessStrict),
					TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
				}}),
			}),
			SelfHostedDomains:      cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
			ServiceAuth401Redirect: cloudflare.F(true),
			SessionDuration:        cloudflare.F("24h"),
			SkipInterstitial:       cloudflare.F(true),
			Tags:                   cloudflare.F([]string{"engineers"}),
		},
		AccountID: cloudflare.F("account_id"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationUpdateParams{
			Body: zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplication{
				Domain:                   cloudflare.F("test.example.com/admin"),
				Type:                     cloudflare.F("self_hosted"),
				AllowAuthenticateViaWARP: cloudflare.F(true),
				AllowedIdPs:              cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252"}),
				AppLauncherVisible:       cloudflare.F(true),
				AutoRedirectToIdentity:   cloudflare.F(true),
				CORSHeaders: cloudflare.F(zero_trust.CORSHeadersParam{
					AllowAllHeaders:  cloudflare.F(true),
					AllowAllMethods:  cloudflare.F(true),
					AllowAllOrigins:  cloudflare.F(true),
					AllowCredentials: cloudflare.F(true),
					AllowedHeaders:   cloudflare.F([]zero_trust.AllowedHeadersParam{"string"}),
					AllowedMethods:   cloudflare.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
					AllowedOrigins:   cloudflare.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
					MaxAge:           cloudflare.F(-1.000000),
				}),
				CustomDenyMessage:        cloudflare.F("custom_deny_message"),
				CustomDenyURL:            cloudflare.F("custom_deny_url"),
				CustomNonIdentityDenyURL: cloudflare.F("custom_non_identity_deny_url"),
				CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252"}),
				Destinations: cloudflare.F([]zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestination{{
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestinationsTypePublic),
					URI:  cloudflare.F("test.example.com/admin"),
				}, {
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestinationsTypePublic),
					URI:  cloudflare.F("test.anotherexample.com/staff"),
				}, {
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestinationsTypePublic),
					URI:  cloudflare.F("10.5.0.2"),
				}, {
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestinationsTypePublic),
					URI:  cloudflare.F("10.5.0.3/32:1234-4321"),
				}, {
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationDestinationsTypePublic),
					URI:  cloudflare.F("private-sni.example.com"),
				}}),
				EnableBindingCookie:     cloudflare.F(true),
				HTTPOnlyCookieAttribute: cloudflare.F(true),
				LogoURL:                 cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				Name:                    cloudflare.F("Admin Site"),
				OptionsPreflightBypass:  cloudflare.F(true),
				PathCookieAttribute:     cloudflare.F(true),
				Policies: cloudflare.F([]zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}}),
				SameSiteCookieAttribute: cloudflare.F("strict"),
				SCIMConfig: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig{
					IdPUID:    cloudflare.F("idp_uid"),
					RemoteURI: cloudflare.F("remote_uri"),
					Authentication: cloudflare.F[zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
						Password: cloudflare.F("password"),
						Scheme:   cloudflare.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
						User:     cloudflare.F("user"),
					}),
					DeactivateOnDelete: cloudflare.F(true),
					Enabled:            cloudflare.F(true),
					Mappings: cloudflare.F([]zero_trust.SCIMConfigMappingParam{{
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						Strictness:       cloudflare.F(zero_trust.SCIMConfigMappingStrictnessStrict),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}}),
				}),
				SelfHostedDomains:      cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
				ServiceAuth401Redirect: cloudflare.F(true),
				SessionDuration:        cloudflare.F("24h"),
				SkipInterstitial:       cloudflare.F(true),
				Tags:                   cloudflare.F([]string{"engineers"}),
			},
			AccountID: cloudflare.F("account_id"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.List(context.TODO(), zero_trust.AccessApplicationListParams{
		AccountID: cloudflare.F("account_id"),
		AUD:       cloudflare.F("aud"),
		Domain:    cloudflare.F("domain"),
		Name:      cloudflare.F("name"),
		Search:    cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationDeleteWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationDeleteParams{
			AccountID: cloudflare.F("account_id"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationGetParams{
			AccountID: cloudflare.F("account_id"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessApplicationRevokeTokensWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.RevokeTokens(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationRevokeTokensParams{
			AccountID: cloudflare.F("account_id"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
