// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/workers_for_platforms"
)

func TestDispatchNamespaceScriptAssetUploadNew(t *testing.T) {
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
	_, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.AssetUpload.New(
		context.TODO(),
		"my-dispatch-namespace",
		"this-is_my_script-01",
		workers_for_platforms.DispatchNamespaceScriptAssetUploadNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Manifest: cloudflare.F(map[string]workers_for_platforms.DispatchNamespaceScriptAssetUploadNewParamsManifest{
				"foo": {
					Hash: cloudflare.F("hash"),
					Size: cloudflare.F(int64(0)),
				},
			}),
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
