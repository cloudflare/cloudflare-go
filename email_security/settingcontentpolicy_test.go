// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestSettingContentPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.New(context.TODO(), email_security.SettingContentPolicyNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Enabled:   cloudflare.F(true),
		Name:      cloudflare.F("Block phishing keywords"),
		Pattern:   cloudflare.F("urgent.*verify.*account"),
		Targets:   cloudflare.F([]email_security.SettingContentPolicyNewParamsTarget{email_security.SettingContentPolicyNewParamsTargetSubject}),
		Notes:     cloudflare.F("Blocks common phishing subject lines"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingContentPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.List(context.TODO(), email_security.SettingContentPolicyListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(email_security.SettingContentPolicyListParamsDirectionAsc),
		Enabled:   cloudflare.F(true),
		Name:      cloudflare.F("name"),
		Order:     cloudflare.F(email_security.SettingContentPolicyListParamsOrderName),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(20)),
		Search:    cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingContentPolicyDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingContentPolicyDeleteParams{
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

func TestSettingContentPolicyBatch(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.Batch(context.TODO(), email_security.SettingContentPolicyBatchParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Deletes: cloudflare.F([]email_security.SettingContentPolicyBatchParamsDelete{{
			ID: cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		}}),
		Patches: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPatch{{
			Enabled: cloudflare.F(true),
			Name:    cloudflare.F("Block phishing keywords"),
			Notes:   cloudflare.F("Blocks common phishing subject lines"),
			Pattern: cloudflare.F("urgent.*verify.*account"),
			Targets: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPatchesTarget{email_security.SettingContentPolicyBatchParamsPatchesTargetSubject}),
		}}),
		Posts: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPost{{
			Enabled: cloudflare.F(true),
			Name:    cloudflare.F("Block phishing keywords"),
			Pattern: cloudflare.F("urgent.*verify.*account"),
			Targets: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPostsTarget{email_security.SettingContentPolicyBatchParamsPostsTargetSubject}),
			Notes:   cloudflare.F("Blocks common phishing subject lines"),
		}}),
		Puts: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPut{{
			Enabled: cloudflare.F(true),
			Name:    cloudflare.F("Block phishing keywords"),
			Pattern: cloudflare.F("urgent.*verify.*account"),
			Targets: cloudflare.F([]email_security.SettingContentPolicyBatchParamsPutsTarget{email_security.SettingContentPolicyBatchParamsPutsTargetSubject}),
			Notes:   cloudflare.F("Blocks common phishing subject lines"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingContentPolicyEditWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingContentPolicyEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Enabled:   cloudflare.F(true),
			Name:      cloudflare.F("Block phishing keywords"),
			Notes:     cloudflare.F("Blocks common phishing subject lines"),
			Pattern:   cloudflare.F("urgent.*verify.*account"),
			Targets:   cloudflare.F([]email_security.SettingContentPolicyEditParamsTarget{email_security.SettingContentPolicyEditParamsTargetSubject}),
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

func TestSettingContentPolicyGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.ContentPolicies.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingContentPolicyGetParams{
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
