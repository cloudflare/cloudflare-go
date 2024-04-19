// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/pagerules"
)

func TestPageruleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.New(context.TODO(), pagerules.PageruleNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: cloudflare.F([]pagerules.RouteParam{{
			Name: cloudflare.F(pagerules.RouteNameForwardURL),
			Value: cloudflare.F(pagerules.RouteValueParam{
				Type: cloudflare.F(pagerules.RouteValueTypeTemporary),
				URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}}),
		Targets: cloudflare.F([]pagerules.TargetParam{{
			Constraint: cloudflare.F(pagerules.TargetConstraintParam{
				Operator: cloudflare.F(pagerules.TargetConstraintOperatorMatches),
				Value:    cloudflare.F("*example.com/images/*"),
			}),
			Target: cloudflare.F(pagerules.TargetTargetURL),
		}}),
		Priority: cloudflare.F(int64(0)),
		Status:   cloudflare.F(pagerules.PageruleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]pagerules.RouteParam{{
				Name: cloudflare.F(pagerules.RouteNameForwardURL),
				Value: cloudflare.F(pagerules.RouteValueParam{
					Type: cloudflare.F(pagerules.RouteValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: cloudflare.F([]pagerules.TargetParam{{
				Constraint: cloudflare.F(pagerules.TargetConstraintParam{
					Operator: cloudflare.F(pagerules.TargetConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(pagerules.TargetTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(pagerules.PageruleUpdateParamsStatusActive),
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

func TestPageruleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.List(context.TODO(), pagerules.PageruleListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(pagerules.PageruleListParamsDirectionDesc),
		Match:     cloudflare.F(pagerules.PageruleListParamsMatchAny),
		Order:     cloudflare.F(pagerules.PageruleListParamsOrderStatus),
		Status:    cloudflare.F(pagerules.PageruleListParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleDelete(t *testing.T) {
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
	_, err := client.Pagerules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:   map[string]interface{}{},
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

func TestPageruleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Pagerules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]pagerules.RouteParam{{
				Name: cloudflare.F(pagerules.RouteNameForwardURL),
				Value: cloudflare.F(pagerules.RouteValueParam{
					Type: cloudflare.F(pagerules.RouteValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(pagerules.PageruleEditParamsStatusActive),
			Targets: cloudflare.F([]pagerules.TargetParam{{
				Constraint: cloudflare.F(pagerules.TargetConstraintParam{
					Operator: cloudflare.F(pagerules.TargetConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(pagerules.TargetTargetURL),
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

func TestPageruleGet(t *testing.T) {
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
	_, err := client.Pagerules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleGetParams{
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
