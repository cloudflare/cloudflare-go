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
		BuildConfig: cloudflare.F(pages.ProjectNewParamsBuildConfig{
			BuildCaching:      cloudflare.F(true),
			BuildCommand:      cloudflare.F("npm run build"),
			DestinationDir:    cloudflare.F("build"),
			RootDir:           cloudflare.F("/"),
			WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
			WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
		}),
		DeploymentConfigs: cloudflare.F(pages.ProjectNewParamsDeploymentConfigs{
			Preview: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreview{
				AIBindings: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAIBindings{
					AIBinding: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding{
						ProjectID: cloudflare.F("project_id"),
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
				CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
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
						Jurisdiction: cloudflare.F("eu"),
						Name:         cloudflare.F("some-bucket"),
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
						ProjectID: cloudflare.F("project_id"),
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
				CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
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
						Jurisdiction: cloudflare.F("eu"),
						Name:         cloudflare.F("some-bucket"),
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

func TestProjectEditWithOptionalParams(t *testing.T) {
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
			BuildConfig: cloudflare.F(pages.ProjectEditParamsBuildConfig{
				BuildCaching:      cloudflare.F(true),
				BuildCommand:      cloudflare.F("npm run build"),
				DestinationDir:    cloudflare.F("build"),
				RootDir:           cloudflare.F("/"),
				WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
				WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
			}),
			DeploymentConfigs: cloudflare.F(pages.ProjectEditParamsDeploymentConfigs{
				Preview: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreview{
					AIBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewAIBindings{
						AIBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewAIBindingsAIBinding{
							ProjectID: cloudflare.F("project_id"),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					Browsers: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewBrowsers{
						Browser: cloudflare.F[any](map[string]interface{}{}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
					D1Databases: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewD1Databases{
						D1Binding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces{
						DoBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewEnvVars{
						EnvironmentVariable: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					HyperdriveBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings{
						Hyperdrive: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindingsHyperdrive{
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						}),
					}),
					KVNamespaces: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewKVNamespaces{
						KVBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewKVNamespacesKVBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					MTLSCertificates: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates{
						MTLS: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewMTLSCertificatesMTLS{
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						}),
					}),
					Placement: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewQueueProducers{
						QueueProducerBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewR2Buckets{
						R2Binding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewR2BucketsR2Binding{
							Jurisdiction: cloudflare.F("eu"),
							Name:         cloudflare.F("some-bucket"),
						}),
					}),
					Services: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewServices{
						ServiceBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewServicesServiceBinding{
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
					VectorizeBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings{
						Vectorize: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewVectorizeBindingsVectorize{
							IndexName: cloudflare.F("my_index"),
						}),
					}),
				}),
				Production: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProduction{
					AIBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionAIBindings{
						AIBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionAIBindingsAIBinding{
							ProjectID: cloudflare.F("project_id"),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					Browsers: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionBrowsers{
						Browser: cloudflare.F[any](map[string]interface{}{}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
					D1Databases: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionD1Databases{
						D1Binding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces{
						DoBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionEnvVars{
						EnvironmentVariable: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					HyperdriveBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings{
						Hyperdrive: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionHyperdriveBindingsHyperdrive{
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						}),
					}),
					KVNamespaces: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionKVNamespaces{
						KVBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionKVNamespacesKVBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					MTLSCertificates: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionMTLSCertificates{
						MTLS: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionMTLSCertificatesMTLS{
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						}),
					}),
					Placement: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionQueueProducers{
						QueueProducerBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionR2Buckets{
						R2Binding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionR2BucketsR2Binding{
							Jurisdiction: cloudflare.F("eu"),
							Name:         cloudflare.F("some-bucket"),
						}),
					}),
					Services: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionServices{
						ServiceBinding: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionServicesServiceBinding{
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
					VectorizeBindings: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionVectorizeBindings{
						Vectorize: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionVectorizeBindingsVectorize{
							IndexName: cloudflare.F("my_index"),
						}),
					}),
				}),
			}),
			Name:             cloudflare.F("NextJS Blog"),
			ProductionBranch: cloudflare.F("main"),
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
