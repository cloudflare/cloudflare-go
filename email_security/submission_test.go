// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/email_security"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestSubmissionListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Submissions.List(context.TODO(), email_security.SubmissionListParams{
		AccountID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		End:                  cloudflare.F(time.Now()),
		OriginalDisposition:  cloudflare.F(email_security.SubmissionListParamsOriginalDispositionMalicious),
		OutcomeDisposition:   cloudflare.F(email_security.SubmissionListParamsOutcomeDispositionMalicious),
		Page:                 cloudflare.F(int64(1)),
		PerPage:              cloudflare.F(int64(1)),
		Query:                cloudflare.F("query"),
		RequestedDisposition: cloudflare.F(email_security.SubmissionListParamsRequestedDispositionMalicious),
		Start:                cloudflare.F(time.Now()),
		Status:               cloudflare.F("status"),
		SubmissionID:         cloudflare.F("submission_id"),
		Type:                 cloudflare.F(email_security.SubmissionListParamsTypeTeam),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
