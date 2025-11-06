// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/token_validation"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.New(context.TODO(), token_validation.RuleNewParams{
		ZoneID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Action:      cloudflare.F(token_validation.RuleNewParamsActionLog),
		Description: cloudflare.F("Long description for Token Validation Rule"),
		Enabled:     cloudflare.F(true),
		Expression:  cloudflare.F(`is_jwt_valid("52973293-cb04-4a97-8f55-e7d2ad1107dd") or is_jwt_valid("46eab8d1-6376-45e3-968f-2c649d77d423")`),
		Selector: cloudflare.F(token_validation.RuleNewParamsSelector{
			Exclude: cloudflare.F([]token_validation.RuleNewParamsSelectorExclude{{
				OperationIDs: cloudflare.F([]string{"f9c5615e-fe15-48ce-bec6-cfc1946f1bec", "56828eae-035a-4396-ba07-51c66d680a04"}),
			}}),
			Include: cloudflare.F([]token_validation.RuleNewParamsSelectorInclude{{
				Host: cloudflare.F([]string{"v1.example.com", "v2.example.com"}),
			}}),
		}),
		Title: cloudflare.F("Example Token Validation Rule"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.List(context.TODO(), token_validation.RuleListParams{
		ZoneID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ID:                 cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		Action:             cloudflare.F(token_validation.RuleListParamsActionLog),
		Enabled:            cloudflare.F(true),
		Host:               cloudflare.F("www.example.com"),
		Hostname:           cloudflare.F("www.example.com"),
		Page:               cloudflare.F(int64(1)),
		PerPage:            cloudflare.F(int64(5)),
		RuleID:             cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		TokenConfiguration: cloudflare.F([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleDelete(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.Delete(
		context.TODO(),
		"4a7ee8d3-dd63-4ceb-9d5f-c27831854ce7",
		token_validation.RuleDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestRuleBulkNew(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.BulkNew(context.TODO(), token_validation.RuleBulkNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: []token_validation.RuleBulkNewParamsBody{{
			Action:      cloudflare.F(token_validation.RuleBulkNewParamsBodyActionLog),
			Description: cloudflare.F("Long description for Token Validation Rule"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F(`is_jwt_valid("52973293-cb04-4a97-8f55-e7d2ad1107dd") or is_jwt_valid("46eab8d1-6376-45e3-968f-2c649d77d423")`),
			Selector: cloudflare.F(token_validation.RuleBulkNewParamsBodySelector{
				Exclude: cloudflare.F([]token_validation.RuleBulkNewParamsBodySelectorExclude{{
					OperationIDs: cloudflare.F([]string{"f9c5615e-fe15-48ce-bec6-cfc1946f1bec", "56828eae-035a-4396-ba07-51c66d680a04"}),
				}}),
				Include: cloudflare.F([]token_validation.RuleBulkNewParamsBodySelectorInclude{{
					Host: cloudflare.F([]string{"v1.example.com", "v2.example.com"}),
				}}),
			}),
			Title: cloudflare.F("Example Token Validation Rule"),
		}},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleBulkEdit(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.BulkEdit(context.TODO(), token_validation.RuleBulkEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: []token_validation.RuleBulkEditParamsBody{{
			ID:          cloudflare.F("0d9bf70c-92e1-4bb3-9411-34a3bcc59003"),
			Action:      cloudflare.F(token_validation.RuleBulkEditParamsBodyActionLog),
			Description: cloudflare.F("Long description for Token Validation Rule"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F(`is_jwt_valid("52973293-cb04-4a97-8f55-e7d2ad1107dd") or is_jwt_valid("46eab8d1-6376-45e3-968f-2c649d77d423")`),
			Position: cloudflare.F[token_validation.RuleBulkEditParamsBodyPositionUnion](token_validation.RuleBulkEditParamsBodyPositionAPIShieldIndex{
				Index: cloudflare.F(int64(2)),
			}),
			Selector: cloudflare.F(token_validation.RuleBulkEditParamsBodySelector{
				Exclude: cloudflare.F([]token_validation.RuleBulkEditParamsBodySelectorExclude{{
					OperationIDs: cloudflare.F([]string{"f9c5615e-fe15-48ce-bec6-cfc1946f1bec", "56828eae-035a-4396-ba07-51c66d680a04"}),
				}}),
				Include: cloudflare.F([]token_validation.RuleBulkEditParamsBodySelectorInclude{{
					Host: cloudflare.F([]string{"v1.example.com", "v2.example.com"}),
				}}),
			}),
			Title: cloudflare.F("Example Token Validation Rule"),
		}},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.Edit(
		context.TODO(),
		"4a7ee8d3-dd63-4ceb-9d5f-c27831854ce7",
		token_validation.RuleEditParams{
			ZoneID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:      cloudflare.F(token_validation.RuleEditParamsActionLog),
			Description: cloudflare.F("Long description for Token Validation Rule"),
			Enabled:     cloudflare.F(true),
			Expression:  cloudflare.F(`is_jwt_valid("52973293-cb04-4a97-8f55-e7d2ad1107dd") or is_jwt_valid("46eab8d1-6376-45e3-968f-2c649d77d423")`),
			Position: cloudflare.F[token_validation.RuleEditParamsPositionUnion](token_validation.RuleEditParamsPositionAPIShieldIndex{
				Index: cloudflare.F(int64(2)),
			}),
			Selector: cloudflare.F(token_validation.RuleEditParamsSelector{
				Exclude: cloudflare.F([]token_validation.RuleEditParamsSelectorExclude{{
					OperationIDs: cloudflare.F([]string{"f9c5615e-fe15-48ce-bec6-cfc1946f1bec", "56828eae-035a-4396-ba07-51c66d680a04"}),
				}}),
				Include: cloudflare.F([]token_validation.RuleEditParamsSelectorInclude{{
					Host: cloudflare.F([]string{"v1.example.com", "v2.example.com"}),
				}}),
			}),
			Title: cloudflare.F("Example Token Validation Rule"),
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

func TestRuleGet(t *testing.T) {
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
	_, err := client.TokenValidation.Rules.Get(
		context.TODO(),
		"4a7ee8d3-dd63-4ceb-9d5f-c27831854ce7",
		token_validation.RuleGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
