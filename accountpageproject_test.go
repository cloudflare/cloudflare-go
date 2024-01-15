// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountPageProjectGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Pages.Projects.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPageProjectUpdate(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Pages.Projects.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
		cloudflare.AccountPageProjectUpdateParams{
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

func TestAccountPageProjectDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Pages.Projects.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPageProjectPagesProjectNewProjectWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Pages.Projects.PagesProjectNewProject(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountPageProjectPagesProjectNewProjectParams{
			BuildConfig: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsBuildConfig{
				BuildCaching:      cloudflare.F(true),
				BuildCommand:      cloudflare.F("npm run build"),
				DestinationDir:    cloudflare.F("build"),
				RootDir:           cloudflare.F("/"),
				WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
				WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
			}),
			CanonicalDeployment: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsCanonicalDeployment{}),
			DeploymentConfigs: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigs{
				Preview: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreview{
					AIBindings: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindings{
						AIBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAIBindingsAIBinding{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1Databases{
						D1Binding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespaces{
						DoBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVars{
						EnvironmentVariable: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					KvNamespaces: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespaces{
						KvBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewKvNamespacesKvBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					Placement: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducers{
						QueueProducerBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2Buckets{
						R2Binding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewR2BucketsR2Binding{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					ServiceBindings: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindings{
						ServiceBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsPreviewServiceBindingsServiceBinding{
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
				}),
				Production: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProduction{
					AIBindings: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindings{
						AIBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAIBindingsAIBinding{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1Databases{
						D1Binding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespaces{
						DoBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVars{
						EnvironmentVariable: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					KvNamespaces: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespaces{
						KvBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionKvNamespacesKvBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					Placement: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducers{
						QueueProducerBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2Buckets{
						R2Binding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionR2BucketsR2Binding{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					ServiceBindings: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindings{
						ServiceBinding: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsDeploymentConfigsProductionServiceBindingsServiceBinding{
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
				}),
			}),
			LatestDeployment: cloudflare.F(cloudflare.AccountPageProjectPagesProjectNewProjectParamsLatestDeployment{}),
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

func TestAccountPageProjectPagesProjectGetProjects(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Pages.Projects.PagesProjectGetProjects(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
