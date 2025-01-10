// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_security"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestInvestigateListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.List(context.TODO(), email_security.InvestigateListParams{
		AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ActionLog:        cloudflare.F(true),
		AlertID:          cloudflare.F("alert_id"),
		DetectionsOnly:   cloudflare.F(true),
		Domain:           cloudflare.F("domain"),
		End:              cloudflare.F(time.Now()),
		FinalDisposition: cloudflare.F(email_security.InvestigateListParamsFinalDispositionMalicious),
		MessageAction:    cloudflare.F(email_security.InvestigateListParamsMessageActionPreview),
		MessageID:        cloudflare.F("message_id"),
		Metric:           cloudflare.F("metric"),
		Page:             cloudflare.F(int64(1)),
		PerPage:          cloudflare.F(int64(1)),
		Query:            cloudflare.F("query"),
		Recipient:        cloudflare.F("recipient"),
		Sender:           cloudflare.F("sender"),
		Start:            cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestigateGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.Get(
		context.TODO(),
		"4Njp3P0STMz2c02Q",
		email_security.InvestigateGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
