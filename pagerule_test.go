// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestPageruleNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Pagerules.New(context.TODO(), cloudflare.PageruleNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: cloudflare.F([]cloudflare.PageruleNewParamsAction{{
			Name: cloudflare.F(cloudflare.PageruleNewParamsActionsNameForwardURL),
			Value: cloudflare.F(cloudflare.PageruleNewParamsActionsValue{
				Type: cloudflare.F(cloudflare.PageruleNewParamsActionsValueTypeTemporary),
				URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}}),
		Targets: cloudflare.F([]cloudflare.PageruleNewParamsTarget{{
			Constraint: cloudflare.F(cloudflare.PageruleNewParamsTargetsConstraint{
				Operator: cloudflare.F(cloudflare.PageruleNewParamsTargetsConstraintOperatorMatches),
				Value:    cloudflare.F("*example.com/images/*"),
			}),
			Target: cloudflare.F(cloudflare.PageruleNewParamsTargetsTargetURL),
		}}),
		Priority: cloudflare.F(int64(0)),
		Status:   cloudflare.F(cloudflare.PageruleNewParamsStatusActive),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		cloudflare.PageruleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]cloudflare.PageruleUpdateParamsAction{{
				Name: cloudflare.F(cloudflare.PageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValue{
					Type: cloudflare.F(cloudflare.PageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: cloudflare.F([]cloudflare.PageruleUpdateParamsTarget{{
				Constraint: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsConstraint{
					Operator: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(cloudflare.PageruleUpdateParamsTargetsTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(cloudflare.PageruleUpdateParamsStatusActive),
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
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Pagerules.List(context.TODO(), cloudflare.PageruleListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(cloudflare.PageruleListParamsDirectionDesc),
		Match:     cloudflare.F(cloudflare.PageruleListParamsMatchAny),
		Order:     cloudflare.F(cloudflare.PageruleListParamsOrderStatus),
		Status:    cloudflare.F(cloudflare.PageruleListParamsStatusActive),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		cloudflare.PageruleDeleteParams{
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

func TestPageruleEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		cloudflare.PageruleEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]cloudflare.PageruleEditParamsAction{{
				Name: cloudflare.F(cloudflare.PageruleEditParamsActionsNameForwardURL),
				Value: cloudflare.F(cloudflare.PageruleEditParamsActionsValue{
					Type: cloudflare.F(cloudflare.PageruleEditParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(cloudflare.PageruleEditParamsStatusActive),
			Targets: cloudflare.F([]cloudflare.PageruleEditParamsTarget{{
				Constraint: cloudflare.F(cloudflare.PageruleEditParamsTargetsConstraint{
					Operator: cloudflare.F(cloudflare.PageruleEditParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(cloudflare.PageruleEditParamsTargetsTargetURL),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		cloudflare.PageruleGetParams{
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
