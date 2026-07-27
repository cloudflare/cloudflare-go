// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

func TestThreatEventTagIndicatorByDatasetListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Indicators.ByDataset.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"tag_uuid",
		cloudforce_one.ThreatEventTagIndicatorByDatasetListParams{
			AccountID:     cloudflare.F("account_id"),
			IndicatorType: cloudflare.F("indicatorType"),
			Page:          cloudflare.F(0.000000),
			PageSize:      cloudflare.F(0.000000),
			RelatedEvent:  cloudflare.F([]string{"string"}),
			Search: cloudflare.F([]cloudforce_one.ThreatEventTagIndicatorByDatasetListParamsSearch{{
				Field: cloudflare.F(cloudforce_one.ThreatEventTagIndicatorByDatasetListParamsSearchFieldValue),
				Op:    cloudflare.F(cloudforce_one.ThreatEventTagIndicatorByDatasetListParamsSearchOpContains),
				Value: cloudflare.F[cloudforce_one.ThreatEventTagIndicatorByDatasetListParamsSearchValueUnion](shared.UnionString("malicious")),
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
