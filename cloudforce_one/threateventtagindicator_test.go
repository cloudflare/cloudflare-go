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

func TestThreatEventTagIndicatorListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Indicators.List(
		context.TODO(),
		"tag_uuid",
		cloudforce_one.ThreatEventTagIndicatorListParams{
			AccountID:     cloudflare.F("account_id"),
			DatasetIDs:    cloudflare.F([]string{"string"}),
			IndicatorType: cloudflare.F("indicatorType"),
			Page:          cloudflare.F(0.000000),
			PageSize:      cloudflare.F(0.000000),
			RelatedEvent:  cloudflare.F([]string{"string"}),
			Search: cloudflare.F([]cloudforce_one.ThreatEventTagIndicatorListParamsSearch{{
				Field: cloudflare.F(cloudforce_one.ThreatEventTagIndicatorListParamsSearchFieldValue),
				Op:    cloudflare.F(cloudforce_one.ThreatEventTagIndicatorListParamsSearchOpContains),
				Value: cloudflare.F[cloudforce_one.ThreatEventTagIndicatorListParamsSearchValueUnion](shared.UnionString("malicious")),
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
