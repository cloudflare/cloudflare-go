// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/page_rules"
	"github.com/cloudflare/cloudflare-go/v4/zones"
)

func TestPageRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PageRules.New(context.TODO(), page_rules.PageRuleNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: cloudflare.F([]page_rules.PageRuleNewParamsActionUnion{zones.BrowserCheckParam{
			ID:    cloudflare.F(zones.BrowserCheckIDBrowserCheck),
			Value: cloudflare.F(zones.BrowserCheckValueOn),
		}}),
		Targets: cloudflare.F([]page_rules.TargetParam{{
			Constraint: cloudflare.F(page_rules.TargetConstraintParam{
				Operator: cloudflare.F(page_rules.TargetConstraintOperatorMatches),
				Value:    cloudflare.F("*example.com/images/*"),
			}),
			Target: cloudflare.F(page_rules.TargetTargetURL),
		}}),
		Priority: cloudflare.F(int64(0)),
		Status:   cloudflare.F(page_rules.PageRuleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PageRules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_rules.PageRuleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]page_rules.PageRuleUpdateParamsActionUnion{zones.BrowserCheckParam{
				ID:    cloudflare.F(zones.BrowserCheckIDBrowserCheck),
				Value: cloudflare.F(zones.BrowserCheckValueOn),
			}}),
			Targets: cloudflare.F([]page_rules.TargetParam{{
				Constraint: cloudflare.F(page_rules.TargetConstraintParam{
					Operator: cloudflare.F(page_rules.TargetConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(page_rules.TargetTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(page_rules.PageRuleUpdateParamsStatusActive),
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

func TestPageRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.PageRules.List(context.TODO(), page_rules.PageRuleListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(page_rules.PageRuleListParamsDirectionAsc),
		Match:     cloudflare.F(page_rules.PageRuleListParamsMatchAny),
		Order:     cloudflare.F(page_rules.PageRuleListParamsOrderStatus),
		Status:    cloudflare.F(page_rules.PageRuleListParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageRuleDelete(t *testing.T) {
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
	_, err := client.PageRules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_rules.PageRuleDeleteParams{
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

func TestPageRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.PageRules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_rules.PageRuleEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]page_rules.PageRuleEditParamsActionUnion{zones.BrowserCheckParam{
				ID:    cloudflare.F(zones.BrowserCheckIDBrowserCheck),
				Value: cloudflare.F(zones.BrowserCheckValueOn),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(page_rules.PageRuleEditParamsStatusActive),
			Targets: cloudflare.F([]page_rules.TargetParam{{
				Constraint: cloudflare.F(page_rules.TargetConstraintParam{
					Operator: cloudflare.F(page_rules.TargetConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(page_rules.TargetTargetURL),
			}}),
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

func TestPageRuleGet(t *testing.T) {
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
	_, err := client.PageRules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		page_rules.PageRuleGetParams{
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
