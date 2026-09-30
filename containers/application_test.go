// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/containers"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestApplicationNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Containers.Applications.New(context.TODO(), containers.ApplicationNewParams{
		AccountID: cloudflare.F("account-123"),
		Body: containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequest{
			Configuration: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfiguration{
				Image: cloudflare.F("image"),
				AuthorizedKeys: cloudflare.F([]containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationAuthorizedKey{{
					PublicKey: cloudflare.F("public_key"),
					Name:      cloudflare.F("name"),
				}}),
				Command:    cloudflare.F([]string{"myapp", "--default-option"}),
				Entrypoint: cloudflare.F([]string{"/bin/bash"}),
				EnvironmentVariables: cloudflare.F([]containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationEnvironmentVariable{{
					Name:  cloudflare.F("name"),
					Value: cloudflare.F("value"),
				}}),
				InstanceType: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationInstanceTypeLite),
				Observability: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservability{
					Logs: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConfigurationObservabilityLogs{
						Enabled: cloudflare.F(true),
					}),
				}),
			}),
			Instances:        cloudflare.F(int64(0)),
			MaxInstances:     cloudflare.F(int64(0)),
			Name:             cloudflare.F("name"),
			SchedulingPolicy: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestSchedulingPolicyDefault),
			Constraints: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestConstraints{
				Jurisdiction: cloudflare.F("jurisdiction"),
				Regions:      cloudflare.F([]string{"WNAM"}),
			}),
			DurableObjects: cloudflare.F[containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsUnion](containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestDurableObjectsCcDurableObjectsConfigurationNamespaceID{
				NamespaceID: cloudflare.F("14758f1afd44c09b7992073ccf00b43d"),
			}),
			Observability: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservability{
				Logs: cloudflare.F(containers.ApplicationNewParamsBodyCcContainersCreateScheduledApplicationRequestObservabilityLogs{
					Enabled: cloudflare.F(true),
				}),
			}),
			RolloutActiveGracePeriod: cloudflare.F(int64(0)),
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

func TestApplicationListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Containers.Applications.List(context.TODO(), containers.ApplicationListParams{
		AccountID: cloudflare.F("account-123"),
		Image:     cloudflare.F("image"),
		Name:      cloudflare.F("name"),
		PageToken: cloudflare.F("page_token"),
		PerPage:   cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Containers.Applications.Delete(
		context.TODO(),
		"application_id",
		containers.ApplicationDeleteParams{
			AccountID: cloudflare.F("account-123"),
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

func TestApplicationEditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Containers.Applications.Edit(
		context.TODO(),
		"application_id",
		containers.ApplicationEditParams{
			AccountID: cloudflare.F("account-123"),
			Configuration: cloudflare.F(containers.ApplicationEditParamsConfiguration{
				AuthorizedKeys: cloudflare.F([]containers.ApplicationEditParamsConfigurationAuthorizedKey{{
					PublicKey: cloudflare.F("public_key"),
					Name:      cloudflare.F("name"),
				}}),
				WranglerSSH: cloudflare.F(containers.ApplicationEditParamsConfigurationWranglerSSH{
					Enabled: cloudflare.F(true),
					Port:    cloudflare.F(int64(1)),
				}),
			}),
			Constraints: cloudflare.F(containers.ApplicationEditParamsConstraints{
				Jurisdiction: cloudflare.F("jurisdiction"),
				Regions:      cloudflare.F([]string{"WNAM"}),
			}),
			MaxInstances: cloudflare.F(int64(0)),
			Observability: cloudflare.F(containers.ApplicationEditParamsObservability{
				Logs: cloudflare.F(containers.ApplicationEditParamsObservabilityLogs{
					Enabled: cloudflare.F(true),
				}),
			}),
			RolloutActiveGracePeriod: cloudflare.F(int64(0)),
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

func TestApplicationGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Containers.Applications.Get(
		context.TODO(),
		"application_id",
		containers.ApplicationGetParams{
			AccountID: cloudflare.F("account-123"),
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
