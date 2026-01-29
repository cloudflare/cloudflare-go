// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/token_validation"
)

func TestConfigurationCredentialUpdate(t *testing.T) {
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
	_, err := client.TokenValidation.Configuration.Credentials.Update(
		context.TODO(),
		"4a7ee8d3-dd63-4ceb-9d5f-c27831854ce7",
		token_validation.ConfigurationCredentialUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Keys: cloudflare.F([]token_validation.ConfigurationCredentialUpdateParamsKeyUnion{token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256{
				Alg: cloudflare.F(token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256AlgEs256),
				Crv: cloudflare.F(token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256CrvP256),
				Kid: cloudflare.F("38013f13-c266-4eec-a72a-92ec92779f21"),
				Kty: cloudflare.F(token_validation.ConfigurationCredentialUpdateParamsKeysAPIShieldCredentialsJWTKeyEcEs256KtyEc),
				X:   cloudflare.F("KN53JRwN3wCjm2o39bvZUX2VdrsHzS8pxOAGjm8m7EQ"),
				Y:   cloudflare.F("lnkkzIxaveggz-HFhcMWW15nxvOj0Z_uQsXbpK0GFcY"),
			}}),
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
