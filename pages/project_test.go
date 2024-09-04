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
        BuildCaching: cloudflare.F(true),
        BuildCommand: cloudflare.F("npm run build"),
        DestinationDir: cloudflare.F("build"),
        RootDir: cloudflare.F("/"),
        WebAnalyticsTag: cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
        WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
      }),
      DeploymentConfigs: cloudflare.F(pages.ProjectDeploymentConfigsParam{
        Preview: cloudflare.F(pages.ProjectDeploymentConfigsPreviewParam{
          AIBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewAIBindingParam{
            "AI_BINDING": pages.ProjectDeploymentConfigsPreviewAIBindingParam{
              ProjectID: cloudflare.F("some-project-id"),
            },
          }),
          AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
            "ANALYTICS_ENGINE_BINDING": pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
              Dataset: cloudflare.F("api_analytics"),
            },
          }),
          Browsers: cloudflare.F(map[string]interface{}{
            "BROWSER": "bar",
          }),
          CompatibilityDate: cloudflare.F("2022-01-01"),
          CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
          D1Databases: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
            "D1_BINDING": pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
              ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
            },
          }),
          DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
            "DO_BINDING": pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
              NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
            },
          }),
          EnvVars: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewEnvVarParam{
            "foo": pages.ProjectDeploymentConfigsPreviewEnvVarParam{
              Value: cloudflare.F("hello world"),
              Type: cloudflare.F(pages.ProjectDeploymentConfigsPreviewEnvVarsTypePlainText),
            },
          }),
          HyperdriveBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
            "HYPERDRIVE": pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
              ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
            },
          }),
          KVNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
            "KV_BINDING": pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
              NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
            },
          }),
          MTLSCertificates: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
            "MTLS": pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
              CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
            },
          }),
          Placement: cloudflare.F(pages.ProjectDeploymentConfigsPreviewPlacementParam{
            Mode: cloudflare.F("smart"),
          }),
          QueueProducers: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
            "QUEUE_PRODUCER_BINDING": pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
              Name: cloudflare.F("some-queue"),
            },
          }),
          R2Buckets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewR2BucketParam{
            "R2_BINDING": pages.ProjectDeploymentConfigsPreviewR2BucketParam{
              Jurisdiction: cloudflare.F("eu"),
              Name: cloudflare.F("some-bucket"),
            },
          }),
          Services: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewServiceParam{
            "SERVICE_BINDING": pages.ProjectDeploymentConfigsPreviewServiceParam{
              Entrypoint: cloudflare.F("MyHandler"),
              Environment: cloudflare.F("production"),
              Service: cloudflare.F("example-worker"),
            },
          }),
          VectorizeBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
            "VECTORIZE": pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
              IndexName: cloudflare.F("my_index"),
            },
          }),
        }),
        Production: cloudflare.F(pages.ProjectDeploymentConfigsProductionParam{
          AIBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionAIBindingParam{
            "AI_BINDING": pages.ProjectDeploymentConfigsProductionAIBindingParam{
              ProjectID: cloudflare.F("some-project-id"),
            },
          }),
          AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
            "ANALYTICS_ENGINE_BINDING": pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
              Dataset: cloudflare.F("api_analytics"),
            },
          }),
          Browsers: cloudflare.F(map[string]interface{}{
            "BROWSER": "bar",
          }),
          CompatibilityDate: cloudflare.F("2022-01-01"),
          CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
          D1Databases: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
            "D1_BINDING": pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
              ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
            },
          }),
          DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
            "DO_BINDING": pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
              NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
            },
          }),
          EnvVars: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionEnvVarParam{
            "foo": pages.ProjectDeploymentConfigsProductionEnvVarParam{
              Value: cloudflare.F("hello world"),
              Type: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
            },
          }),
          HyperdriveBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
            "HYPERDRIVE": pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
              ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
            },
          }),
          KVNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
            "KV_BINDING": pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
              NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
            },
          }),
          MTLSCertificates: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
            "MTLS": pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
              CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
            },
          }),
          Placement: cloudflare.F(pages.ProjectDeploymentConfigsProductionPlacementParam{
            Mode: cloudflare.F("smart"),
          }),
          QueueProducers: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionQueueProducerParam{
            "QUEUE_PRODUCER_BINDING": pages.ProjectDeploymentConfigsProductionQueueProducerParam{
              Name: cloudflare.F("some-queue"),
            },
          }),
          R2Buckets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionR2BucketParam{
            "R2_BINDING": pages.ProjectDeploymentConfigsProductionR2BucketParam{
              Jurisdiction: cloudflare.F("eu"),
              Name: cloudflare.F("some-bucket"),
            },
          }),
          Services: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionServiceParam{
            "SERVICE_BINDING": pages.ProjectDeploymentConfigsProductionServiceParam{
              Entrypoint: cloudflare.F("MyHandler"),
              Environment: cloudflare.F("production"),
              Service: cloudflare.F("example-worker"),
            },
          }),
          VectorizeBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
            "VECTORIZE": pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
              IndexName: cloudflare.F("my_index"),
            },
          }),
        }),
      }),
      Name: cloudflare.F("NextJS Blog"),
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
      Project: pages.ProjectParam{
        BuildConfig: cloudflare.F(pages.ProjectBuildConfigParam{
          BuildCaching: cloudflare.F(true),
          BuildCommand: cloudflare.F("npm run build"),
          DestinationDir: cloudflare.F("build"),
          RootDir: cloudflare.F("/"),
          WebAnalyticsTag: cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
          WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
        }),
        DeploymentConfigs: cloudflare.F(pages.ProjectDeploymentConfigsParam{
          Preview: cloudflare.F(pages.ProjectDeploymentConfigsPreviewParam{
            AIBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewAIBindingParam{
              "AI_BINDING": pages.ProjectDeploymentConfigsPreviewAIBindingParam{
                ProjectID: cloudflare.F("some-project-id"),
              },
            }),
            AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
              "ANALYTICS_ENGINE_BINDING": pages.ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam{
                Dataset: cloudflare.F("api_analytics"),
              },
            }),
            Browsers: cloudflare.F(map[string]interface{}{
              "BROWSER": "bar",
            }),
            CompatibilityDate: cloudflare.F("2022-01-01"),
            CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
            D1Databases: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
              "D1_BINDING": pages.ProjectDeploymentConfigsPreviewD1DatabaseParam{
                ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
              },
            }),
            DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
              "DO_BINDING": pages.ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam{
                NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
              },
            }),
            EnvVars: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewEnvVarParam{
              "foo": pages.ProjectDeploymentConfigsPreviewEnvVarParam{
                Value: cloudflare.F("hello world"),
                Type: cloudflare.F(pages.ProjectDeploymentConfigsPreviewEnvVarsTypePlainText),
              },
            }),
            HyperdriveBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
              "HYPERDRIVE": pages.ProjectDeploymentConfigsPreviewHyperdriveBindingParam{
                ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
              },
            }),
            KVNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
              "KV_BINDING": pages.ProjectDeploymentConfigsPreviewKVNamespaceParam{
                NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
              },
            }),
            MTLSCertificates: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
              "MTLS": pages.ProjectDeploymentConfigsPreviewMTLSCertificateParam{
                CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
              },
            }),
            Placement: cloudflare.F(pages.ProjectDeploymentConfigsPreviewPlacementParam{
              Mode: cloudflare.F("smart"),
            }),
            QueueProducers: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
              "QUEUE_PRODUCER_BINDING": pages.ProjectDeploymentConfigsPreviewQueueProducerParam{
                Name: cloudflare.F("some-queue"),
              },
            }),
            R2Buckets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewR2BucketParam{
              "R2_BINDING": pages.ProjectDeploymentConfigsPreviewR2BucketParam{
                Jurisdiction: cloudflare.F("eu"),
                Name: cloudflare.F("some-bucket"),
              },
            }),
            Services: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewServiceParam{
              "SERVICE_BINDING": pages.ProjectDeploymentConfigsPreviewServiceParam{
                Entrypoint: cloudflare.F("MyHandler"),
                Environment: cloudflare.F("production"),
                Service: cloudflare.F("example-worker"),
              },
            }),
            VectorizeBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
              "VECTORIZE": pages.ProjectDeploymentConfigsPreviewVectorizeBindingParam{
                IndexName: cloudflare.F("my_index"),
              },
            }),
          }),
          Production: cloudflare.F(pages.ProjectDeploymentConfigsProductionParam{
            AIBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionAIBindingParam{
              "AI_BINDING": pages.ProjectDeploymentConfigsProductionAIBindingParam{
                ProjectID: cloudflare.F("some-project-id"),
              },
            }),
            AnalyticsEngineDatasets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
              "ANALYTICS_ENGINE_BINDING": pages.ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam{
                Dataset: cloudflare.F("api_analytics"),
              },
            }),
            Browsers: cloudflare.F(map[string]interface{}{
              "BROWSER": "bar",
            }),
            CompatibilityDate: cloudflare.F("2022-01-01"),
            CompatibilityFlags: cloudflare.F([]string{"url_standard"}),
            D1Databases: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
              "D1_BINDING": pages.ProjectDeploymentConfigsProductionD1DatabaseParam{
                ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
              },
            }),
            DurableObjectNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
              "DO_BINDING": pages.ProjectDeploymentConfigsProductionDurableObjectNamespaceParam{
                NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
              },
            }),
            EnvVars: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionEnvVarParam{
              "BUILD_VERSION": pages.ProjectDeploymentConfigsProductionEnvVarParam{
                Value: cloudflare.F("3.3"),
                Type: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
              },
              "delete_this_env_var": pages.ProjectDeploymentConfigsProductionEnvVarParam{
                Value: cloudflare.F("value"),
                Type: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
              },
              "secret_var": pages.ProjectDeploymentConfigsProductionEnvVarParam{
                Value: cloudflare.F("A_CMS_API_TOKEN"),
                Type: cloudflare.F(pages.ProjectDeploymentConfigsProductionEnvVarsTypePlainText),
              },
            }),
            HyperdriveBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
              "HYPERDRIVE": pages.ProjectDeploymentConfigsProductionHyperdriveBindingParam{
                ID: cloudflare.F("a76a99bc342644deb02c38d66082262a"),
              },
            }),
            KVNamespaces: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
              "KV_BINDING": pages.ProjectDeploymentConfigsProductionKVNamespaceParam{
                NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
              },
            }),
            MTLSCertificates: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
              "MTLS": pages.ProjectDeploymentConfigsProductionMTLSCertificateParam{
                CertificateID: cloudflare.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
              },
            }),
            Placement: cloudflare.F(pages.ProjectDeploymentConfigsProductionPlacementParam{
              Mode: cloudflare.F("smart"),
            }),
            QueueProducers: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionQueueProducerParam{
              "QUEUE_PRODUCER_BINDING": pages.ProjectDeploymentConfigsProductionQueueProducerParam{
                Name: cloudflare.F("some-queue"),
              },
            }),
            R2Buckets: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionR2BucketParam{
              "R2_BINDING": pages.ProjectDeploymentConfigsProductionR2BucketParam{
                Jurisdiction: cloudflare.F("eu"),
                Name: cloudflare.F("some-bucket"),
              },
            }),
            Services: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionServiceParam{
              "SERVICE_BINDING": pages.ProjectDeploymentConfigsProductionServiceParam{
                Entrypoint: cloudflare.F("MyHandler"),
                Environment: cloudflare.F("production"),
                Service: cloudflare.F("example-worker"),
              },
            }),
            VectorizeBindings: cloudflare.F(map[string]pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
              "VECTORIZE": pages.ProjectDeploymentConfigsProductionVectorizeBindingParam{
                IndexName: cloudflare.F("my_index"),
              },
            }),
          }),
        }),
        Name: cloudflare.F("NextJS Blog"),
        ProductionBranch: cloudflare.F("main"),
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
