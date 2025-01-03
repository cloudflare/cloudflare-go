// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
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
		Body: zero_trust.AccessApplicationNewParamsBodyBookmarkApplication{
			Domain:             cloudflare.F("test.example.com/admin"),
			Type:               cloudflare.F("self_hosted"),
			AppLauncherVisible: cloudflare.F(true),
			LogoURL:            cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
			Name:               cloudflare.F("Admin Site"),
			SCIMConfig: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfig{
				IdPUID:    cloudflare.F("idp_uid"),
				RemoteURI: cloudflare.F("remote_uri"),
				Authentication: cloudflare.F[zero_trust.AccessApplicationNewParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
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
			Tags: cloudflare.F([]string{"engineers"}),
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
			Body: zero_trust.AccessApplicationUpdateParamsBodyBookmarkApplication{
				Domain:             cloudflare.F("test.example.com/admin"),
				Type:               cloudflare.F("self_hosted"),
				AppLauncherVisible: cloudflare.F(true),
				LogoURL:            cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				Name:               cloudflare.F("Admin Site"),
				SCIMConfig: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfig{
					IdPUID:    cloudflare.F("idp_uid"),
					RemoteURI: cloudflare.F("remote_uri"),
					Authentication: cloudflare.F[zero_trust.AccessApplicationUpdateParamsBodyBookmarkApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
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
				Tags: cloudflare.F([]string{"engineers"}),
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
