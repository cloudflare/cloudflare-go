// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestAccessRuleNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.AccessRules.New(context.TODO(), firewall.AccessRuleNewParams{
		Configuration: cloudflare.F[firewall.AccessRuleNewParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
			Target: cloudflare.F(firewall.AccessRuleIPConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		}),
		Mode:      cloudflare.F(firewall.AccessRuleNewParamsModeBlock),
		AccountID: cloudflare.F("account_id"),
		Notes:     cloudflare.F("This rule is enabled because of an event that occurred on date X."),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessRuleListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.AccessRules.List(context.TODO(), firewall.AccessRuleListParams{
		AccountID: cloudflare.F("account_id"),
		Configuration: cloudflare.F(firewall.AccessRuleListParamsConfiguration{
			Target: cloudflare.F(firewall.AccessRuleListParamsConfigurationTargetIP),
			Value:  cloudflare.F("198.51.100.4"),
		}),
		Direction: cloudflare.F(firewall.AccessRuleListParamsDirectionAsc),
		Match:     cloudflare.F(firewall.AccessRuleListParamsMatchAny),
		Mode:      cloudflare.F(firewall.AccessRuleListParamsModeBlock),
		Notes:     cloudflare.F("my note"),
		Order:     cloudflare.F(firewall.AccessRuleListParamsOrderConfigurationTarget),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(20.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessRuleDeleteWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.AccessRules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.AccessRuleDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestAccessRuleEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.AccessRules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.AccessRuleEditParams{
			Configuration: cloudflare.F[firewall.AccessRuleEditParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
				Target: cloudflare.F(firewall.AccessRuleIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			}),
			Mode:      cloudflare.F(firewall.AccessRuleEditParamsModeBlock),
			AccountID: cloudflare.F("account_id"),
			Notes:     cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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

func TestAccessRuleGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Firewall.AccessRules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		firewall.AccessRuleGetParams{
			AccountID: cloudflare.F("account_id"),
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
