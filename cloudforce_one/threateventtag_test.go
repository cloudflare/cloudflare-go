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
		AccountID:      cloudflare.F("account_id"),
		Value:          cloudflare.F("APT28"),
		ActiveDuration: cloudflare.F[cloudforce_one.ThreatEventTagNewParamsActiveDurationUnion](shared.UnionString("string")),
		ActorCategory:  cloudflare.F[cloudforce_one.ThreatEventTagNewParamsActorCategoryUnion](shared.UnionString("string")),
		Aliases: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsAlias{{
			Value:      cloudflare.F("Fancy Bear"),
			Confidence: cloudflare.F(int64(8)),
			TLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsAliasesTLPAmber),
		}}),
		AliasGroupNames:         cloudflare.F([]string{"string"}),
		AliasGroupNamesInternal: cloudflare.F([]string{"string"}),
		AttributionOrganization: cloudflare.F[cloudforce_one.ThreatEventTagNewParamsAttributionOrganizationUnion](shared.UnionString("string")),
		CategoryUUID:            cloudflare.F("12345678-1234-1234-1234-1234567890ab"),
		Confidence:              cloudflare.F(int64(8)),
		DateOfDiscovery:         cloudflare.F("2024-01-15T00:00:00Z"),
		Description:             cloudflare.F("A suspected state-sponsored group."),
		ExternalReferenceLinks:  cloudflare.F([]string{"string"}),
		ExternalReferences: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsExternalReference{{
			URL:         cloudflare.F("https://example.com/report"),
			Description: cloudflare.F("Vendor threat report"),
		}}),
		InternalAliases: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsInternalAlias{{
			Value:      cloudflare.F("Fancy Bear"),
			Confidence: cloudflare.F(int64(8)),
			TLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsInternalAliasesTLPAmber),
		}}),
		InternalDescription: cloudflare.F("internalDescription"),
		LastSeen:            cloudflare.F("lastSeen"),
		Motive:              cloudflare.F[cloudforce_one.ThreatEventTagNewParamsMotiveUnion](shared.UnionString("string")),
		OpsecLevel:          cloudflare.F[cloudforce_one.ThreatEventTagNewParamsOpsecLevelUnion](shared.UnionString("string")),
		OriginCountryISO:    cloudflare.F[cloudforce_one.ThreatEventTagNewParamsOriginCountryISOUnion](shared.UnionString("string")),
		Priority:            cloudflare.F[cloudforce_one.ThreatEventTagNewParamsPriorityUnion](shared.UnionFloat(0.000000)),
		Properties: cloudflare.F(map[string]interface{}{
			"foo": "bar",
		}),
		SophisticationLevel: cloudflare.F[cloudforce_one.ThreatEventTagNewParamsSophisticationLevelUnion](shared.UnionString("string")),
		TLP:                 cloudflare.F(cloudforce_one.ThreatEventTagNewParamsTLPAmber),
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
			Field: cloudflare.F("value"),
			Op:    cloudflare.F(cloudforce_one.ThreatEventTagListParamsFiltersOpIn),
			Value: cloudflare.F[cloudforce_one.ThreatEventTagListParamsFiltersValueUnion](shared.UnionString("APT28")),
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
			AccountID:      cloudflare.F("account_id"),
			ActiveDuration: cloudflare.F[cloudforce_one.ThreatEventTagEditParamsActiveDurationUnion](shared.UnionString("string")),
			ActorCategory:  cloudflare.F[cloudforce_one.ThreatEventTagEditParamsActorCategoryUnion](shared.UnionString("string")),
			Aliases: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsAlias{{
				Value:      cloudflare.F("Fancy Bear"),
				Confidence: cloudflare.F(int64(8)),
				TLP:        cloudflare.F(cloudforce_one.ThreatEventTagEditParamsAliasesTLPAmber),
			}}),
			AliasGroupNames:         cloudflare.F([]string{"string"}),
			AliasGroupNamesInternal: cloudflare.F([]string{"string"}),
			AttributionOrganization: cloudflare.F[cloudforce_one.ThreatEventTagEditParamsAttributionOrganizationUnion](shared.UnionString("string")),
			CategoryUUID:            cloudflare.F("12345678-1234-1234-1234-1234567890ab"),
			Confidence:              cloudflare.F(int64(8)),
			DateOfDiscovery:         cloudflare.F("2024-01-15T00:00:00Z"),
			Description:             cloudflare.F("description"),
			ExternalReferenceLinks:  cloudflare.F([]string{"string"}),
			ExternalReferences: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsExternalReference{{
				URL:         cloudflare.F("https://example.com/report"),
				Description: cloudflare.F("Vendor threat report"),
			}}),
			InternalAliases: cloudflare.F([]cloudforce_one.ThreatEventTagEditParamsInternalAlias{{
				Value:      cloudflare.F("Fancy Bear"),
				Confidence: cloudflare.F(int64(8)),
				TLP:        cloudflare.F(cloudforce_one.ThreatEventTagEditParamsInternalAliasesTLPAmber),
			}}),
			InternalDescription: cloudflare.F("internalDescription"),
			LastSeen:            cloudflare.F("lastSeen"),
			Motive:              cloudflare.F[cloudforce_one.ThreatEventTagEditParamsMotiveUnion](shared.UnionString("string")),
			OpsecLevel:          cloudflare.F[cloudforce_one.ThreatEventTagEditParamsOpsecLevelUnion](shared.UnionString("string")),
			OriginCountryISO:    cloudflare.F[cloudforce_one.ThreatEventTagEditParamsOriginCountryISOUnion](shared.UnionString("string")),
			Priority:            cloudflare.F[cloudforce_one.ThreatEventTagEditParamsPriorityUnion](shared.UnionFloat(0.000000)),
			Properties: cloudflare.F(map[string]interface{}{
				"foo": "bar",
			}),
			SophisticationLevel: cloudflare.F[cloudforce_one.ThreatEventTagEditParamsSophisticationLevelUnion](shared.UnionString("string")),
			TLP:                 cloudflare.F(cloudforce_one.ThreatEventTagEditParamsTLPAmber),
			Value:               cloudflare.F("APT28"),
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
