// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pages"
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
		AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:             cloudflare.F("my-pages-app"),
		ProductionBranch: cloudflare.F("main"),
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
				AIBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewAIBindings{
					"AI_BINDING": {
						ProjectID: cloudflare.F("some-project-id"),
					},
				}),
				AlwaysUseLatestCompatibilityDate: cloudflare.F(false),
				AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
					"ANALYTICS_ENGINE_BINDING": {
						Dataset: cloudflare.F("api_analytics"),
					},
				}),
				Browsers: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewBrowsers{
					"BROWSER": {},
				}),
				BuildImageMajorVersion: cloudflare.F(int64(3)),
				CompatibilityDate:      cloudflare.F("2025-01-01"),
				CompatibilityFlags:     cloudflare.F([]string{"url_standard"}),
				D1Databases: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewD1Databases{
					"D1_BINDING": {
						ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
					},
				}),
				DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces{
					"DO_BINDING": {
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					},
				}),
				EnvVars: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion{
					"foo": pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{
						Type:  cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText),
						Value: cloudflare.F("hello world"),
					},
				}),
				FailOpen: cloudflare.F(true),
				HyperdriveBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings{
					"HYPERDRIVE": {
						ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
					},
				}),
				KVNamespaces: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewKVNamespaces{
					"KV_BINDING": {
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					},
				}),
				Limits: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewLimits{
					CPUMs: cloudflare.F(int64(100)),
				}),
				MTLSCertificates: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates{
					"MTLS": {
						CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
					},
				}),
				Placement: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewPlacement{
					Mode: cloudflare.F("smart"),
				}),
				QueueProducers: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewQueueProducers{
					"QUEUE_PRODUCER_BINDING": {
						Name: cloudflare.F("some-queue"),
					},
				}),
				R2Buckets: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewR2Buckets{
					"R2_BINDING": {
						Name:         cloudflare.F("some-bucket"),
						Jurisdiction: cloudflare.F("eu"),
					},
				}),
				Services: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewServices{
					"SERVICE_BINDING": {
						Service:     cloudflare.F("example-worker"),
						Entrypoint:  cloudflare.F("MyHandler"),
						Environment: cloudflare.F("production"),
					},
				}),
				UsageModel: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsPreviewUsageModelStandard),
				VectorizeBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings{
					"VECTORIZE": {
						IndexName: cloudflare.F("my_index"),
					},
				}),
				WranglerConfigHash: cloudflare.F("abc123def456"),
			}),
			Production: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProduction{
				AIBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionAIBindings{
					"AI_BINDING": {
						ProjectID: cloudflare.F("some-project-id"),
					},
				}),
				AlwaysUseLatestCompatibilityDate: cloudflare.F(false),
				AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
					"ANALYTICS_ENGINE_BINDING": {
						Dataset: cloudflare.F("api_analytics"),
					},
				}),
				Browsers: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionBrowsers{
					"BROWSER": {},
				}),
				BuildImageMajorVersion: cloudflare.F(int64(3)),
				CompatibilityDate:      cloudflare.F("2025-01-01"),
				CompatibilityFlags:     cloudflare.F([]string{"url_standard"}),
				D1Databases: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionD1Databases{
					"D1_BINDING": {
						ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
					},
				}),
				DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces{
					"DO_BINDING": {
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					},
				}),
				EnvVars: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsUnion{
					"foo": pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{
						Type:  cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText),
						Value: cloudflare.F("hello world"),
					},
				}),
				FailOpen: cloudflare.F(true),
				HyperdriveBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings{
					"HYPERDRIVE": {
						ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
					},
				}),
				KVNamespaces: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionKVNamespaces{
					"KV_BINDING": {
						NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
					},
				}),
				Limits: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionLimits{
					CPUMs: cloudflare.F(int64(100)),
				}),
				MTLSCertificates: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionMTLSCertificates{
					"MTLS": {
						CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
					},
				}),
				Placement: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionPlacement{
					Mode: cloudflare.F("smart"),
				}),
				QueueProducers: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionQueueProducers{
					"QUEUE_PRODUCER_BINDING": {
						Name: cloudflare.F("some-queue"),
					},
				}),
				R2Buckets: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionR2Buckets{
					"R2_BINDING": {
						Name:         cloudflare.F("some-bucket"),
						Jurisdiction: cloudflare.F("eu"),
					},
				}),
				Services: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionServices{
					"SERVICE_BINDING": {
						Service:     cloudflare.F("example-worker"),
						Entrypoint:  cloudflare.F("MyHandler"),
						Environment: cloudflare.F("production"),
					},
				}),
				UsageModel: cloudflare.F(pages.ProjectNewParamsDeploymentConfigsProductionUsageModelStandard),
				VectorizeBindings: cloudflare.F(map[string]pages.ProjectNewParamsDeploymentConfigsProductionVectorizeBindings{
					"VECTORIZE": {
						IndexName: cloudflare.F("my_index"),
					},
				}),
				WranglerConfigHash: cloudflare.F("abc123def456"),
			}),
		}),
		Source: cloudflare.F(pages.ProjectNewParamsSource{
			Config: cloudflare.F(pages.ProjectNewParamsSourceConfig{
				DeploymentsEnabled:           cloudflare.F(true),
				Owner:                        cloudflare.F("my-org"),
				OwnerID:                      cloudflare.F("12345678"),
				PathExcludes:                 cloudflare.F([]string{"string"}),
				PathIncludes:                 cloudflare.F([]string{"string"}),
				PrCommentsEnabled:            cloudflare.F(true),
				PreviewBranchExcludes:        cloudflare.F([]string{"string"}),
				PreviewBranchIncludes:        cloudflare.F([]string{"string"}),
				PreviewDeploymentSetting:     cloudflare.F(pages.ProjectNewParamsSourceConfigPreviewDeploymentSettingAll),
				ProductionBranch:             cloudflare.F("main"),
				ProductionDeploymentsEnabled: cloudflare.F(true),
				RepoID:                       cloudflare.F("12345678"),
				RepoName:                     cloudflare.F("my-repo"),
			}),
			Type: cloudflare.F(pages.ProjectNewParamsSourceTypeGitHub),
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

func TestProjectListWithOptionalParams(t *testing.T) {
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
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(10)),
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
					AIBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewAIBindings{
						"AI_BINDING": {
							ProjectID: cloudflare.F("some-project-id"),
						},
					}),
					AlwaysUseLatestCompatibilityDate: cloudflare.F(false),
					AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
						"ANALYTICS_ENGINE_BINDING": {
							Dataset: cloudflare.F("api_analytics"),
						},
					}),
					Browsers: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewBrowsers{
						"BROWSER": {},
					}),
					BuildImageMajorVersion: cloudflare.F(int64(3)),
					CompatibilityDate:      cloudflare.F("2025-01-01"),
					CompatibilityFlags:     cloudflare.F([]string{"url_standard"}),
					D1Databases: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewD1Databases{
						"D1_BINDING": {
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						},
					}),
					DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces{
						"DO_BINDING": {
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					EnvVars: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion{
						"foo": pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{
							Type:  cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText),
							Value: cloudflare.F("hello world"),
						},
					}),
					FailOpen: cloudflare.F(true),
					HyperdriveBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings{
						"HYPERDRIVE": {
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						},
					}),
					KVNamespaces: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewKVNamespaces{
						"KV_BINDING": {
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					Limits: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewLimits{
						CPUMs: cloudflare.F(int64(100)),
					}),
					MTLSCertificates: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates{
						"MTLS": {
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						},
					}),
					Placement: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewQueueProducers{
						"QUEUE_PRODUCER_BINDING": {
							Name: cloudflare.F("some-queue"),
						},
					}),
					R2Buckets: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewR2Buckets{
						"R2_BINDING": {
							Name:         cloudflare.F("some-bucket"),
							Jurisdiction: cloudflare.F("eu"),
						},
					}),
					Services: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewServices{
						"SERVICE_BINDING": {
							Service:     cloudflare.F("example-worker"),
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
						},
					}),
					UsageModel: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsPreviewUsageModelStandard),
					VectorizeBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings{
						"VECTORIZE": {
							IndexName: cloudflare.F("my_index"),
						},
					}),
					WranglerConfigHash: cloudflare.F("abc123def456"),
				}),
				Production: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProduction{
					AIBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionAIBindings{
						"AI_BINDING": {
							ProjectID: cloudflare.F("some-project-id"),
						},
					}),
					AlwaysUseLatestCompatibilityDate: cloudflare.F(false),
					AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
						"ANALYTICS_ENGINE_BINDING": {
							Dataset: cloudflare.F("api_analytics"),
						},
					}),
					Browsers: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionBrowsers{
						"BROWSER": {},
					}),
					BuildImageMajorVersion: cloudflare.F(int64(3)),
					CompatibilityDate:      cloudflare.F("2025-01-01"),
					CompatibilityFlags:     cloudflare.F([]string{"url_standard"}),
					D1Databases: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionD1Databases{
						"D1_BINDING": {
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						},
					}),
					DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces{
						"DO_BINDING": {
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					EnvVars: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsUnion{
						"foo": pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{
							Type:  cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText),
							Value: cloudflare.F("hello world"),
						},
					}),
					FailOpen: cloudflare.F(true),
					HyperdriveBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings{
						"HYPERDRIVE": {
							ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
						},
					}),
					KVNamespaces: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionKVNamespaces{
						"KV_BINDING": {
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						},
					}),
					Limits: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionLimits{
						CPUMs: cloudflare.F(int64(100)),
					}),
					MTLSCertificates: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionMTLSCertificates{
						"MTLS": {
							CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
						},
					}),
					Placement: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionQueueProducers{
						"QUEUE_PRODUCER_BINDING": {
							Name: cloudflare.F("some-queue"),
						},
					}),
					R2Buckets: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionR2Buckets{
						"R2_BINDING": {
							Name:         cloudflare.F("some-bucket"),
							Jurisdiction: cloudflare.F("eu"),
						},
					}),
					Services: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionServices{
						"SERVICE_BINDING": {
							Service:     cloudflare.F("example-worker"),
							Entrypoint:  cloudflare.F("MyHandler"),
							Environment: cloudflare.F("production"),
						},
					}),
					UsageModel: cloudflare.F(pages.ProjectEditParamsDeploymentConfigsProductionUsageModelStandard),
					VectorizeBindings: cloudflare.F(map[string]pages.ProjectEditParamsDeploymentConfigsProductionVectorizeBindings{
						"VECTORIZE": {
							IndexName: cloudflare.F("my_index"),
						},
					}),
					WranglerConfigHash: cloudflare.F("abc123def456"),
				}),
			}),
			Name:             cloudflare.F("my-pages-app"),
			ProductionBranch: cloudflare.F("main"),
			Source: cloudflare.F(pages.ProjectEditParamsSource{
				Config: cloudflare.F(pages.ProjectEditParamsSourceConfig{
					DeploymentsEnabled:           cloudflare.F(true),
					Owner:                        cloudflare.F("my-org"),
					OwnerID:                      cloudflare.F("12345678"),
					PathExcludes:                 cloudflare.F([]string{"string"}),
					PathIncludes:                 cloudflare.F([]string{"string"}),
					PrCommentsEnabled:            cloudflare.F(true),
					PreviewBranchExcludes:        cloudflare.F([]string{"string"}),
					PreviewBranchIncludes:        cloudflare.F([]string{"string"}),
					PreviewDeploymentSetting:     cloudflare.F(pages.ProjectEditParamsSourceConfigPreviewDeploymentSettingAll),
					ProductionBranch:             cloudflare.F("main"),
					ProductionDeploymentsEnabled: cloudflare.F(true),
					RepoID:                       cloudflare.F("12345678"),
					RepoName:                     cloudflare.F("my-repo"),
				}),
				Type: cloudflare.F(pages.ProjectEditParamsSourceTypeGitHub),
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
