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

func TestStreamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Streams.New(context.TODO(), pipelines.StreamNewParams{
		AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("my_stream"),
		Format: cloudflare.F[pipelines.StreamNewParamsFormatUnion](pipelines.StreamNewParamsFormatJson{
			Type:            cloudflare.F(pipelines.StreamNewParamsFormatJsonTypeJson),
			DecimalEncoding: cloudflare.F(pipelines.StreamNewParamsFormatJsonDecimalEncodingNumber),
			TimestampFormat: cloudflare.F(pipelines.StreamNewParamsFormatJsonTimestampFormatRfc3339),
			Unstructured:    cloudflare.F(true),
		}),
		HTTP: cloudflare.F(pipelines.StreamNewParamsHTTP{
			Authentication: cloudflare.F(false),
			Enabled:        cloudflare.F(true),
			CORS: cloudflare.F(pipelines.StreamNewParamsHTTPCORS{
				Origins: cloudflare.F([]string{"string"}),
			}),
		}),
		Schema: cloudflare.F(pipelines.StreamNewParamsSchema{
			Fields: cloudflare.F([]pipelines.StreamNewParamsSchemaFieldUnion{pipelines.StreamNewParamsSchemaFieldsInt32{
				Type:        cloudflare.F(pipelines.StreamNewParamsSchemaFieldsInt32TypeInt32),
				MetadataKey: cloudflare.F("metadata_key"),
				Name:        cloudflare.F("name"),
				Required:    cloudflare.F(true),
				SqlName:     cloudflare.F("sql_name"),
			}}),
			Format: cloudflare.F[pipelines.StreamNewParamsSchemaFormatUnion](pipelines.StreamNewParamsSchemaFormatJson{
				Type:            cloudflare.F(pipelines.StreamNewParamsSchemaFormatJsonTypeJson),
				DecimalEncoding: cloudflare.F(pipelines.StreamNewParamsSchemaFormatJsonDecimalEncodingNumber),
				TimestampFormat: cloudflare.F(pipelines.StreamNewParamsSchemaFormatJsonTimestampFormatRfc3339),
				Unstructured:    cloudflare.F(true),
			}),
			Inferred: cloudflare.F(true),
		}),
		WorkerBinding: cloudflare.F(pipelines.StreamNewParamsWorkerBinding{
			Enabled: cloudflare.F(true),
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

func TestStreamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Streams.Update(
		context.TODO(),
		"033e105f4ecef8ad9ca31a8372d0c353",
		pipelines.StreamUpdateParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
			HTTP: cloudflare.F(pipelines.StreamUpdateParamsHTTP{
				Authentication: cloudflare.F(false),
				Enabled:        cloudflare.F(true),
				CORS: cloudflare.F(pipelines.StreamUpdateParamsHTTPCORS{
					Origins: cloudflare.F([]string{"string"}),
				}),
			}),
			WorkerBinding: cloudflare.F(pipelines.StreamUpdateParamsWorkerBinding{
				Enabled: cloudflare.F(true),
			}),
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

func TestStreamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Streams.List(context.TODO(), pipelines.StreamListParams{
		AccountID:  cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Page:       cloudflare.F(0.000000),
		PerPage:    cloudflare.F(0.000000),
		PipelineID: cloudflare.F("043e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStreamDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Pipelines.Streams.Delete(
		context.TODO(),
		"033e105f4ecef8ad9ca31a8372d0c353",
		pipelines.StreamDeleteParams{
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

func TestStreamGet(t *testing.T) {
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
	_, err := client.Pipelines.Streams.Get(
		context.TODO(),
		"033e105f4ecef8ad9ca31a8372d0c353",
		pipelines.StreamGetParams{
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
