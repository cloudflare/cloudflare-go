// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/r2"
)

func TestSuperSlurperConnectivityPrecheckSourceWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.SuperSlurper.ConnectivityPrecheck.Source(context.TODO(), r2.SuperSlurperConnectivityPrecheckSourceParams{
		AccountID: cloudflare.F("account_id"),
		Body: r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchema{
			Bucket: cloudflare.F("bucket"),
			Secret: cloudflare.F(r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaSecret{
				AccessKeyID:     cloudflare.F("accessKeyId"),
				SecretAccessKey: cloudflare.F("secretAccessKey"),
			}),
			Vendor:     cloudflare.F(r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendorS3),
			Endpoint:   cloudflare.F("endpoint"),
			Keys:       cloudflare.F([]string{"string"}),
			PathPrefix: cloudflare.F("pathPrefix"),
			Region:     cloudflare.F("region"),
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

func TestSuperSlurperConnectivityPrecheckTargetWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.SuperSlurper.ConnectivityPrecheck.Target(context.TODO(), r2.SuperSlurperConnectivityPrecheckTargetParams{
		AccountID: cloudflare.F("account_id"),
		Bucket:    cloudflare.F("bucket"),
		Secret: cloudflare.F(r2.SuperSlurperConnectivityPrecheckTargetParamsSecret{
			AccessKeyID:     cloudflare.F("accessKeyId"),
			SecretAccessKey: cloudflare.F("secretAccessKey"),
		}),
		Vendor:       cloudflare.F(r2.ProviderR2),
		Jurisdiction: cloudflare.F(r2.SuperSlurperConnectivityPrecheckTargetParamsJurisdictionDefault),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
