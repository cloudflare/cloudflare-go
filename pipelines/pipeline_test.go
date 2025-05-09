// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipelines_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/pipelines"
)

func TestPipelineNew(t *testing.T) {
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
	_, err := client.Pipelines.New(context.TODO(), pipelines.PipelineNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Destination: cloudflare.F(pipelines.PipelineNewParamsDestination{
			Batch: cloudflare.F(pipelines.PipelineNewParamsDestinationBatch{
				MaxBytes:     cloudflare.F(int64(1000)),
				MaxDurationS: cloudflare.F(0.250000),
				MaxRows:      cloudflare.F(int64(100)),
			}),
			Compression: cloudflare.F(pipelines.PipelineNewParamsDestinationCompression{
				Type: cloudflare.F(pipelines.PipelineNewParamsDestinationCompressionTypeGzip),
			}),
			Credentials: cloudflare.F(pipelines.PipelineNewParamsDestinationCredentials{
				AccessKeyID:     cloudflare.F("<access key id>"),
				Endpoint:        cloudflare.F("https://123f8a8258064ed892a347f173372359.r2.cloudflarestorage.com"),
				SecretAccessKey: cloudflare.F("<secret key>"),
			}),
			Format: cloudflare.F(pipelines.PipelineNewParamsDestinationFormatJson),
			Path: cloudflare.F(pipelines.PipelineNewParamsDestinationPath{
				Bucket:   cloudflare.F("bucket"),
				Filename: cloudflare.F("${slug}${extension}"),
				Filepath: cloudflare.F("${date}/${hour}"),
				Prefix:   cloudflare.F("base"),
			}),
			Type: cloudflare.F(pipelines.PipelineNewParamsDestinationTypeR2),
		}),
		Name: cloudflare.F("sample_pipeline"),
		Source: cloudflare.F([]pipelines.PipelineNewParamsSourceUnion{pipelines.PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource{
			Format:         cloudflare.F(pipelines.PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson),
			Type:           cloudflare.F("type"),
			Authentication: cloudflare.F(true),
			CORS: cloudflare.F(pipelines.PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS{
				Origins: cloudflare.F([]string{"*"}),
			}),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Update(
		context.TODO(),
		"sample_pipeline",
		pipelines.PipelineUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Destination: cloudflare.F(pipelines.PipelineUpdateParamsDestination{
				Batch: cloudflare.F(pipelines.PipelineUpdateParamsDestinationBatch{
					MaxBytes:     cloudflare.F(int64(1000)),
					MaxDurationS: cloudflare.F(0.250000),
					MaxRows:      cloudflare.F(int64(100)),
				}),
				Compression: cloudflare.F(pipelines.PipelineUpdateParamsDestinationCompression{
					Type: cloudflare.F(pipelines.PipelineUpdateParamsDestinationCompressionTypeGzip),
				}),
				Format: cloudflare.F(pipelines.PipelineUpdateParamsDestinationFormatJson),
				Path: cloudflare.F(pipelines.PipelineUpdateParamsDestinationPath{
					Bucket:   cloudflare.F("bucket"),
					Filename: cloudflare.F("${slug}${extension}"),
					Filepath: cloudflare.F("${date}/${hour}"),
					Prefix:   cloudflare.F("base"),
				}),
				Type: cloudflare.F(pipelines.PipelineUpdateParamsDestinationTypeR2),
				Credentials: cloudflare.F(pipelines.PipelineUpdateParamsDestinationCredentials{
					AccessKeyID:     cloudflare.F("<access key id>"),
					Endpoint:        cloudflare.F("https://123f8a8258064ed892a347f173372359.r2.cloudflarestorage.com"),
					SecretAccessKey: cloudflare.F("<secret key>"),
				}),
			}),
			Name: cloudflare.F("sample_pipeline"),
			Source: cloudflare.F([]pipelines.PipelineUpdateParamsSourceUnion{pipelines.PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource{
				Format:         cloudflare.F(pipelines.PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson),
				Type:           cloudflare.F("type"),
				Authentication: cloudflare.F(true),
				CORS: cloudflare.F(pipelines.PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS{
					Origins: cloudflare.F([]string{"*"}),
				}),
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

func TestPipelineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.List(context.TODO(), pipelines.PipelineListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F("page"),
		PerPage:   cloudflare.F("per_page"),
		Search:    cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineDelete(t *testing.T) {
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
	err := client.Pipelines.Delete(
		context.TODO(),
		"sample_pipeline",
		pipelines.PipelineDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestPipelineGet(t *testing.T) {
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
	_, err := client.Pipelines.Get(
		context.TODO(),
		"sample_pipeline",
		pipelines.PipelineGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
