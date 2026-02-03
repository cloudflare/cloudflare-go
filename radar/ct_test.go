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

func TestCTSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.CT.Summary(
		context.TODO(),
		radar.CTSummaryParamsDimensionCA,
		radar.CTSummaryParams{
			CA:                 cloudflare.F([]string{"string"}),
			CAOwner:            cloudflare.F([]string{"string"}),
			DateEnd:            cloudflare.F([]time.Time{time.Now()}),
			DateRange:          cloudflare.F([]string{"7d"}),
			DateStart:          cloudflare.F([]time.Time{time.Now()}),
			Duration:           cloudflare.F([]radar.CTSummaryParamsDuration{radar.CTSummaryParamsDurationLte3D}),
			EntryType:          cloudflare.F([]radar.CTSummaryParamsEntryType{radar.CTSummaryParamsEntryTypePrecertificate}),
			ExpirationStatus:   cloudflare.F([]radar.CTSummaryParamsExpirationStatus{radar.CTSummaryParamsExpirationStatusExpired}),
			Format:             cloudflare.F(radar.CTSummaryParamsFormatJson),
			HasIPs:             cloudflare.F([]bool{true}),
			HasWildcards:       cloudflare.F([]bool{true}),
			LimitPerGroup:      cloudflare.F(int64(10)),
			Log:                cloudflare.F([]string{"string"}),
			LogAPI:             cloudflare.F([]radar.CTSummaryParamsLogAPI{radar.CTSummaryParamsLogAPIRfc6962}),
			LogOperator:        cloudflare.F([]string{"string"}),
			Name:               cloudflare.F([]string{"main_series"}),
			Normalization:      cloudflare.F(radar.CTSummaryParamsNormalizationRawValues),
			PublicKeyAlgorithm: cloudflare.F([]radar.CTSummaryParamsPublicKeyAlgorithm{radar.CTSummaryParamsPublicKeyAlgorithmDsa}),
			SignatureAlgorithm: cloudflare.F([]radar.CTSummaryParamsSignatureAlgorithm{radar.CTSummaryParamsSignatureAlgorithmDsaSha1}),
			TLD:                cloudflare.F([]string{"com"}),
			UniqueEntries:      cloudflare.F([]radar.CTSummaryParamsUniqueEntry{radar.CTSummaryParamsUniqueEntryTrue}),
			ValidationLevel:    cloudflare.F([]radar.CTSummaryParamsValidationLevel{radar.CTSummaryParamsValidationLevelDomain}),
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

func TestCTTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.CT.Timeseries(context.TODO(), radar.CTTimeseriesParams{
		AggInterval:        cloudflare.F(radar.CTTimeseriesParamsAggInterval1h),
		CA:                 cloudflare.F([]string{"string"}),
		CAOwner:            cloudflare.F([]string{"string"}),
		DateEnd:            cloudflare.F([]time.Time{time.Now()}),
		DateRange:          cloudflare.F([]string{"7d"}),
		DateStart:          cloudflare.F([]time.Time{time.Now()}),
		Duration:           cloudflare.F([]radar.CTTimeseriesParamsDuration{radar.CTTimeseriesParamsDurationLte3D}),
		EntryType:          cloudflare.F([]radar.CTTimeseriesParamsEntryType{radar.CTTimeseriesParamsEntryTypePrecertificate}),
		ExpirationStatus:   cloudflare.F([]radar.CTTimeseriesParamsExpirationStatus{radar.CTTimeseriesParamsExpirationStatusExpired}),
		Format:             cloudflare.F(radar.CTTimeseriesParamsFormatJson),
		HasIPs:             cloudflare.F([]bool{true}),
		HasWildcards:       cloudflare.F([]bool{true}),
		Log:                cloudflare.F([]string{"string"}),
		LogAPI:             cloudflare.F([]radar.CTTimeseriesParamsLogAPI{radar.CTTimeseriesParamsLogAPIRfc6962}),
		LogOperator:        cloudflare.F([]string{"string"}),
		Name:               cloudflare.F([]string{"main_series"}),
		PublicKeyAlgorithm: cloudflare.F([]radar.CTTimeseriesParamsPublicKeyAlgorithm{radar.CTTimeseriesParamsPublicKeyAlgorithmDsa}),
		SignatureAlgorithm: cloudflare.F([]radar.CTTimeseriesParamsSignatureAlgorithm{radar.CTTimeseriesParamsSignatureAlgorithmDsaSha1}),
		TLD:                cloudflare.F([]string{"com"}),
		UniqueEntries:      cloudflare.F([]radar.CTTimeseriesParamsUniqueEntry{radar.CTTimeseriesParamsUniqueEntryTrue}),
		ValidationLevel:    cloudflare.F([]radar.CTTimeseriesParamsValidationLevel{radar.CTTimeseriesParamsValidationLevelDomain}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCTTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.CT.TimeseriesGroups(
		context.TODO(),
		radar.CTTimeseriesGroupsParamsDimensionCA,
		radar.CTTimeseriesGroupsParams{
			AggInterval:        cloudflare.F(radar.CTTimeseriesGroupsParamsAggInterval1h),
			CA:                 cloudflare.F([]string{"string"}),
			CAOwner:            cloudflare.F([]string{"string"}),
			DateEnd:            cloudflare.F([]time.Time{time.Now()}),
			DateRange:          cloudflare.F([]string{"7d"}),
			DateStart:          cloudflare.F([]time.Time{time.Now()}),
			Duration:           cloudflare.F([]radar.CTTimeseriesGroupsParamsDuration{radar.CTTimeseriesGroupsParamsDurationLte3D}),
			EntryType:          cloudflare.F([]radar.CTTimeseriesGroupsParamsEntryType{radar.CTTimeseriesGroupsParamsEntryTypePrecertificate}),
			ExpirationStatus:   cloudflare.F([]radar.CTTimeseriesGroupsParamsExpirationStatus{radar.CTTimeseriesGroupsParamsExpirationStatusExpired}),
			Format:             cloudflare.F(radar.CTTimeseriesGroupsParamsFormatJson),
			HasIPs:             cloudflare.F([]bool{true}),
			HasWildcards:       cloudflare.F([]bool{true}),
			LimitPerGroup:      cloudflare.F(int64(10)),
			Log:                cloudflare.F([]string{"string"}),
			LogAPI:             cloudflare.F([]radar.CTTimeseriesGroupsParamsLogAPI{radar.CTTimeseriesGroupsParamsLogAPIRfc6962}),
			LogOperator:        cloudflare.F([]string{"string"}),
			Name:               cloudflare.F([]string{"main_series"}),
			Normalization:      cloudflare.F(radar.CTTimeseriesGroupsParamsNormalizationRawValues),
			PublicKeyAlgorithm: cloudflare.F([]radar.CTTimeseriesGroupsParamsPublicKeyAlgorithm{radar.CTTimeseriesGroupsParamsPublicKeyAlgorithmDsa}),
			SignatureAlgorithm: cloudflare.F([]radar.CTTimeseriesGroupsParamsSignatureAlgorithm{radar.CTTimeseriesGroupsParamsSignatureAlgorithmDsaSha1}),
			TLD:                cloudflare.F([]string{"com"}),
			UniqueEntries:      cloudflare.F([]radar.CTTimeseriesGroupsParamsUniqueEntry{radar.CTTimeseriesGroupsParamsUniqueEntryTrue}),
			ValidationLevel:    cloudflare.F([]radar.CTTimeseriesGroupsParamsValidationLevel{radar.CTTimeseriesGroupsParamsValidationLevelDomain}),
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
