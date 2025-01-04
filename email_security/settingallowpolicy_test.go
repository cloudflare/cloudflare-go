// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_security"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestSettingAllowPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPolicies.New(context.TODO(), email_security.SettingAllowPolicyNewParams{
		AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IsAcceptableSender: cloudflare.F(false),
		IsExemptRecipient:  cloudflare.F(false),
		IsRegex:            cloudflare.F(false),
		IsTrustedSender:    cloudflare.F(true),
		Pattern:            cloudflare.F("test@example.com"),
		PatternType:        cloudflare.F(email_security.SettingAllowPolicyNewParamsPatternTypeEmail),
		VerifySender:       cloudflare.F(true),
		Comments:           cloudflare.F("Trust all messages send from test@example.com"),
		IsRecipient:        cloudflare.F(false),
		IsSender:           cloudflare.F(true),
		IsSpoof:            cloudflare.F(false),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPolicies.List(context.TODO(), email_security.SettingAllowPolicyListParams{
		AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction:          cloudflare.F(email_security.SettingAllowPolicyListParamsDirectionAsc),
		IsAcceptableSender: cloudflare.F(true),
		IsExemptRecipient:  cloudflare.F(true),
		IsRecipient:        cloudflare.F(true),
		IsSender:           cloudflare.F(true),
		IsSpoof:            cloudflare.F(true),
		IsTrustedSender:    cloudflare.F(true),
		Order:              cloudflare.F(email_security.SettingAllowPolicyListParamsOrderPattern),
		Page:               cloudflare.F(int64(1)),
		PatternType:        cloudflare.F(email_security.SettingAllowPolicyListParamsPatternTypeEmail),
		PerPage:            cloudflare.F(int64(1)),
		Search:             cloudflare.F("search"),
		VerifySender:       cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingAllowPolicyDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPolicies.Delete(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPolicyDeleteParams{
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

func TestSettingAllowPolicyEditWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPolicies.Edit(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPolicyEditParams{
			AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Comments:           cloudflare.F("comments"),
			IsAcceptableSender: cloudflare.F(true),
			IsExemptRecipient:  cloudflare.F(true),
			IsRegex:            cloudflare.F(true),
			IsTrustedSender:    cloudflare.F(true),
			Pattern:            cloudflare.F("x"),
			PatternType:        cloudflare.F(email_security.SettingAllowPolicyEditParamsPatternTypeEmail),
			VerifySender:       cloudflare.F(true),
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

func TestSettingAllowPolicyGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.AllowPolicies.Get(
		context.TODO(),
		int64(2401),
		email_security.SettingAllowPolicyGetParams{
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
