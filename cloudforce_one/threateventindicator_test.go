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

func TestThreatEventIndicatorListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Indicators.List(context.TODO(), cloudforce_one.ThreatEventIndicatorListParams{
		AccountID:          cloudflare.F("account_id"),
		CreatedAfter:       cloudflare.F(time.Now()),
		CreatedBefore:      cloudflare.F(time.Now()),
		DatasetIDs:         cloudflare.F([]string{"string"}),
		Format:             cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsFormatJson),
		IncludeTags:        cloudflare.F(true),
		IncludeTotalCount:  cloudflare.F(true),
		IndicatorType:      cloudflare.F("indicatorType"),
		Name:               cloudflare.F("name"),
		Page:               cloudflare.F(0.000000),
		PageSize:           cloudflare.F(0.000000),
		RelatedEvents:      cloudflare.F([]string{"string"}),
		RelatedEventsLimit: cloudflare.F(2.000000),
		Search: cloudflare.F([]cloudforce_one.ThreatEventIndicatorListParamsSearch{{
			Field: cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsSearchFieldValue),
			Op:    cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsSearchOpContains),
			Value: cloudflare.F[cloudforce_one.ThreatEventIndicatorListParamsSearchValueUnion](shared.UnionString("malicious")),
		}}),
		Source: cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsSourceDo),
		Tags:   cloudflare.F([]string{"string"}),
		TagSearch: cloudflare.F([]cloudforce_one.ThreatEventIndicatorListParamsTagSearch{{
			Field: cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsTagSearchFieldOriginCountryISO),
			Op:    cloudflare.F(cloudforce_one.ThreatEventIndicatorListParamsTagSearchOpIn),
			Value: cloudflare.F[cloudforce_one.ThreatEventIndicatorListParamsTagSearchValueUnion](shared.UnionString("IR")),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
