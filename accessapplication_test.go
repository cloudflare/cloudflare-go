// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.New(context.TODO(), cloudflare.AccessApplicationNewParams{
		AccountID:                cloudflare.F("string"),
		ZoneID:                   cloudflare.F("string"),
		AllowAuthenticateViaWARP: cloudflare.F(true),
		AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
		AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
		AutoRedirectToIdentity:   cloudflare.F(true),
		CorsHeaders: cloudflare.F(cloudflare.AccessApplicationNewParamsCorsHeaders{
			AllowAllHeaders:  cloudflare.F(true),
			AllowAllMethods:  cloudflare.F(true),
			AllowAllOrigins:  cloudflare.F(true),
			AllowCredentials: cloudflare.F(true),
			AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
			AllowedMethods:   cloudflare.F([]cloudflare.AccessApplicationNewParamsCorsHeadersAllowedMethod{cloudflare.AccessApplicationNewParamsCorsHeadersAllowedMethodGet}),
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
		SaasApp: cloudflare.F[cloudflare.AccessApplicationNewParamsSaasApp](cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasApp(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasApp{
			AuthType:           cloudflare.F(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
			ConsumerServiceURL: cloudflare.F("https://example.com"),
			CustomAttributes: cloudflare.F(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributes{
				Name:       cloudflare.F("family_name"),
				NameFormat: cloudflare.F(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
				Source: cloudflare.F(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
					Name: cloudflare.F("last_name"),
				}),
			}),
			DefaultRelayState:      cloudflare.F("https://example.com"),
			IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
			NameIDFormat:           cloudflare.F(cloudflare.AccessApplicationNewParamsSaasAppAccessSamlSaasAppNameIDFormatID),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.Update(
		context.TODO(),
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		cloudflare.AccessApplicationUpdateParams{
			AccountID:                cloudflare.F("string"),
			ZoneID:                   cloudflare.F("string"),
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F[any](map[string]interface{}{}),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CorsHeaders: cloudflare.F(cloudflare.AccessApplicationUpdateParamsCorsHeaders{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				AllowedMethods:   cloudflare.F([]cloudflare.AccessApplicationUpdateParamsCorsHeadersAllowedMethod{cloudflare.AccessApplicationUpdateParamsCorsHeadersAllowedMethodGet}),
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
			SaasApp: cloudflare.F[cloudflare.AccessApplicationUpdateParamsSaasApp](cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasApp{
				AuthType:           cloudflare.F(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppAuthTypeSaml),
				ConsumerServiceURL: cloudflare.F("https://example.com"),
				CustomAttributes: cloudflare.F(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributes{
					Name:       cloudflare.F("family_name"),
					NameFormat: cloudflare.F(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesNameFormatUrnOasisNamesTcSaml2_0AttrnameFormatBasic),
					Source: cloudflare.F(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppCustomAttributesSource{
						Name: cloudflare.F("last_name"),
					}),
				}),
				DefaultRelayState:      cloudflare.F("https://example.com"),
				IdpEntityID:            cloudflare.F("https://example.cloudflareaccess.com"),
				NameIDFormat:           cloudflare.F(cloudflare.AccessApplicationUpdateParamsSaasAppAccessSamlSaasAppNameIDFormatID),
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

func TestAccessApplicationList(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.List(context.TODO(), cloudflare.AccessApplicationListParams{
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

func TestAccessApplicationDelete(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.Delete(
		context.TODO(),
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		cloudflare.AccessApplicationDeleteParams{
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

func TestAccessApplicationGet(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.Get(
		context.TODO(),
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		cloudflare.AccessApplicationGetParams{
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

func TestAccessApplicationRevokeTokens(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Access.Applications.RevokeTokens(
		context.TODO(),
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		cloudflare.AccessApplicationRevokeTokensParams{
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
