// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestDLPDatasetVersionNew(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Datasets.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		int64(0),
		zero_trust.DLPDatasetVersionNewParams{
			AccountID: cloudflare.F("account_id"),
			Body: []zero_trust.DLPDatasetVersionNewParamsBodyUnion{zero_trust.DLPDatasetVersionNewParamsBodyExistingColumn{
				EntryID:    cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				HeaderName: cloudflare.F("header_name"),
				NumCells:   cloudflare.F(int64(0)),
			}},
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
