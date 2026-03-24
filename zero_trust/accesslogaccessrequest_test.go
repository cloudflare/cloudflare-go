// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
)

func TestAccessLogAccessRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Logs.AccessRequests.List(context.TODO(), zero_trust.AccessLogAccessRequestListParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AllowedOp:     cloudflare.F(zero_trust.AccessLogAccessRequestListParamsAllowedOpEq),
		AppTypeOp:     cloudflare.F(zero_trust.AccessLogAccessRequestListParamsAppTypeOpEq),
		AppUIDOp:      cloudflare.F(zero_trust.AccessLogAccessRequestListParamsAppUIDOpEq),
		CountryCodeOp: cloudflare.F(zero_trust.AccessLogAccessRequestListParamsCountryCodeOpEq),
		Direction:     cloudflare.F(zero_trust.AccessLogAccessRequestListParamsDirectionDesc),
		Email:         cloudflare.F("user@example.com"),
		EmailExact:    cloudflare.F(true),
		EmailOp:       cloudflare.F(zero_trust.AccessLogAccessRequestListParamsEmailOpEq),
		Fields:        cloudflare.F("fields"),
		IdPOp:         cloudflare.F(zero_trust.AccessLogAccessRequestListParamsIdPOpEq),
		Limit:         cloudflare.F(int64(0)),
		NonIdentityOp: cloudflare.F(zero_trust.AccessLogAccessRequestListParamsNonIdentityOpEq),
		Page:          cloudflare.F(int64(0)),
		PerPage:       cloudflare.F(int64(0)),
		RayIDOp:       cloudflare.F(zero_trust.AccessLogAccessRequestListParamsRayIDOpEq),
		Since:         cloudflare.F(time.Now()),
		Until:         cloudflare.F(time.Now()),
		UserID:        cloudflare.F("f757c5c3-c1b2-50f7-9126-150a099b6f7e"),
		UserIDOp:      cloudflare.F(zero_trust.AccessLogAccessRequestListParamsUserIDOpEq),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
