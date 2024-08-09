// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_security"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestSettingAllowPatternNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.New(context.TODO(), email_security.SettingAllowPatternNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: email_security.SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern{
			Comments:     cloudflare.F("Trust all messages send from test@example.com"),
			IsRecipient:  cloudflare.F(false),
			IsRegex:      cloudflare.F(false),
			IsSender:     cloudflare.F(true),
			IsSpoof:      cloudflare.F(false),
			Pattern:      cloudflare.F("test@example.com"),
			PatternType:  cloudflare.F(email_security.SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeEmail),
			VerifySender: cloudflare.F(true),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPatternListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.List(context.TODO(), email_security.SettingAllowPatternListParams{
		AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:    cloudflare.F(email_security.SettingAllowPatternListParamsDirectionAsc),
		IsRecipient:  cloudflare.F(true),
		IsSender:     cloudflare.F(true),
		IsSpoof:      cloudflare.F(true),
		Order:        cloudflare.F(email_security.SettingAllowPatternListParamsOrderPattern),
		Page:         cloudflare.F(int64(1)),
		PatternType:  cloudflare.F(email_security.SettingAllowPatternListParamsPatternTypeEmail),
		PerPage:      cloudflare.F(int64(1)),
		Search:       cloudflare.F("search"),
		VerifySender: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPatternDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Delete(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternDeleteParams{
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

func TestSettingAllowPatternEditWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Edit(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternEditParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Comments:     cloudflare.F("comments"),
			IsRecipient:  cloudflare.F(true),
			IsRegex:      cloudflare.F(true),
			IsSender:     cloudflare.F(true),
			IsSpoof:      cloudflare.F(true),
			Pattern:      cloudflare.F("x"),
			PatternType:  cloudflare.F(email_security.SettingAllowPatternEditParamsPatternTypeEmail),
			VerifySender: cloudflare.F(true),
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

func TestSettingAllowPatternGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPatterns.Get(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPatternGetParams{
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
