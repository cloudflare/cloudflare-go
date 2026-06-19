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
		ActiveDuration: cloudflare.F("activeDuration"),
		ActorCategory:  cloudflare.F("Nation State"),
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
		InternalAliases: cloudflare.F([]cloudforce_one.ThreatEventTagNewParamsInternalAlias{{
			Value:      cloudflare.F("Fancy Bear"),
			Confidence: cloudflare.F(int64(8)),
			TLP:        cloudflare.F(cloudforce_one.ThreatEventTagNewParamsInternalAliasesTLPAmber),
		}}),
		InternalDescription: cloudflare.F("internalDescription"),
		Motive:              cloudflare.F("Espionage"),
		OpsecLevel:          cloudflare.F("opsecLevel"),
		OriginCountryISO:    cloudflare.F("originCountryISO"),
		Priority:            cloudflare.F(0.000000),
		SophisticationLevel: cloudflare.F("sophisticationLevel"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
