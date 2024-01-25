// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestPageruleUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.PageruleUpdateParams{
			Actions: cloudflare.F([]cloudflare.PageruleUpdateParamsAction{{
				Name: cloudflare.F(cloudflare.PageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.PageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}, {
				Name: cloudflare.F(cloudflare.PageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(cloudflare.PageruleUpdateParamsStatusActive),
			Targets: cloudflare.F([]cloudflare.PageruleUpdateParamsTarget{{
				Constraint: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsConstraint{
					Operator: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsTargetURL),
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
