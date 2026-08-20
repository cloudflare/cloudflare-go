// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

func TestThreatEventIndicatorAggregateListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Indicators.Aggregate.List(context.TODO(), cloudforce_one.ThreatEventIndicatorAggregateListParams{
		AccountID:       cloudflare.F("account_id"),
		AggregateBy:     cloudflare.F("aggregateBy"),
		CreatedAfter:    cloudflare.F[cloudforce_one.ThreatEventIndicatorAggregateListParamsCreatedAfterUnion](shared.UnionTime(time.Now())),
		CreatedBefore:   cloudflare.F[cloudforce_one.ThreatEventIndicatorAggregateListParamsCreatedBeforeUnion](shared.UnionTime(time.Now())),
		DatasetIDs:      cloudflare.F([]string{"string"}),
		EventDateAfter:  cloudflare.F("eventDateAfter"),
		EventDateBefore: cloudflare.F("eventDateBefore"),
		Limit:           cloudflare.F(1.000000),
		Measure:         cloudflare.F(cloudforce_one.ThreatEventIndicatorAggregateListParamsMeasureIndicators),
		TagUUID:         cloudflare.F("tagUuid"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
