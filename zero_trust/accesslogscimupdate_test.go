// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestAccessLogSCIMUpdateListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Logs.SCIM.Updates.List(context.TODO(), zero_trust.AccessLogSCIMUpdateListParams{
		AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IdPID:             cloudflare.F([]string{"df7e2w5f-02b7-4d9d-af26-8d1988fca630", "0194ae2c-efcf-7cfb-8884-055f1a161fa5"}),
		CfResourceID:      cloudflare.F("bd97ef8d-7986-43e3-9ee0-c25dda33e4b0"),
		Direction:         cloudflare.F(zero_trust.AccessLogSCIMUpdateListParamsDirectionDesc),
		IdPResourceID:     cloudflare.F("idp_resource_id"),
		Limit:             cloudflare.F(int64(10)),
		RequestMethod:     cloudflare.F([]zero_trust.AccessLogSCIMUpdateListParamsRequestMethod{zero_trust.AccessLogSCIMUpdateListParamsRequestMethodDelete, zero_trust.AccessLogSCIMUpdateListParamsRequestMethodPatch}),
		ResourceGroupName: cloudflare.F("ALL_EMPLOYEES"),
		ResourceType:      cloudflare.F([]zero_trust.AccessLogSCIMUpdateListParamsResourceType{zero_trust.AccessLogSCIMUpdateListParamsResourceTypeUser, zero_trust.AccessLogSCIMUpdateListParamsResourceTypeGroup}),
		ResourceUserEmail: cloudflare.F("john.smith@example.com"),
		Since:             cloudflare.F(time.Now()),
		Status:            cloudflare.F([]zero_trust.AccessLogSCIMUpdateListParamsStatus{zero_trust.AccessLogSCIMUpdateListParamsStatusFailure, zero_trust.AccessLogSCIMUpdateListParamsStatusSuccess}),
		Until:             cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
