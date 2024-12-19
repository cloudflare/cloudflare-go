// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/r2"
)

func TestTemporaryCredentialNewWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.TemporaryCredentials.New(context.TODO(), r2.TemporaryCredentialNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		TemporaryCredential: r2.TemporaryCredentialParam{
			Bucket:            cloudflare.F("example-bucket"),
			ParentAccessKeyID: cloudflare.F("example-access-key-id"),
			Permission:        cloudflare.F(r2.TemporaryCredentialPermissionAdminReadWrite),
			TTLSeconds:        cloudflare.F(3600.000000),
			Objects:           cloudflare.F([]string{"example-object"}),
			Prefixes:          cloudflare.F([]string{"example-prefix/"}),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
