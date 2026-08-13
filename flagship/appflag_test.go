// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/flagship"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

func TestAppFlagNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Flagship.Apps.Flags.New(
		context.TODO(),
		"app_id",
		flagship.AppFlagNewParams{
			AccountID:        cloudflare.F("account_id"),
			DefaultVariation: cloudflare.F("x"),
			Enabled:          cloudflare.F(true),
			Key:              cloudflare.F("x"),
			Rules: cloudflare.F([]flagship.AppFlagNewParamsRule{{
				Conditions: cloudflare.F([]flagship.AppFlagNewParamsRulesConditionUnion{flagship.AppFlagNewParamsRulesConditionsObject{
					Attribute: cloudflare.F("x"),
					Operator:  cloudflare.F(flagship.AppFlagNewParamsRulesConditionsObjectOperatorEquals),
					Value:     cloudflare.F[any](map[string]interface{}{}),
				}}),
				Priority:       cloudflare.F(int64(1)),
				ServeVariation: cloudflare.F("x"),
				Rollout: cloudflare.F(flagship.AppFlagNewParamsRulesRollout{
					Percentage: cloudflare.F(0.000000),
					Attribute:  cloudflare.F("x"),
				}),
			}}),
			Variations: cloudflare.F(map[string]flagship.AppFlagNewParamsVariationsUnion{
				"foo": shared.UnionString("string"),
			}),
			Description: cloudflare.F("description"),
			Type:        cloudflare.F(flagship.AppFlagNewParamsTypeBoolean),
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

func TestAppFlagUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Flagship.Apps.Flags.Update(
		context.TODO(),
		"app_id",
		"flag_key",
		flagship.AppFlagUpdateParams{
			AccountID:        cloudflare.F("account_id"),
			DefaultVariation: cloudflare.F("x"),
			Enabled:          cloudflare.F(true),
			Key:              cloudflare.F("x"),
			Rules: cloudflare.F([]flagship.AppFlagUpdateParamsRule{{
				Conditions: cloudflare.F([]flagship.AppFlagUpdateParamsRulesConditionUnion{flagship.AppFlagUpdateParamsRulesConditionsObject{
					Attribute: cloudflare.F("x"),
					Operator:  cloudflare.F(flagship.AppFlagUpdateParamsRulesConditionsObjectOperatorEquals),
					Value:     cloudflare.F[any](map[string]interface{}{}),
				}}),
				Priority:       cloudflare.F(int64(1)),
				ServeVariation: cloudflare.F("x"),
				Rollout: cloudflare.F(flagship.AppFlagUpdateParamsRulesRollout{
					Percentage: cloudflare.F(0.000000),
					Attribute:  cloudflare.F("x"),
				}),
			}}),
			Variations: cloudflare.F(map[string]flagship.AppFlagUpdateParamsVariationsUnion{
				"foo": shared.UnionString("string"),
			}),
			Description: cloudflare.F("description"),
			Type:        cloudflare.F(flagship.AppFlagUpdateParamsTypeBoolean),
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

func TestAppFlagListWithOptionalParams(t *testing.T) {
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
	_, err := client.Flagship.Apps.Flags.List(
		context.TODO(),
		"app_id",
		flagship.AppFlagListParams{
			AccountID: cloudflare.F("account_id"),
			Cursor:    cloudflare.F("cursor"),
			Limit:     cloudflare.F("limit"),
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

func TestAppFlagDelete(t *testing.T) {
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
	_, err := client.Flagship.Apps.Flags.Delete(
		context.TODO(),
		"app_id",
		"flag_key",
		flagship.AppFlagDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestAppFlagGet(t *testing.T) {
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
	_, err := client.Flagship.Apps.Flags.Get(
		context.TODO(),
		"app_id",
		"flag_key",
		flagship.AppFlagGetParams{
			AccountID: cloudflare.F("account_id"),
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
