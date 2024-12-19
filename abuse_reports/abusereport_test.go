// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/abuse_reports"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestAbuseReportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AbuseReports.New(
		context.TODO(),
		abuse_reports.AbuseReportNewParamsReportTypeAbuseDmca,
		abuse_reports.AbuseReportNewParams{
			AccountID:                  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Act:                        cloudflare.F(abuse_reports.AbuseReportNewParamsActAbuseDmca),
			Email:                      cloudflare.F("email"),
			Email2:                     cloudflare.F("email2"),
			HostNotification:           cloudflare.F(abuse_reports.AbuseReportNewParamsHostNotificationSend),
			NcmecNotification:          cloudflare.F(abuse_reports.AbuseReportNewParamsNcmecNotificationSend),
			OwnerNotification:          cloudflare.F(abuse_reports.AbuseReportNewParamsOwnerNotificationSend),
			URLs:                       cloudflare.F("urls"),
			Address1:                   cloudflare.F("x"),
			AgentName:                  cloudflare.F("x"),
			Agree:                      cloudflare.F(abuse_reports.AbuseReportNewParamsAgree0),
			City:                       cloudflare.F("x"),
			Comments:                   cloudflare.F("x"),
			Company:                    cloudflare.F("x"),
			Country:                    cloudflare.F("x"),
			DestinationIPs:             cloudflare.F("destination_ips"),
			Justification:              cloudflare.F("x"),
			Name:                       cloudflare.F("x"),
			NcseiSubjectRepresentation: cloudflare.F(true),
			OriginalWork:               cloudflare.F("x"),
			PortsProtocols:             cloudflare.F("ports_protocols"),
			Signature:                  cloudflare.F("signature"),
			SourceIPs:                  cloudflare.F("source_ips"),
			State:                      cloudflare.F("x"),
			Tele:                       cloudflare.F("x"),
			Title:                      cloudflare.F("x"),
			TrademarkNumber:            cloudflare.F("x"),
			TrademarkOffice:            cloudflare.F("x"),
			TrademarkSymbol:            cloudflare.F("x"),
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
