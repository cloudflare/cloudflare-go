// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_sending"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

func TestEmailSendingSendWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSending.Send(context.TODO(), email_sending.EmailSendingSendParams{
		AccountID: cloudflare.F("account_id"),
		From:      cloudflare.F[email_sending.EmailSendingSendParamsFromUnion](shared.UnionString("sender@example.com")),
		Subject:   cloudflare.F("Monthly Report"),
		To:        cloudflare.F[email_sending.EmailSendingSendParamsToUnion](email_sending.EmailSendingSendParamsToEmailSendingEmailAddressList([]string{"recipient@example.com"})),
		Attachments: cloudflare.F([]email_sending.EmailSendingSendParamsAttachmentUnion{email_sending.EmailSendingSendParamsAttachmentsEmailSendingEmailAttachment{
			Content:     cloudflare.F("JVBERi0xLjQK..."),
			Disposition: cloudflare.F(email_sending.EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDispositionAttachment),
			Filename:    cloudflare.F("report.pdf"),
			Type:        cloudflare.F("application/pdf"),
		}}),
		Bcc: cloudflare.F[email_sending.EmailSendingSendParamsBccUnion](shared.UnionString("user@example.com")),
		Cc:  cloudflare.F[email_sending.EmailSendingSendParamsCcUnion](shared.UnionString("user@example.com")),
		Headers: cloudflare.F(map[string]string{
			"X-Custom-Header": "value",
		}),
		HTML:    cloudflare.F("<h1>Hello</h1><p>Please find your report attached.</p>"),
		ReplyTo: cloudflare.F[email_sending.EmailSendingSendParamsReplyToUnion](shared.UnionString("user@example.com")),
		Text:    cloudflare.F("Hello\n\nPlease find your report attached."),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailSendingSendRaw(t *testing.T) {
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
	_, err := client.EmailSending.SendRaw(context.TODO(), email_sending.EmailSendingSendRawParams{
		AccountID:   cloudflare.F("account_id"),
		From:        cloudflare.F("sender@example.com"),
		MimeMessage: cloudflare.F("From: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Hello\r\nContent-Type: text/plain\r\n\r\nHello, World!"),
		Recipients:  cloudflare.F([]string{"recipient@example.com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
