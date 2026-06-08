// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/iam"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestOAuthClientNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.New(context.TODO(), iam.OAuthClientNewParams{
		AccountID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ClientName:              cloudflare.F("My OAuth App"),
		GrantTypes:              cloudflare.F([]iam.OAuthClientNewParamsGrantType{iam.OAuthClientNewParamsGrantTypeAuthorizationCode, iam.OAuthClientNewParamsGrantTypeRefreshToken}),
		RedirectURIs:            cloudflare.F([]string{"https://example.com/callback"}),
		ResponseTypes:           cloudflare.F([]iam.OAuthClientNewParamsResponseType{iam.OAuthClientNewParamsResponseTypeCode}),
		Scopes:                  cloudflare.F([]string{"account.read"}),
		TokenEndpointAuthMethod: cloudflare.F(iam.OAuthClientNewParamsTokenEndpointAuthMethodClientSecretPost),
		AllowedCORSOrigins:      cloudflare.F([]string{"https://example.com"}),
		ClientURI:               cloudflare.F("https://example.com"),
		LogoURI:                 cloudflare.F("https://example.com/logo.png"),
		PolicyURI:               cloudflare.F("https://example.com/privacy"),
		PostLogoutRedirectURIs:  cloudflare.F([]string{"https://example.com/logout"}),
		TosURI:                  cloudflare.F("https://example.com/tos"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthClientUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.Update(
		context.TODO(),
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4",
		iam.OAuthClientUpdateParams{
			AccountID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedCORSOrigins:      cloudflare.F([]string{"https://example.com"}),
			ClientName:              cloudflare.F("My OAuth App"),
			ClientURI:               cloudflare.F("https://example.com"),
			GrantTypes:              cloudflare.F([]iam.OAuthClientUpdateParamsGrantType{iam.OAuthClientUpdateParamsGrantTypeAuthorizationCode, iam.OAuthClientUpdateParamsGrantTypeRefreshToken}),
			LogoURI:                 cloudflare.F("https://example.com/logo.png"),
			PolicyURI:               cloudflare.F("https://example.com/privacy"),
			PostLogoutRedirectURIs:  cloudflare.F([]string{"https://example.com/logout"}),
			RedirectURIs:            cloudflare.F([]string{"https://example.com/callback"}),
			ResponseTypes:           cloudflare.F([]iam.OAuthClientUpdateParamsResponseType{iam.OAuthClientUpdateParamsResponseTypeCode}),
			Scopes:                  cloudflare.F([]string{"account.read"}),
			TokenEndpointAuthMethod: cloudflare.F(iam.OAuthClientUpdateParamsTokenEndpointAuthMethodClientSecretPost),
			TosURI:                  cloudflare.F("https://example.com/tos"),
			Visibility:              cloudflare.F(iam.OAuthClientUpdateParamsVisibilityPublic),
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

func TestOAuthClientList(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.List(context.TODO(), iam.OAuthClientListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthClientDelete(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.Delete(
		context.TODO(),
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4",
		iam.OAuthClientDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestOAuthClientDeleteRotatedSecret(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.DeleteRotatedSecret(
		context.TODO(),
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4",
		iam.OAuthClientDeleteRotatedSecretParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestOAuthClientGet(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.Get(
		context.TODO(),
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4",
		iam.OAuthClientGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestOAuthClientRotateSecret(t *testing.T) {
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
	_, err := client.IAM.OAuthClients.RotateSecret(
		context.TODO(),
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4",
		iam.OAuthClientRotateSecretParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
