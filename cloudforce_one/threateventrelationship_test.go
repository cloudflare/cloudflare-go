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

func TestThreatEventRelationshipListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Relationships.List(
		context.TODO(),
		"event_id",
		cloudforce_one.ThreatEventRelationshipListParams{
			AccountID:         cloudflare.F("account_id"),
			DatasetID:         cloudflare.F("datasetId"),
			Direction:         cloudflare.F(cloudforce_one.ThreatEventRelationshipListParamsDirectionAncestors),
			IncludeParent:     cloudflare.F(true),
			IndicatorTypeIDs:  cloudflare.F([]string{"string"}),
			MaxDepth:          cloudflare.F(0.000000),
			Page:              cloudflare.F(0.000000),
			PageSize:          cloudflare.F(0.000000),
			RelationshipTypes: cloudflare.F[cloudforce_one.ThreatEventRelationshipListParamsRelationshipTypesUnion](shared.UnionString("string")),
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
