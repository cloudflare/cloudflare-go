// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/logpush"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.New(context.TODO(), logpush.JobNewParams{
		DestinationConf: cloudflare.F("s3://mybucket/logs?region=us-west-2"),
		AccountID:       cloudflare.F("string"),
		ZoneID:          cloudflare.F("string"),
		Dataset:         cloudflare.F("http_requests"),
		Enabled:         cloudflare.F(false),
		Frequency:       cloudflare.F(logpush.JobNewParamsFrequencyHigh),
		LogpullOptions:  cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
		Name:            cloudflare.F("example.com"),
		OutputOptions: cloudflare.F(logpush.OutputOptionsParam{
			Cve2021_4428:    cloudflare.F(true),
			BatchPrefix:     cloudflare.F("string"),
			BatchSuffix:     cloudflare.F("string"),
			FieldDelimiter:  cloudflare.F("string"),
			FieldNames:      cloudflare.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
			OutputType:      cloudflare.F(logpush.OutputOptionsOutputTypeNdjson),
			RecordDelimiter: cloudflare.F("string"),
			RecordPrefix:    cloudflare.F("string"),
			RecordSuffix:    cloudflare.F("string"),
			RecordTemplate:  cloudflare.F("string"),
			SampleRate:      cloudflare.F(0.000000),
			TimestampFormat: cloudflare.F(logpush.OutputOptionsTimestampFormatUnixnano),
		}),
		OwnershipChallenge: cloudflare.F("00000000000000000000"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Update(
		context.TODO(),
		int64(1),
		logpush.JobUpdateParams{
			AccountID:       cloudflare.F("string"),
			ZoneID:          cloudflare.F("string"),
			DestinationConf: cloudflare.F("s3://mybucket/logs?region=us-west-2"),
			Enabled:         cloudflare.F(false),
			Frequency:       cloudflare.F(logpush.JobUpdateParamsFrequencyHigh),
			LogpullOptions:  cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			OutputOptions: cloudflare.F(logpush.OutputOptionsParam{
				Cve2021_4428:    cloudflare.F(true),
				BatchPrefix:     cloudflare.F("string"),
				BatchSuffix:     cloudflare.F("string"),
				FieldDelimiter:  cloudflare.F("string"),
				FieldNames:      cloudflare.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
				OutputType:      cloudflare.F(logpush.OutputOptionsOutputTypeNdjson),
				RecordDelimiter: cloudflare.F("string"),
				RecordPrefix:    cloudflare.F("string"),
				RecordSuffix:    cloudflare.F("string"),
				RecordTemplate:  cloudflare.F("string"),
				SampleRate:      cloudflare.F(0.000000),
				TimestampFormat: cloudflare.F(logpush.OutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: cloudflare.F("00000000000000000000"),
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

func TestJobListWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.List(context.TODO(), logpush.JobListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Delete(
		context.TODO(),
		int64(1),
		logpush.JobDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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

func TestJobGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Get(
		context.TODO(),
		int64(1),
		logpush.JobGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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
