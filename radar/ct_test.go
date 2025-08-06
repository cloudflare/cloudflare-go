// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/radar"
)

func TestCtSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ct.Summary(
		context.TODO(),
		radar.CtSummaryParamsDimensionCA,
		radar.CtSummaryParams{
			CA:                 cloudflare.F([]string{"string"}),
			CAOwner:            cloudflare.F([]string{"string"}),
			DateEnd:            cloudflare.F([]time.Time{time.Now()}),
			DateRange:          cloudflare.F([]string{"7d"}),
			DateStart:          cloudflare.F([]time.Time{time.Now()}),
			Duration:           cloudflare.F([]radar.CtSummaryParamsDuration{radar.CtSummaryParamsDurationLte3D}),
			EntryType:          cloudflare.F([]radar.CtSummaryParamsEntryType{radar.CtSummaryParamsEntryTypePrecertificate}),
			ExpirationStatus:   cloudflare.F([]radar.CtSummaryParamsExpirationStatus{radar.CtSummaryParamsExpirationStatusExpired}),
			Format:             cloudflare.F(radar.CtSummaryParamsFormatJson),
			HasIPs:             cloudflare.F([]bool{true}),
			HasWildcards:       cloudflare.F([]bool{true}),
			LimitPerGroup:      cloudflare.F(int64(10)),
			Log:                cloudflare.F([]string{"string"}),
			LogAPI:             cloudflare.F([]radar.CtSummaryParamsLogAPI{radar.CtSummaryParamsLogAPIRfc6962}),
			LogOperator:        cloudflare.F([]string{"string"}),
			Name:               cloudflare.F([]string{"main_series"}),
			Normalization:      cloudflare.F(radar.CtSummaryParamsNormalizationRawValues),
			PublicKeyAlgorithm: cloudflare.F([]radar.CtSummaryParamsPublicKeyAlgorithm{radar.CtSummaryParamsPublicKeyAlgorithmDsa}),
			SignatureAlgorithm: cloudflare.F([]radar.CtSummaryParamsSignatureAlgorithm{radar.CtSummaryParamsSignatureAlgorithmDsaSha1}),
			Tld:                cloudflare.F([]string{"string"}),
			UniqueEntries:      cloudflare.F([]radar.CtSummaryParamsUniqueEntry{radar.CtSummaryParamsUniqueEntryTrue}),
			ValidationLevel:    cloudflare.F([]radar.CtSummaryParamsValidationLevel{radar.CtSummaryParamsValidationLevelDomain}),
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

func TestCtTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ct.Timeseries(context.TODO(), radar.CtTimeseriesParams{
		AggInterval:        cloudflare.F(radar.CtTimeseriesParamsAggInterval1h),
		CA:                 cloudflare.F([]string{"string"}),
		CAOwner:            cloudflare.F([]string{"string"}),
		DateEnd:            cloudflare.F([]time.Time{time.Now()}),
		DateRange:          cloudflare.F([]string{"7d"}),
		DateStart:          cloudflare.F([]time.Time{time.Now()}),
		Duration:           cloudflare.F([]radar.CtTimeseriesParamsDuration{radar.CtTimeseriesParamsDurationLte3D}),
		EntryType:          cloudflare.F([]radar.CtTimeseriesParamsEntryType{radar.CtTimeseriesParamsEntryTypePrecertificate}),
		ExpirationStatus:   cloudflare.F([]radar.CtTimeseriesParamsExpirationStatus{radar.CtTimeseriesParamsExpirationStatusExpired}),
		Format:             cloudflare.F(radar.CtTimeseriesParamsFormatJson),
		HasIPs:             cloudflare.F([]bool{true}),
		HasWildcards:       cloudflare.F([]bool{true}),
		Log:                cloudflare.F([]string{"string"}),
		LogAPI:             cloudflare.F([]radar.CtTimeseriesParamsLogAPI{radar.CtTimeseriesParamsLogAPIRfc6962}),
		LogOperator:        cloudflare.F([]string{"string"}),
		Name:               cloudflare.F([]string{"main_series"}),
		PublicKeyAlgorithm: cloudflare.F([]radar.CtTimeseriesParamsPublicKeyAlgorithm{radar.CtTimeseriesParamsPublicKeyAlgorithmDsa}),
		SignatureAlgorithm: cloudflare.F([]radar.CtTimeseriesParamsSignatureAlgorithm{radar.CtTimeseriesParamsSignatureAlgorithmDsaSha1}),
		Tld:                cloudflare.F([]string{"string"}),
		UniqueEntries:      cloudflare.F([]radar.CtTimeseriesParamsUniqueEntry{radar.CtTimeseriesParamsUniqueEntryTrue}),
		ValidationLevel:    cloudflare.F([]radar.CtTimeseriesParamsValidationLevel{radar.CtTimeseriesParamsValidationLevelDomain}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCtTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ct.TimeseriesGroups(
		context.TODO(),
		radar.CtTimeseriesGroupsParamsDimensionCA,
		radar.CtTimeseriesGroupsParams{
			AggInterval:        cloudflare.F(radar.CtTimeseriesGroupsParamsAggInterval1h),
			CA:                 cloudflare.F([]string{"string"}),
			CAOwner:            cloudflare.F([]string{"string"}),
			DateEnd:            cloudflare.F([]time.Time{time.Now()}),
			DateRange:          cloudflare.F([]string{"7d"}),
			DateStart:          cloudflare.F([]time.Time{time.Now()}),
			Duration:           cloudflare.F([]radar.CtTimeseriesGroupsParamsDuration{radar.CtTimeseriesGroupsParamsDurationLte3D}),
			EntryType:          cloudflare.F([]radar.CtTimeseriesGroupsParamsEntryType{radar.CtTimeseriesGroupsParamsEntryTypePrecertificate}),
			ExpirationStatus:   cloudflare.F([]radar.CtTimeseriesGroupsParamsExpirationStatus{radar.CtTimeseriesGroupsParamsExpirationStatusExpired}),
			Format:             cloudflare.F(radar.CtTimeseriesGroupsParamsFormatJson),
			HasIPs:             cloudflare.F([]bool{true}),
			HasWildcards:       cloudflare.F([]bool{true}),
			LimitPerGroup:      cloudflare.F(int64(10)),
			Log:                cloudflare.F([]string{"string"}),
			LogAPI:             cloudflare.F([]radar.CtTimeseriesGroupsParamsLogAPI{radar.CtTimeseriesGroupsParamsLogAPIRfc6962}),
			LogOperator:        cloudflare.F([]string{"string"}),
			Name:               cloudflare.F([]string{"main_series"}),
			Normalization:      cloudflare.F(radar.CtTimeseriesGroupsParamsNormalizationRawValues),
			PublicKeyAlgorithm: cloudflare.F([]radar.CtTimeseriesGroupsParamsPublicKeyAlgorithm{radar.CtTimeseriesGroupsParamsPublicKeyAlgorithmDsa}),
			SignatureAlgorithm: cloudflare.F([]radar.CtTimeseriesGroupsParamsSignatureAlgorithm{radar.CtTimeseriesGroupsParamsSignatureAlgorithmDsaSha1}),
			Tld:                cloudflare.F([]string{"string"}),
			UniqueEntries:      cloudflare.F([]radar.CtTimeseriesGroupsParamsUniqueEntry{radar.CtTimeseriesGroupsParamsUniqueEntryTrue}),
			ValidationLevel:    cloudflare.F([]radar.CtTimeseriesGroupsParamsValidationLevel{radar.CtTimeseriesGroupsParamsValidationLevelDomain}),
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
