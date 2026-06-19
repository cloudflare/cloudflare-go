// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestInvestigateBulkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.Bulk.New(context.TODO(), email_security.InvestigateBulkNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Action:    cloudflare.F(email_security.InvestigateBulkNewParamsActionMove),
		SearchParams: cloudflare.F(email_security.InvestigateBulkNewParamsSearchParams{
			ActionLog:        cloudflare.F(true),
			AlertID:          cloudflare.F("alert_id"),
			DeliveryStatus:   cloudflare.F(email_security.InvestigateBulkNewParamsSearchParamsDeliveryStatusDelivered),
			DetectionsOnly:   cloudflare.F(true),
			Domain:           cloudflare.F("domain"),
			End:              cloudflare.F(time.Now()),
			ExactSubject:     cloudflare.F("exact_subject"),
			FinalDisposition: cloudflare.F(email_security.InvestigateBulkNewParamsSearchParamsFinalDispositionMalicious),
			MessageAction:    cloudflare.F(email_security.InvestigateBulkNewParamsSearchParamsMessageActionPreview),
			MessageID:        cloudflare.F("message_id"),
			Metric:           cloudflare.F("metric"),
			Query:            cloudflare.F("query"),
			Recipient:        cloudflare.F("recipient"),
			Sender:           cloudflare.F("sender"),
			Start:            cloudflare.F(time.Now()),
			Subject:          cloudflare.F("subject"),
			Submissions:      cloudflare.F(true),
		}),
		Comment:             cloudflare.F("comment"),
		Destination:         cloudflare.F(email_security.InvestigateBulkNewParamsDestinationInbox),
		ExpectedDisposition: cloudflare.F(email_security.InvestigateBulkNewParamsExpectedDispositionMalicious),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestigateBulkListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.Bulk.List(context.TODO(), email_security.InvestigateBulkListParams{
		AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ActionType: cloudflare.F(email_security.InvestigateBulkListParamsActionTypeMove),
		Page:       cloudflare.F(int64(1)),
		PerPage:    cloudflare.F(int64(20)),
		Status:     cloudflare.F(email_security.InvestigateBulkListParamsStatusPending),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestigateBulkDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.Bulk.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		email_security.InvestigateBulkDeleteParams{
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

func TestInvestigateBulkGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Investigate.Bulk.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		email_security.InvestigateBulkGetParams{
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
