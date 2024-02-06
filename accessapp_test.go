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

func TestAccessAppGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Get(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessAppUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Update(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
		cloudflare.AccessAppUpdateParamsVariant0{
			AllowAuthenticateViaWarp: cloudflare.F(true),
			AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F(true),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CorsHeaders: cloudflare.F(cloudflare.AccessAppUpdateParamsVariant0CorsHeaders{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				AllowedMethods:   cloudflare.F([]cloudflare.AccessAppUpdateParamsVariant0CorsHeadersAllowedMethod{cloudflare.AccessAppUpdateParamsVariant0CorsHeadersAllowedMethodGet}),
				AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
				MaxAge:           cloudflare.F(-1.000000),
			}),
			CustomDenyMessage:        cloudflare.F("string"),
			CustomDenyURL:            cloudflare.F("string"),
			CustomNonIdentityDenyURL: cloudflare.F("string"),
			CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			Domain:                   cloudflare.F("test.example.com/admin"),
			EnableBindingCookie:      cloudflare.F(true),
			HTTPOnlyCookieAttribute:  cloudflare.F(true),
			LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                     cloudflare.F("Admin Site"),
			PathCookieAttribute:      cloudflare.F(true),
			SameSiteCookieAttribute:  cloudflare.F("strict"),
			SelfHostedDomains:        cloudflare.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
			ServiceAuth401Redirect:   cloudflare.F(true),
			SessionDuration:          cloudflare.F("24h"),
			SkipInterstitial:         cloudflare.F(true),
			Tags:                     cloudflare.F([]string{"engineers", "engineers", "engineers"}),
			Type:                     cloudflare.F("self_hosted"),
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

func TestAccessAppDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.Delete(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		shared.UnionString("023e105f4ecef8ad9ca31a8372d0c353"),
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessAppAccessApplicationsAddAnApplicationWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.AccessApplicationsAddAnApplication(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccessAppAccessApplicationsAddAnApplicationParamsVariant0{
			AllowAuthenticateViaWarp: cloudflare.F(true),
			AllowedIdps:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			AppLauncherVisible:       cloudflare.F(true),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CorsHeaders: cloudflare.F(cloudflare.AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeaders{
				AllowAllHeaders:  cloudflare.F(true),
				AllowAllMethods:  cloudflare.F(true),
				AllowAllOrigins:  cloudflare.F(true),
				AllowCredentials: cloudflare.F(true),
				AllowedHeaders:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				AllowedMethods:   cloudflare.F([]cloudflare.AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethod{cloudflare.AccessAppAccessApplicationsAddAnApplicationParamsVariant0CorsHeadersAllowedMethodGet}),
				AllowedOrigins:   cloudflare.F([]interface{}{"https://example.com"}),
				MaxAge:           cloudflare.F(-1.000000),
			}),
			CustomDenyMessage:        cloudflare.F("string"),
			CustomDenyURL:            cloudflare.F("string"),
			CustomNonIdentityDenyURL: cloudflare.F("string"),
			CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
			Domain:                   cloudflare.F("test.example.com/admin"),
			EnableBindingCookie:      cloudflare.F(true),
			HTTPOnlyCookieAttribute:  cloudflare.F(true),
			LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:                     cloudflare.F("Admin Site"),
			PathCookieAttribute:      cloudflare.F(true),
			SameSiteCookieAttribute:  cloudflare.F("strict"),
			SelfHostedDomains:        cloudflare.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
			ServiceAuth401Redirect:   cloudflare.F(true),
			SessionDuration:          cloudflare.F("24h"),
			SkipInterstitial:         cloudflare.F(true),
			Tags:                     cloudflare.F([]string{"engineers", "engineers", "engineers"}),
			Type:                     cloudflare.F("self_hosted"),
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

func TestAccessAppAccessApplicationsListAccessApplications(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Access.Apps.AccessApplicationsListAccessApplications(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
