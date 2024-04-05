// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/pages"
)

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pages.Projects.New(context.TODO(), pages.ProjectNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		BuildConfig: cloudflare.F(pages.ProjectNewParamsBuildConfig{
			BuildCaching:      cloudflare.F(true),
			BuildCommand:      cloudflare.F("npm run build"),
			DestinationDir:    cloudflare.F("build"),
			RootDir:           cloudflare.F("/"),
			WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
			WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
		}),
		CanonicalDeployment: cloudflare.F(pages.DeploymentParam{}),
		DeploymentConfigs: cloudflare.F(pages.ProjectNewParamsDeploymentConfigs{
			Preview: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreview{
				AIBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAIBindings{
					AIBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding{
						ProjectID: cloudflare.F[any](map[string]interface{}{}),
					}),
				}),
				AnalyticsEngineDatasets: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
					AnalyticsEngineBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding{
						Dataset: cloudflare.F("api_analytics"),
					}),
				}),
				Browsers: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewBrowsers{
					Browser: cloudflare.F[any](map[string]interface{}{}),
				}),
				CompatibilityDate:  cloudflare.F("2022-01-01"),
				CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
				D1Databases: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewD1Databases{
					D1Binding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding{
						ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
					}),
				}),
				DurableObjectNamespaces: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces{
					DoBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding{
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					}),
				}),
				EnvVars: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewEnvVars{
					EnvironmentVariable: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable{
						Type:  cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText),
						Value: cloudflare.F("hello world"),
					}),
				}),
				HyperdriveBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings{
					Hyperdrive: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive{
						ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
					}),
				}),
				KVNamespaces: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewKVNamespaces{
					KVBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewKVNamespacesKVBinding{
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					}),
				}),
				MTLSCertificates: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates{
					MTLS: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewMTLSCertificatesMTLS{
						CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
					}),
				}),
				Placement: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewPlacement{
					Mode: cloudflare.F("smart"),
				}),
				QueueProducers: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewQueueProducers{
					QueueProducerBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding{
						Name: cloudflare.F("some-queue"),
					}),
				}),
				R2Buckets: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewR2Buckets{
					R2Binding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding{
						Name: cloudflare.F("some-bucket"),
					}),
				}),
				Services: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewServices{
					ServiceBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewServicesServiceBinding{
						Entrypoint:  cloudflare.F("MyHandler"),
						Environment: cloudflare.F("production"),
						Service:     cloudflare.F("example-worker"),
					}),
				}),
				VectorizeBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings{
					Vectorize: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewVectorizeBindingsVectorize{
						IndexName: cloudflare.F("my_index"),
					}),
				}),
			}),
			Production: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProduction{
				AIBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionAIBindings{
					AIBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding{
						ProjectID: cloudflare.F[any](map[string]interface{}{}),
					}),
				}),
				AnalyticsEngineDatasets: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
					AnalyticsEngineBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding{
						Dataset: cloudflare.F("api_analytics"),
					}),
				}),
				Browsers: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionBrowsers{
					Browser: cloudflare.F[any](map[string]interface{}{}),
				}),
				CompatibilityDate:  cloudflare.F("2022-01-01"),
				CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
				D1Databases: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionD1Databases{
					D1Binding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding{
						ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
					}),
				}),
				DurableObjectNamespaces: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces{
					DoBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding{
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					}),
				}),
				EnvVars: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionEnvVars{
					EnvironmentVariable: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable{
						Type:  cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText),
						Value: cloudflare.F("hello world"),
					}),
				}),
				HyperdriveBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings{
					Hyperdrive: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive{
						ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
					}),
				}),
				KVNamespaces: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionKVNamespaces{
					KVBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionKVNamespacesKVBinding{
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					}),
				}),
				MTLSCertificates: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionMTLSCertificates{
					MTLS: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionMTLSCertificatesMTLS{
						CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
					}),
				}),
				Placement: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionPlacement{
					Mode: cloudflare.F("smart"),
				}),
				QueueProducers: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionQueueProducers{
					QueueProducerBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding{
						Name: cloudflare.F("some-queue"),
					}),
				}),
				R2Buckets: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionR2Buckets{
					R2Binding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding{
						Name: cloudflare.F("some-bucket"),
					}),
				}),
				Services: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionServices{
					ServiceBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionServicesServiceBinding{
						Entrypoint:  cloudflare.F("MyHandler"),
						Environment: cloudflare.F("production"),
						Service:     cloudflare.F("example-worker"),
					}),
				}),
				VectorizeBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionVectorizeBindings{
					Vectorize: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionVectorizeBindingsVectorize{
						IndexName: cloudflare.F("my_index"),
					}),
				}),
			}),
		}),
		LatestDeployment: cloudflare.F(pages.DeploymentParam{}),
		Name:             cloudflare.F("NextJS Blog"),
		ProductionBranch: cloudflare.F("main"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectList(t *testing.T) {
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
	_, err := client.Pages.Projects.List(context.TODO(), pages.ProjectListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectDelete(t *testing.T) {
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
	_, err := client.Pages.Projects.Delete(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:      cloudflare.F[any](map[string]interface{}{}),
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

func TestProjectEdit(t *testing.T) {
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
	_, err := client.Pages.Projects.Edit(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F[any](map[string]interface{}{
				"deployment_configs": map[string]interface{}{
					"production": map[string]interface{}{
						"compatibility_date": "2022-01-01",
						"compatibility_flags": map[string]interface{}{
							"0": "url_standard",
						},
						"env_vars": map[string]interface{}{
							"BUILD_VERSION": map[string]interface{}{
								"value": "3.3",
							},
							"delete_this_env_var": nil,
							"secret_var": map[string]interface{}{
								"type":  "secret_text",
								"value": "A_CMS_API_TOKEN",
							},
						},
					},
				},
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

func TestProjectGet(t *testing.T) {
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
	_, err := client.Pages.Projects.Get(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectGetParams{
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

func TestProjectPurgeBuildCache(t *testing.T) {
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
	_, err := client.Pages.Projects.PurgeBuildCache(
		context.TODO(),
		"this-is-my-project-01",
		pages.ProjectPurgeBuildCacheParams{
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
