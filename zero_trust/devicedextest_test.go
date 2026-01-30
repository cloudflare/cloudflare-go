// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
)

func TestDeviceDEXTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.New(context.TODO(), zero_trust.DeviceDEXTestNewParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Data: cloudflare.F(zero_trust.DeviceDEXTestNewParamsData{
			Host:   cloudflare.F("https://dash.cloudflare.com"),
			Kind:   cloudflare.F(zero_trust.DeviceDEXTestNewParamsDataKindHTTP),
			Method: cloudflare.F(zero_trust.DeviceDEXTestNewParamsDataMethodGet),
		}),
		Enabled:     cloudflare.F(true),
		Interval:    cloudflare.F("30m"),
		Name:        cloudflare.F("HTTP dash health check"),
		Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
		TargetPolicies: cloudflare.F([]zero_trust.DeviceDEXTestNewParamsTargetPolicy{{
			ID:      cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
			Default: cloudflare.F(true),
			Name:    cloudflare.F("name"),
		}}),
		Targeted: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceDEXTestUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DeviceDEXTestUpdateParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
			Data: cloudflare.F(zero_trust.DeviceDEXTestUpdateParamsData{
				Host:   cloudflare.F("https://dash.cloudflare.com"),
				Kind:   cloudflare.F(zero_trust.DeviceDEXTestUpdateParamsDataKindHTTP),
				Method: cloudflare.F(zero_trust.DeviceDEXTestUpdateParamsDataMethodGet),
			}),
			Enabled:     cloudflare.F(true),
			Interval:    cloudflare.F("30m"),
			Name:        cloudflare.F("HTTP dash health check"),
			Description: cloudflare.F("Checks the dash endpoint every 30 minutes"),
			TargetPolicies: cloudflare.F([]zero_trust.DeviceDEXTestUpdateParamsTargetPolicy{{
				ID:      cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				Default: cloudflare.F(true),
				Name:    cloudflare.F("name"),
			}}),
			Targeted: cloudflare.F(true),
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

func TestDeviceDEXTestListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.List(context.TODO(), zero_trust.DeviceDEXTestListParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Kind:      cloudflare.F(zero_trust.DeviceDEXTestListParamsKindHTTP),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(1.000000),
		TestName:  cloudflare.F("testName"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceDEXTestDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DeviceDEXTestDeleteParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
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

func TestDeviceDEXTestGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.DEXTests.Get(
		context.TODO(),
		"372e67954025e0ba6aaa6d586b9e0b59",
		zero_trust.DeviceDEXTestGetParams{
			AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
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
