// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/radar"
)

func TestLeakedCredentialSummaryV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.LeakedCredentials.SummaryV2(
		context.TODO(),
		radar.LeakedCredentialSummaryV2ParamsDimensionCompromised,
		radar.LeakedCredentialSummaryV2Params{
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.LeakedCredentialSummaryV2ParamsBotClass{radar.LeakedCredentialSummaryV2ParamsBotClassLikelyAutomated}),
			Compromised:   cloudflare.F([]radar.LeakedCredentialSummaryV2ParamsCompromised{radar.LeakedCredentialSummaryV2ParamsCompromisedClean}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.LeakedCredentialSummaryV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
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

func TestLeakedCredentialTimeseriesGroupsV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.LeakedCredentials.TimeseriesGroupsV2(
		context.TODO(),
		radar.LeakedCredentialTimeseriesGroupsV2ParamsDimensionCompromised,
		radar.LeakedCredentialTimeseriesGroupsV2Params{
			AggInterval:   cloudflare.F(radar.LeakedCredentialTimeseriesGroupsV2ParamsAggInterval1h),
			ASN:           cloudflare.F([]string{"string"}),
			BotClass:      cloudflare.F([]radar.LeakedCredentialTimeseriesGroupsV2ParamsBotClass{radar.LeakedCredentialTimeseriesGroupsV2ParamsBotClassLikelyAutomated}),
			CheckResult:   cloudflare.F([]radar.LeakedCredentialTimeseriesGroupsV2ParamsCheckResult{radar.LeakedCredentialTimeseriesGroupsV2ParamsCheckResultClean}),
			Compromised:   cloudflare.F([]radar.LeakedCredentialTimeseriesGroupsV2ParamsCompromised{radar.LeakedCredentialTimeseriesGroupsV2ParamsCompromisedClean}),
			Continent:     cloudflare.F([]string{"string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now()}),
			DateRange:     cloudflare.F([]string{"7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now()}),
			Format:        cloudflare.F(radar.LeakedCredentialTimeseriesGroupsV2ParamsFormatJson),
			LimitPerGroup: cloudflare.F(int64(10)),
			Location:      cloudflare.F([]string{"string"}),
			Name:          cloudflare.F([]string{"main_series"}),
			Normalization: cloudflare.F(radar.LeakedCredentialTimeseriesGroupsV2ParamsNormalizationMin0Max),
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
