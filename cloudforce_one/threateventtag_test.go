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

func TestThreatEventTagNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism")
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.New(context.TODO(), cloudforce_one.ThreatEventTagNewParams{
		AccountID:               cloudflare.F("account_id"),
		Value:                   cloudflare.F("APT28"),
		ActiveDuration:          cloudflare.F("activeDuration"),
		ActorCategory:           cloudflare.F("Nation State"),
		ActorCategoryConfidence: cloudflare.F(int64(7)),
		Aliases: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsAlias{{
			Value:      cloudflare.F("Fancy Bear"),
			Confidence: cloudflare.F(int64(8)),
			TLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsAliasesTLPAmber),
		}}),
		AliasGroupNames:            cloudflare.F([]string{"string"}),
		AliasGroupNamesInternal:    cloudflare.F([]string{"string"}),
		AnalyticPriority:           cloudflare.F(0.000000),
		AttributionConfidence:      cloudflare.F("attributionConfidence"),
		AttributionConfidenceScore: cloudflare.F(int64(7)),
		AttributionOrganization:    cloudflare.F("attributionOrganization"),
		CategoryUUID:               cloudflare.F("12345678-1234-1234-1234-1234567890ab"),
		DateOfDiscovery:            cloudflare.F("2024-01-15"),
		ExternalReferenceLinks:     cloudflare.F([]string{"string"}),
		ExternalReferences: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsExternalReference{{
			URL:         cloudflare.F("https://example.com/report"),
			Description: cloudflare.F("Vendor threat report"),
		}}),
		InternalAliases: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsInternalAlias{{
			Value:      cloudflare.F("Fancy Bear"),
			Confidence: cloudflare.F(int64(8)),
			TLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsInternalAliasesTLPAmber),
		}}),
		InternalDescription:     cloudflare.F("internalDescription"),
		Motive:                  cloudflare.F("Espionage"),
		MotiveConfidence:        cloudflare.F(int64(7)),
		OpsecLevel:              cloudflare.F("opsecLevel"),
		OriginCountryConfidence: cloudflare.F(int64(7)),
		OriginCountryISO:        cloudflare.F("originCountryISO"),
		OriginCountryTLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsOriginCountryTLPAmber),
		Priority:                cloudflare.F(0.000000),
		SophisticationLevel:     cloudflare.F("sophisticationLevel"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatEventTagListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.List(context.TODO(), cloudforce_one.ThreatEventTagListParams{
		AccountID:    cloudflare.F("account_id"),
		Cache:        cloudflare.F(cloudforce_one.ThreatEventTagListParamsCacheFromGraph),
		CategoryUUID: cloudflare.F("categoryUuid"),
		Filters: cloudflare.F([]cloudforce_one.ThreatEventTagListParamsFilter{{
			Field: cloudflare.F(cloudforce_one.ThreatEventTagListParamsFiltersFieldOriginCountryISO),
			Op:    cloudflare.F(cloudforce_one.ThreatEventTagListParamsFiltersOpIn),
			Value: cloudflare.F[cloudforce_one.ThreatEventTagListParamsFiltersValueUnion](shared.UnionString("IR")),
		}}),
		Page:     cloudflare.F(0.000000),
		PageSize: cloudflare.F(0.000000),
		Search:   cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatEventTagDelete(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Delete(
		context.TODO(),
		"tag_uuid",
		cloudforce_one.ThreatEventTagDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestThreatEventTagEditWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Edit(
		context.TODO(),
		"tag_uuid",
		cloudforce_one.ThreatEventTagEditParams{
			AccountID:               cloudflare.F("account_id"),
			ActiveDuration:          cloudflare.F("activeDuration"),
			ActorCategory:           cloudflare.F("Nation State"),
			ActorCategoryConfidence: cloudflare.F(int64(7)),
			Aliases: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsAlias{{
				Value:      cloudflare.F("Fancy Bear"),
				Confidence: cloudflare.F(int64(8)),
				TLP:        cloudflare.F(cloudforce_one.ThreatEventTagEditParamsAliasesTLPAmber),
			}}),
			AliasGroupNames:            cloudflare.F([]string{"string"}),
			AliasGroupNamesInternal:    cloudflare.F([]string{"string"}),
			AnalyticPriority:           cloudflare.F(0.000000),
			AttributionConfidence:      cloudflare.F("attributionConfidence"),
			AttributionConfidenceScore: cloudflare.F(int64(7)),
			AttributionOrganization:    cloudflare.F("attributionOrganization"),
			CategoryUUID:               cloudflare.F("12345678-1234-1234-1234-1234567890ab"),
			DateOfDiscovery:            cloudflare.F("2024-01-15"),
			ExternalReferenceLinks:     cloudflare.F([]string{"string"}),
			ExternalReferences: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsExternalReference{{
				URL:         cloudflare.F("https://example.com/report"),
				Description: cloudflare.F("Vendor threat report"),
			}}),
			InternalAliases: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsInternalAlias{{
				Value:      cloudflare.F("Fancy Bear"),
				Confidence: cloudflare.F(int64(8)),
				TLP:        cloudflare.F(cloudforce_one.ThreatEventTagEditParamsInternalAliasesTLPAmber),
			}}),
			InternalDescription:     cloudflare.F("internalDescription"),
			Motive:                  cloudflare.F("Espionage"),
			MotiveConfidence:        cloudflare.F(int64(7)),
			OpsecLevel:              cloudflare.F("opsecLevel"),
			OriginCountryConfidence: cloudflare.F(int64(7)),
			OriginCountryISO:        cloudflare.F("originCountryISO"),
			OriginCountryTLP:        cloudflare.F(cloudforce_one.ThreatEventTagEditParamsOriginCountryTLPAmber),
			Priority:                cloudflare.F(0.000000),
			SophisticationLevel:     cloudflare.F("sophisticationLevel"),
			Value:                   cloudflare.F("APT28"),
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
