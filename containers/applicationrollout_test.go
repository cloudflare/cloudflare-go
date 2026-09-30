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

func TestApplicationRolloutNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Containers.Applications.Rollouts.New(
		context.TODO(),
		"application_id",
		containers.ApplicationRolloutNewParams{
			AccountID:   cloudflare.F("account-123"),
			Description: cloudflare.F("description"),
			Strategy:    cloudflare.F(containers.ApplicationRolloutNewParamsStrategyRolling),
			TargetConfiguration: cloudflare.F(containers.ApplicationRolloutNewParamsTargetConfiguration{
				AuthorizedKeys: cloudflare.F([]containers.ApplicationRolloutNewParamsTargetConfigurationAuthorizedKey{{
					PublicKey: cloudflare.F("public_key"),
					Name:      cloudflare.F("name"),
				}}),
				Command:    cloudflare.F([]string{"myapp", "--default-option"}),
				Entrypoint: cloudflare.F([]string{"/bin/bash"}),
				EnvironmentVariables: cloudflare.F([]containers.ApplicationRolloutNewParamsTargetConfigurationEnvironmentVariable{{
					Name:  cloudflare.F("name"),
					Value: cloudflare.F("value"),
				}}),
				Image:        cloudflare.F("image"),
				InstanceType: cloudflare.F(containers.ApplicationRolloutNewParamsTargetConfigurationInstanceTypeLite),
				Observability: cloudflare.F(containers.ApplicationRolloutNewParamsTargetConfigurationObservability{
					Logs: cloudflare.F(containers.ApplicationRolloutNewParamsTargetConfigurationObservabilityLogs{
						Enabled: cloudflare.F(true),
					}),
				}),
			}),
			Kind:           cloudflare.F(containers.ApplicationRolloutNewParamsKindFullAuto),
			Percentage:     cloudflare.F(int64(0)),
			StepPercentage: cloudflare.F(containers.ApplicationRolloutNewParamsStepPercentage5),
			Steps: cloudflare.F([]containers.ApplicationRolloutNewParamsStep{{
				Description: cloudflare.F("description"),
				StepSize: cloudflare.F(containers.ApplicationRolloutNewParamsStepsStepSize{
					Percentage: cloudflare.F(int64(0)),
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
