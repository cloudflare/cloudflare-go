// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/hyperdrive"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestConfigNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Hyperdrive.Configs.New(context.TODO(), hyperdrive.ConfigNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Hyperdrive: hyperdrive.HyperdriveParam{
			Name: cloudflare.F("example-hyperdrive"),
			Origin: cloudflare.F[hyperdrive.HyperdriveOriginUnionParam](hyperdrive.HyperdriveOriginPublicDatabaseParam{
				Database: cloudflare.F("postgres"),
				Host:     cloudflare.F("database.example.com"),
				Password: cloudflare.F("password"),
				Port:     cloudflare.F(int64(5432)),
				Scheme:   cloudflare.F(hyperdrive.HyperdriveOriginPublicDatabaseSchemePostgres),
				User:     cloudflare.F("postgres"),
			}),
			Caching: cloudflare.F[hyperdrive.HyperdriveCachingUnionParam](hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommonParam{
				Disabled: cloudflare.F(true),
			}),
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

func TestConfigUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Hyperdrive.Configs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Hyperdrive: hyperdrive.HyperdriveParam{
				Name: cloudflare.F("example-hyperdrive"),
				Origin: cloudflare.F[hyperdrive.HyperdriveOriginUnionParam](hyperdrive.HyperdriveOriginPublicDatabaseParam{
					Database: cloudflare.F("postgres"),
					Host:     cloudflare.F("database.example.com"),
					Password: cloudflare.F("password"),
					Port:     cloudflare.F(int64(5432)),
					Scheme:   cloudflare.F(hyperdrive.HyperdriveOriginPublicDatabaseSchemePostgres),
					User:     cloudflare.F("postgres"),
				}),
				Caching: cloudflare.F[hyperdrive.HyperdriveCachingUnionParam](hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommonParam{
					Disabled: cloudflare.F(true),
				}),
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

func TestConfigList(t *testing.T) {
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
	_, err := client.Hyperdrive.Configs.List(context.TODO(), hyperdrive.ConfigListParams{
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

func TestConfigDelete(t *testing.T) {
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
	_, err := client.Hyperdrive.Configs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigDeleteParams{
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

func TestConfigEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Hyperdrive.Configs.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Caching: cloudflare.F[hyperdrive.ConfigEditParamsCachingUnion](hyperdrive.ConfigEditParamsCachingHyperdriveHyperdriveCachingCommon{
				Disabled: cloudflare.F(true),
			}),
			Name: cloudflare.F("example-hyperdrive"),
			Origin: cloudflare.F[hyperdrive.ConfigEditParamsOriginUnion](hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveDatabase{
				Database: cloudflare.F("postgres"),
				Password: cloudflare.F("password"),
				Scheme:   cloudflare.F(hyperdrive.ConfigEditParamsOriginHyperdriveHyperdriveDatabaseSchemePostgres),
				User:     cloudflare.F("postgres"),
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

func TestConfigGet(t *testing.T) {
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
	_, err := client.Hyperdrive.Configs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		hyperdrive.ConfigGetParams{
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
