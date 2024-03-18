// File generated from our OpenAPI spec by Stainless.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestAccessApplicationNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		Domain:                   cloudflare.F[any]("https://mybookmark.com"),
		Type:                     cloudflare.F(zero_trust.AccessApplicationNewParamsTypeBookmark),
		AccountID:                cloudflare.F("string"),
		ZoneID:                   cloudflare.F("string"),
		AllowAuthenticateViaWARP: cloudflare.F(true),
		AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
		AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
		AutoRedirectToIdentity:   cloudflare.F(true),
		CorsHeaders: cloudflare.F(zero_trust.AccessApplicationNewParamsCorsHeaders{
			AllowAllHeaders:  cloudflare.F(true),
			AllowAllMethods:  cloudflare.F(true),
			AllowAllOrigins:  cloudflare.F(true),
			AllowCredentials: cloudflare.F(true),
			AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
			AllowedMethods:   cloudflare.F([]zero_trust.AccessApplicationNewParamsCorsHeadersAllowedMethod{zero_trust.AccessApplicationNewParamsCorsHeadersAllowedMethodGet}),
			AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
			MaxAge:           cloudflare.F(-1.000000),
		}),
		CustomDenyMessage:        cloudflare.F("string"),
		CustomDenyURL:            cloudflare.F("string"),
		CustomNonIdentityDenyURL: cloudflare.F("string"),
		CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
		EnableBindingCookie:      cloudflare.F(true),
		HTTPOnlyCookieAttribute:  cloudflare.F(true),
		LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
		Name:                     cloudflare.F("Admin Site"),
		PathCookieAttribute:      cloudflare.F(true),
		SaasApp: cloudflare.F[zero_trust.AccessApplicationNewParamsSaasApp](zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasApp(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasApp{
			AuthType:           cloudflare.F(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
			ConsumerServiceURL: cloudflare.F("https://example.com"),
			CustomAttributes: cloudflare.F(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes{
				Name:       cloudflare.F("family_name"),
				NameFormat: cloudflare.F(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
				Source: cloudflare.F(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
					Name: cloudflare.F("last_name"),
				}),
			}),
			DefaultRelayState:      cloudflare.F("https://example.com"),
			IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
			NameIDFormat:           cloudflare.F(zero_trust.AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatID),
			NameIDTransformJsonata: cloudflare.F("$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
			PublicKey:              cloudflare.F("example unique name"),
			SpEntityID:             cloudflare.F("example unique name"),
			SSOEndpoint:            cloudflare.F("https://example.cloudflareaccess.com/cdn-cgi/access/sso/saml/b3f58a2b414e0b51d45c8c2af26fccca0e27c63763c426fa52f98dcf0b3b3bfd"),
		})),
		SameSiteCookieAttribute: cloudflare.F("strict"),
		SelfHostedDomains:       cloudflare.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
		ServiceAuth401Redirect:  cloudflare.F(true),
		SessionDuration:         cloudflare.F("24h"),
		SkipInterstitial:        cloudflare.F(true),
		Tags:                    cloudflare.F([]string{"engineers", "engineers", "engineers"}),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		zero_trust.AccessApplicationUpdateParams{
			Domain:                   cloudflare.F[any]("https://mybookmark.com"),
			Type:                     cloudflare.F(zero_trust.AccessApplicationUpdateParamsTypeBookmark),
			AccountID:                cloudflare.F("string"),
			ZoneID:                   cloudflare.F("string"),
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CorsHeaders: cloudflare.F(zero_trust.AccessApplicationUpdateParamsCorsHeaders{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				AllowedMethods:   cloudflare.F([]zero_trust.AccessApplicationUpdateParamsCorsHeadersAllowedMethod{zero_trust.AccessApplicationUpdateParamsCorsHeadersAllowedMethodGet}),
				AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
				MaxAge:           cloudflare.F(-1.000000),
			}),
			CustomDenyMessage:        cloudflare.F("string"),
			CustomDenyURL:            cloudflare.F("string"),
			CustomNonIdentityDenyURL: cloudflare.F("string"),
			CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			EnableBindingCookie:      cloudflare.F(true),
			HTTPOnlyCookieAttribute:  cloudflare.F(true),
			LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                     cloudflare.F("Admin Site"),
			PathCookieAttribute:      cloudflare.F(true),
			SaasApp: cloudflare.F[zero_trust.AccessApplicationUpdateParamsSaasApp](zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp{
				AuthType:           cloudflare.F(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
				ConsumerServiceURL: cloudflare.F("https://example.com"),
				CustomAttributes: cloudflare.F(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes{
					Name:       cloudflare.F("family_name"),
					NameFormat: cloudflare.F(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
					Source: cloudflare.F(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
						Name: cloudflare.F("last_name"),
					}),
				}),
				DefaultRelayState:      cloudflare.F("https://example.com"),
				IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
				NameIDFormat:           cloudflare.F(zero_trust.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatID),
				NameIDTransformJsonata: cloudflare.F("$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"),
				PublicKey:              cloudflare.F("example unique name"),
				SpEntityID:             cloudflare.F("example unique name"),
				SSOEndpoint:            cloudflare.F("https://example.cloudflareaccess.com/cdn-cgi/access/sso/saml/b3f58a2b414e0b51d45c8c2af26fccca0e27c63763c426fa52f98dcf0b3b3bfd"),
			})),
			SameSiteCookieAttribute: cloudflare.F("strict"),
			SelfHostedDomains:       cloudflare.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
			ServiceAuth401Redirect:  cloudflare.F(true),
			SessionDuration:         cloudflare.F("24h"),
			SkipInterstitial:        cloudflare.F(true),
			Tags:                    cloudflare.F([]string{"engineers", "engineers", "engineers"}),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		zero_trust.AccessApplicationDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		zero_trust.AccessApplicationGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		zero_trust.AccessApplicationRevokeTokensParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
