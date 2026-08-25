// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestThreatEventTagCategoryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Categories.New(context.TODO(), cloudforce_one.ThreatEventTagCategoryNewParams{
		AccountID:   cloudflare.F("account_id"),
		Name:        cloudflare.F("Actor"),
		Description: cloudflare.F("description"),
		Schema: cloudflare.F([]cloudforce_one.ThreatEventTagCategoryNewParamsSchema{{
			Key:           cloudflare.F("family"),
			Kind:          cloudflare.F(cloudforce_one.ThreatEventTagCategoryNewParamsSchemaKindString),
			AllowedValues: cloudflare.F([]string{"low", "medium", "high", "critical"}),
			Annotations: cloudflare.F(cloudforce_one.ThreatEventTagCategoryNewParamsSchemaAnnotations{
				Confidence: cloudflare.F(true),
				TLP:        cloudflare.F(true),
			}),
			Enforcement: cloudflare.F(cloudforce_one.ThreatEventTagCategoryNewParamsSchemaEnforcementError),
			Format:      cloudflare.F(cloudforce_one.ThreatEventTagCategoryNewParamsSchemaFormatDate),
			Label:       cloudflare.F("Attacker Name"),
			MaxLength:   cloudflare.F(int64(1)),
			NumberConstraint: cloudflare.F(cloudforce_one.ThreatEventTagCategoryNewParamsSchemaNumberConstraint{
				Integer: cloudflare.F(true),
				Max:     cloudflare.F(0.000000),
				Min:     cloudflare.F(0.000000),
			}),
			Properties: cloudflare.F(map[string]interface{}{}),
			Required:   cloudflare.F(true),
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

func TestThreatEventTagCategoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Categories.List(context.TODO(), cloudforce_one.ThreatEventTagCategoryListParams{
		AccountID: cloudflare.F("account_id"),
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

func TestThreatEventTagCategoryDelete(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Categories.Delete(
		context.TODO(),
		"category_uuid",
		cloudforce_one.ThreatEventTagCategoryDeleteParams{
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

func TestThreatEventTagCategoryEditWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatEvents.Tags.Categories.Edit(
		context.TODO(),
		"category_uuid",
		cloudforce_one.ThreatEventTagCategoryEditParams{
			AccountID:   cloudflare.F("account_id"),
			Description: cloudflare.F("description"),
			Name:        cloudflare.F("name"),
			Schema: cloudflare.F([]cloudforce_one.ThreatEventTagCategoryEditParamsSchema{{
				Key:           cloudflare.F("family"),
				Kind:          cloudflare.F(cloudforce_one.ThreatEventTagCategoryEditParamsSchemaKindString),
				AllowedValues: cloudflare.F([]string{"low", "medium", "high", "critical"}),
				Annotations: cloudflare.F(cloudforce_one.ThreatEventTagCategoryEditParamsSchemaAnnotations{
					Confidence: cloudflare.F(true),
					TLP:        cloudflare.F(true),
				}),
				Enforcement: cloudflare.F(cloudforce_one.ThreatEventTagCategoryEditParamsSchemaEnforcementError),
				Format:      cloudflare.F(cloudforce_one.ThreatEventTagCategoryEditParamsSchemaFormatDate),
				Label:       cloudflare.F("Attacker Name"),
				MaxLength:   cloudflare.F(int64(1)),
				NumberConstraint: cloudflare.F(cloudforce_one.ThreatEventTagCategoryEditParamsSchemaNumberConstraint{
					Integer: cloudflare.F(true),
					Max:     cloudflare.F(0.000000),
					Min:     cloudflare.F(0.000000),
				}),
				Properties: cloudflare.F(map[string]interface{}{}),
				Required:   cloudflare.F(true),
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
