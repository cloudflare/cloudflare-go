// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/logpush"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.New(context.TODO(), logpush.JobNewParams{
		DestinationConf:          cloudflare.F("s3://mybucket/logs?region=us-west-2"),
		AccountID:                cloudflare.F("account_id"),
		Dataset:                  cloudflare.F(logpush.JobNewParamsDatasetGatewayDNS),
		Enabled:                  cloudflare.F(false),
		Filter:                   cloudflare.F(`{"where":{"and":[{"key":"ClientRequestPath","operator":"contains","value":"/static"},{"key":"ClientRequestHost","operator":"eq","value":"example.com"}]}}`),
		Frequency:                cloudflare.F(logpush.JobNewParamsFrequencyHigh),
		Kind:                     cloudflare.F(logpush.JobNewParamsKindEmpty),
		LogpullOptions:           cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
		MaxUploadBytes:           cloudflare.F(logpush.JobNewParamsMaxUploadBytes0),
		MaxUploadIntervalSeconds: cloudflare.F(logpush.JobNewParamsMaxUploadIntervalSeconds0),
		MaxUploadRecords:         cloudflare.F(logpush.JobNewParamsMaxUploadRecords0),
		Name:                     cloudflare.F("example.com"),
		OutputOptions: cloudflare.F(logpush.OutputOptionsParam{
			BatchPrefix:     cloudflare.F(""),
			BatchSuffix:     cloudflare.F(""),
			Cve2021_44228:   cloudflare.F(false),
			FieldDelimiter:  cloudflare.F(","),
			FieldNames:      cloudflare.F([]string{"Datetime", "DstIP", "SrcIP"}),
			OutputType:      cloudflare.F(logpush.OutputOptionsOutputTypeNdjson),
			RecordDelimiter: cloudflare.F(""),
			RecordPrefix:    cloudflare.F("{"),
			RecordSuffix:    cloudflare.F("}\n"),
			RecordTemplate:  cloudflare.F("record_template"),
			SampleRate:      cloudflare.F(1.000000),
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
	_, err := client.Logpush.Jobs.Update(
		context.TODO(),
		int64(1),
		logpush.JobUpdateParams{
			AccountID:                cloudflare.F("account_id"),
			DestinationConf:          cloudflare.F("s3://mybucket/logs?region=us-west-2"),
			Enabled:                  cloudflare.F(false),
			Filter:                   cloudflare.F(`{"where":{"and":[{"key":"ClientRequestPath","operator":"contains","value":"/static"},{"key":"ClientRequestHost","operator":"eq","value":"example.com"}]}}`),
			Frequency:                cloudflare.F(logpush.JobUpdateParamsFrequencyHigh),
			Kind:                     cloudflare.F(logpush.JobUpdateParamsKindEmpty),
			LogpullOptions:           cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			MaxUploadBytes:           cloudflare.F(logpush.JobUpdateParamsMaxUploadBytes0),
			MaxUploadIntervalSeconds: cloudflare.F(logpush.JobUpdateParamsMaxUploadIntervalSeconds0),
			MaxUploadRecords:         cloudflare.F(logpush.JobUpdateParamsMaxUploadRecords0),
			Name:                     cloudflare.F("example.com"),
			OutputOptions: cloudflare.F(logpush.OutputOptionsParam{
				BatchPrefix:     cloudflare.F(""),
				BatchSuffix:     cloudflare.F(""),
				Cve2021_44228:   cloudflare.F(false),
				FieldDelimiter:  cloudflare.F(","),
				FieldNames:      cloudflare.F([]string{"Datetime", "DstIP", "SrcIP"}),
				OutputType:      cloudflare.F(logpush.OutputOptionsOutputTypeNdjson),
				RecordDelimiter: cloudflare.F(""),
				RecordPrefix:    cloudflare.F("{"),
				RecordSuffix:    cloudflare.F("}\n"),
				RecordTemplate:  cloudflare.F("record_template"),
				SampleRate:      cloudflare.F(1.000000),
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
	_, err := client.Logpush.Jobs.List(context.TODO(), logpush.JobListParams{
		AccountID: cloudflare.F("account_id"),
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
	_, err := client.Logpush.Jobs.Delete(
		context.TODO(),
		int64(1),
		logpush.JobDeleteParams{
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

func TestJobGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Logpush.Jobs.Get(
		context.TODO(),
		int64(1),
		logpush.JobGetParams{
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
