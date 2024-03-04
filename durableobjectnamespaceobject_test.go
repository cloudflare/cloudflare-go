// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestDurableObjectNamespaceObjectListWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.DurableObjects.Namespaces.Objects.List(
		context.TODO(),
		"5fd1cafff895419c8bcc647fc64ab8f0",
		cloudflare.DurableObjectNamespaceObjectListParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Cursor:    cloudflare.F("AAAAANuhDN7SjacTnSVsDu3WW1Lvst6dxJGTjRY5BhxPXdf6L6uTcpd_NVtjhn11OUYRsVEykxoUwF-JQU4dn6QylZSKTOJuG0indrdn_MlHpMRtsxgXjs-RPdHYIVm3odE_uvEQ_dTQGFm8oikZMohns34DLBgrQpc"),
			Limit:     cloudflare.F(10.000000),
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
