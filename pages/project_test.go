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
		Project: pages.ProjectParam{
			BuildConfig: cloudflare.F(pages.ProjectBuildConfigParam{
				BuildCaching:      cloudflare.F(true),
				BuildCommand:      cloudflare.F("npm run build"),
				DestinationDir:    cloudflare.F("build"),
				RootDir:           cloudflare.F("/"),
				WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
				WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
			}),
			CanonicalDeployment: cloudflare.F(pages.DeploymentParam{}),
			DeploymentConfigs: cloudflare.F(pages.ProjectDeploymentConfigsParam{
				Preview: cloudflare.F(pages.ProjectDeploymentConfigsPreviewParam{
					AIBindings: cloudflare.F(pages.ProjectDeploymentConfigsPreviewAIBindingsParam{
						AIBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewAIBindingsAIBindingParam{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsParam{
						AnalyticsEngineBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingParam{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					Browsers: cloudflare.F(pages.ProjectDeploymentConfigsPreviewBrowsersParam{
						Browser: cloudflare.F[any](map[string]interface{}{}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(pages.ProjectDeploymentConfigsPreviewD1DatabasesParam{
						D1Binding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewD1DatabasesD1BindingParam{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(pages.ProjectDeploymentConfigsPreviewDurableObjectNamespacesParam{
						DoBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingParam{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(pages.ProjectDeploymentConfigsPreviewEnvVarsParam{
						EnvironmentVariable: cloudflare.F(pages.ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableParam{
							Type:  cloudflare.F(pages.ProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					HyperdriveBindings: cloudflare.F(pages.ProjectDeploymentConfigsPreviewHyperdriveBindingsParam{
						Hyperdrive: cloudflare.F(pages.ProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveParam{
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						}),
					}),
					KVNamespaces: cloudflare.F(pages.ProjectDeploymentConfigsPreviewKVNamespacesParam{
						KVBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewKVNamespacesKVBindingParam{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					MTLSCertificates: cloudflare.F(pages.ProjectDeploymentConfigsPreviewMTLSCertificatesParam{
						MTLS: cloudflare.F(pages.ProjectDeploymentConfigsPreviewMTLSCertificatesMTLSParam{
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						}),
					}),
					Placement: cloudflare.F(pages.ProjectDeploymentConfigsPreviewPlacementParam{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(pages.ProjectDeploymentConfigsPreviewQueueProducersParam{
						QueueProducerBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingParam{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(pages.ProjectDeploymentConfigsPreviewR2BucketsParam{
						R2Binding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewR2BucketsR2BindingParam{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					Services: cloudflare.F(pages.ProjectDeploymentConfigsPreviewServicesParam{
						ServiceBinding: cloudflare.F(pages.ProjectDeploymentConfigsPreviewServicesServiceBindingParam{
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
					VectorizeBindings: cloudflare.F(pages.ProjectDeploymentConfigsPreviewVectorizeBindingsParam{
						Vectorize: cloudflare.F(pages.ProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeParam{
							IndexName: cloudflare.F("my_index"),
						}),
					}),
				}),
				Production: cloudflare.F(pages.ProjectDeploymentConfigsProductionParam{
					AIBindings: cloudflare.F(pages.ProjectDeploymentConfigsProductionAIBindingsParam{
						AIBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionAIBindingsAIBindingParam{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsParam{
						AnalyticsEngineBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingParam{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					Browsers: cloudflare.F(pages.ProjectDeploymentConfigsProductionBrowsersParam{
						Browser: cloudflare.F[any](map[string]interface{}{}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(pages.ProjectDeploymentConfigsProductionD1DatabasesParam{
						D1Binding: cloudflare.F(pages.ProjectDeploymentConfigsProductionD1DatabasesD1BindingParam{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(pages.ProjectDeploymentConfigsProductionDurableObjectNamespacesParam{
						DoBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingParam{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsParam{
						EnvironmentVariable: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableParam{
							Type:  cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					HyperdriveBindings: cloudflare.F(pages.ProjectDeploymentConfigsProductionHyperdriveBindingsParam{
						Hyperdrive: cloudflare.F(pages.ProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveParam{
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						}),
					}),
					KVNamespaces: cloudflare.F(pages.ProjectDeploymentConfigsProductionKVNamespacesParam{
						KVBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionKVNamespacesKVBindingParam{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					MTLSCertificates: cloudflare.F(pages.ProjectDeploymentConfigsProductionMTLSCertificatesParam{
						MTLS: cloudflare.F(pages.ProjectDeploymentConfigsProductionMTLSCertificatesMTLSParam{
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						}),
					}),
					Placement: cloudflare.F(pages.ProjectDeploymentConfigsProductionPlacementParam{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(pages.ProjectDeploymentConfigsProductionQueueProducersParam{
						QueueProducerBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingParam{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(pages.ProjectDeploymentConfigsProductionR2BucketsParam{
						R2Binding: cloudflare.F(pages.ProjectDeploymentConfigsProductionR2BucketsR2BindingParam{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					Services: cloudflare.F(pages.ProjectDeploymentConfigsProductionServicesParam{
						ServiceBinding: cloudflare.F(pages.ProjectDeploymentConfigsProductionServicesServiceBindingParam{
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
					VectorizeBindings: cloudflare.F(pages.ProjectDeploymentConfigsProductionVectorizeBindingsParam{
						Vectorize: cloudflare.F(pages.ProjectDeploymentConfigsProductionVectorizeBindingsVectorizeParam{
							IndexName: cloudflare.F("my_index"),
						}),
					}),
				}),
			}),
			LatestDeployment: cloudflare.F(pages.DeploymentParam{}),
			Name:             cloudflare.F("NextJS Blog"),
			ProductionBranch: cloudflare.F("main"),
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

func TestProjectList(t *testing.T) {
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
			Body:      map[string]interface{}{},
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
			Body: map[string]interface{}{
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
			},
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
