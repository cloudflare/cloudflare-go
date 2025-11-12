// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipelines_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pipelines"
)

func TestSinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Sinks.New(context.TODO(), pipelines.SinkNewParams{
		AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("my_sink"),
		Type:      cloudflare.F(pipelines.SinkNewParamsTypeR2),
		Config: cloudflare.F[pipelines.SinkNewParamsConfigUnion](pipelines.SinkNewParamsConfigCloudflarePipelinesR2Table{
			AccountID: cloudflare.F("account_id"),
			Bucket:    cloudflare.F("bucket"),
			Credentials: cloudflare.F(pipelines.SinkNewParamsConfigCloudflarePipelinesR2TableCredentials{
				AccessKeyID:     cloudflare.F("access_key_id"),
				SecretAccessKey: cloudflare.F("secret_access_key"),
			}),
			FileNaming: cloudflare.F(pipelines.SinkNewParamsConfigCloudflarePipelinesR2TableFileNaming{
				Prefix:   cloudflare.F("prefix"),
				Strategy: cloudflare.F(pipelines.SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategySerial),
				Suffix:   cloudflare.F("suffix"),
			}),
			Jurisdiction: cloudflare.F("jurisdiction"),
			Partitioning: cloudflare.F(pipelines.SinkNewParamsConfigCloudflarePipelinesR2TablePartitioning{
				TimePattern: cloudflare.F("year=%Y/month=%m/day=%d/hour=%H"),
			}),
			Path: cloudflare.F("path"),
			RollingPolicy: cloudflare.F(pipelines.SinkNewParamsConfigCloudflarePipelinesR2TableRollingPolicy{
				FileSizeBytes:     cloudflare.F(int64(0)),
				InactivitySeconds: cloudflare.F(int64(1)),
				IntervalSeconds:   cloudflare.F(int64(1)),
			}),
		}),
		Format: cloudflare.F[pipelines.SinkNewParamsFormatUnion](pipelines.SinkNewParamsFormatJson{
			Type:            cloudflare.F(pipelines.SinkNewParamsFormatJsonTypeJson),
			DecimalEncoding: cloudflare.F(pipelines.SinkNewParamsFormatJsonDecimalEncodingNumber),
			TimestampFormat: cloudflare.F(pipelines.SinkNewParamsFormatJsonTimestampFormatRfc3339),
			Unstructured:    cloudflare.F(true),
		}),
		Schema: cloudflare.F(pipelines.SinkNewParamsSchema{
			Fields: cloudflare.F([]pipelines.SinkNewParamsSchemaFieldUnion{pipelines.SinkNewParamsSchemaFieldsInt32{
				Type:        cloudflare.F(pipelines.SinkNewParamsSchemaFieldsInt32TypeInt32),
				MetadataKey: cloudflare.F("metadata_key"),
				Name:        cloudflare.F("name"),
				Required:    cloudflare.F(true),
				SqlName:     cloudflare.F("sql_name"),
			}}),
			Format: cloudflare.F[pipelines.SinkNewParamsSchemaFormatUnion](pipelines.SinkNewParamsSchemaFormatJson{
				Type:            cloudflare.F(pipelines.SinkNewParamsSchemaFormatJsonTypeJson),
				DecimalEncoding: cloudflare.F(pipelines.SinkNewParamsSchemaFormatJsonDecimalEncodingNumber),
				TimestampFormat: cloudflare.F(pipelines.SinkNewParamsSchemaFormatJsonTimestampFormatRfc3339),
				Unstructured:    cloudflare.F(true),
			}),
			Inferred: cloudflare.F(true),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Sinks.List(context.TODO(), pipelines.SinkListParams{
		AccountID:  cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Page:       cloudflare.F(0.000000),
		PerPage:    cloudflare.F(0.000000),
		PipelineID: cloudflare.F("pipeline_id"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSinkDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Pipelines.Sinks.Delete(
		context.TODO(),
		"0223105f4ecef8ad9ca31a8372d0c353",
		pipelines.SinkDeleteParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
			Force:     cloudflare.F("force"),
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

func TestSinkGet(t *testing.T) {
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
	_, err := client.Pipelines.Sinks.Get(
		context.TODO(),
		"0223105f4ecef8ad9ca31a8372d0c353",
		pipelines.SinkGetParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
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
