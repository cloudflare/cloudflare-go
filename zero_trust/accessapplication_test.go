// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.New(context.TODO(), zero_trust.AccessApplicationNewParams{
		Body: zero_trust.AccessApplicationNewParamsBodyAccessEndUserProps{
			OAuthConfiguration: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsOAuthConfiguration{
				DynamicClientRegistration: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsOAuthConfigurationDynamicClientRegistration{
					AllowAnyOnLocalhost: cloudflare.F(true),
					AllowAnyOnLoopback:  cloudflare.F(true),
					AllowedURIs:         cloudflare.F([]string{"https://example.com/callback", "com.example.app:/oauth/callback"}),
					Enabled:             cloudflare.F(true),
				}),
				Enabled: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsOAuthConfigurationEnabledTrue),
				Grant: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsOAuthConfigurationGrant{
					AccessTokenLifetime: cloudflare.F("5m"),
					SessionDuration:     cloudflare.F("24h"),
				}),
			}),
			Type:            cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsTypeSelfHosted),
			UserPopulations: cloudflare.F([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}),
			AllowedIdPs:     cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252"}),
			Destinations: cloudflare.F([]zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsDestinationUnion{zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsDestinationsAccessEndUserPublicDestination{
				URI: cloudflare.F("uri"),
				Overrides: cloudflare.F([]zero_trust.DestinationOverrideParam{{
					Behavior:    cloudflare.F(zero_trust.DestinationOverrideBehaviorPublic),
					PathPattern: cloudflare.F("/health/*"),
				}}),
				Type: cloudflare.F(zero_trust.AccessApplicationNewParamsBodyAccessEndUserPropsDestinationsAccessEndUserPublicDestinationTypePublic),
			}}),
			Domain:            cloudflare.F("test.example.com/admin"),
			Name:              cloudflare.F("Admin Site"),
			SelfHostedDomains: cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationUpdateParams{
			Body: zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserProps{
				OAuthConfiguration: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsOAuthConfiguration{
					DynamicClientRegistration: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsOAuthConfigurationDynamicClientRegistration{
						AllowAnyOnLocalhost: cloudflare.F(true),
						AllowAnyOnLoopback:  cloudflare.F(true),
						AllowedURIs:         cloudflare.F([]string{"https://example.com/callback", "com.example.app:/oauth/callback"}),
						Enabled:             cloudflare.F(true),
					}),
					Enabled: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsOAuthConfigurationEnabledTrue),
					Grant: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsOAuthConfigurationGrant{
						AccessTokenLifetime: cloudflare.F("5m"),
						SessionDuration:     cloudflare.F("24h"),
					}),
				}),
				Type:            cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsTypeSelfHosted),
				UserPopulations: cloudflare.F([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}),
				AllowedIdPs:     cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252"}),
				Destinations: cloudflare.F([]zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsDestinationUnion{zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsDestinationsAccessEndUserPublicDestination{
					URI: cloudflare.F("uri"),
					Overrides: cloudflare.F([]zero_trust.DestinationOverrideParam{{
						Behavior:    cloudflare.F(zero_trust.DestinationOverrideBehaviorPublic),
						PathPattern: cloudflare.F("/health/*"),
					}}),
					Type: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodyAccessEndUserPropsDestinationsAccessEndUserPublicDestinationTypePublic),
				}}),
				Domain:            cloudflare.F("test.example.com/admin"),
				Name:              cloudflare.F("Admin Site"),
				SelfHostedDomains: cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Applications.List(context.TODO(), zero_trust.AccessApplicationListParams{
		AccountID:        cloudflare.F("account_id"),
		AUD:              cloudflare.F("aud"),
		Domain:           cloudflare.F("domain"),
		Exact:            cloudflare.F(true),
		Name:             cloudflare.F("name"),
		Page:             cloudflare.F(int64(0)),
		PerPage:          cloudflare.F(int64(1000)),
		Search:           cloudflare.F("search"),
		TargetAttributes: cloudflare.F("target_attributes"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
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
