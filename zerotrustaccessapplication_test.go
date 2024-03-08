// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustAccessApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.New(context.TODO(), cloudflare.ZeroTrustAccessApplicationNewParams{
		AccountID:                cloudflare.F("string"),
		ZoneID:                   cloudflare.F("string"),
		AllowAuthenticateViaWARP: cloudflare.F(true),
		AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
		AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
		AutoRedirectToIdentity:   cloudflare.F(true),
		CorsHeaders: cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsCorsHeaders{
			AllowAllHeaders:  cloudflare.F(true),
			AllowAllMethods:  cloudflare.F(true),
			AllowAllOrigins:  cloudflare.F(true),
			AllowCredentials: cloudflare.F(true),
			AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
			AllowedMethods:   cloudflare.F([]cloudflare.ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethod{cloudflare.ZeroTrustAccessApplicationNewParamsCorsHeadersAllowedMethodGet}),
			AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
			MaxAge:           cloudflare.F(-1.000000),
		}),
		CustomDenyMessage:        cloudflare.F("string"),
		CustomDenyURL:            cloudflare.F("string"),
		CustomNonIdentityDenyURL: cloudflare.F("string"),
		CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
		Domain:                   cloudflare.F[any]("https://mybookmark.com"),
		EnableBindingCookie:      cloudflare.F(true),
		HTTPOnlyCookieAttribute:  cloudflare.F(true),
		LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
		Name:                     cloudflare.F("Admin Site"),
		PathCookieAttribute:      cloudflare.F(true),
		SaasApp: cloudflare.F[cloudflare.ZeroTrustAccessApplicationNewParamsSaasApp](cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasApp{
			AuthType:           cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
			ConsumerServiceURL: cloudflare.F("https://example.com"),
			CustomAttributes: cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes{
				Name:       cloudflare.F("family_name"),
				NameFormat: cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
				Source: cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
					Name: cloudflare.F("last_name"),
				}),
			}),
			DefaultRelayState:      cloudflare.F("https://example.com"),
			IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
			NameIDFormat:           cloudflare.F(cloudflare.ZeroTrustAccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatID),
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
		Type:                    cloudflare.F("bookmark"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustAccessApplicationUpdateWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationUpdateParams{
			AccountID:                cloudflare.F("string"),
			ZoneID:                   cloudflare.F("string"),
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CorsHeaders: cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsCorsHeaders{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				AllowedMethods:   cloudflare.F([]cloudflare.ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethod{cloudflare.ZeroTrustAccessApplicationUpdateParamsCorsHeadersAllowedMethodGet}),
				AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
				MaxAge:           cloudflare.F(-1.000000),
			}),
			CustomDenyMessage:        cloudflare.F("string"),
			CustomDenyURL:            cloudflare.F("string"),
			CustomNonIdentityDenyURL: cloudflare.F("string"),
			CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			Domain:                   cloudflare.F[any]("https://mybookmark.com"),
			EnableBindingCookie:      cloudflare.F(true),
			HTTPOnlyCookieAttribute:  cloudflare.F(true),
			LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                     cloudflare.F("Admin Site"),
			PathCookieAttribute:      cloudflare.F(true),
			SaasApp: cloudflare.F[cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasApp](cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasApp{
				AuthType:           cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
				ConsumerServiceURL: cloudflare.F("https://example.com"),
				CustomAttributes: cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes{
					Name:       cloudflare.F("family_name"),
					NameFormat: cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
					Source: cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
						Name: cloudflare.F("last_name"),
					}),
				}),
				DefaultRelayState:      cloudflare.F("https://example.com"),
				IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
				NameIDFormat:           cloudflare.F(cloudflare.ZeroTrustAccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatID),
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
			Type:                    cloudflare.F("bookmark"),
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

func TestZeroTrustAccessApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.List(context.TODO(), cloudflare.ZeroTrustAccessApplicationListParams{
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

func TestZeroTrustAccessApplicationDeleteWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationDeleteParams{
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

func TestZeroTrustAccessApplicationGetWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationGetParams{
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

func TestZeroTrustAccessApplicationRevokeTokensWithOptionalParams(t *testing.T) {
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
		cloudflare.ZeroTrustAccessApplicationRevokeTokensParams{
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
