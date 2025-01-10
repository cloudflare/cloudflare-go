// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

func TestMemberNewWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.New(context.TODO(), accounts.MemberNewParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Body: accounts.MemberNewParamsBodyIAMCreateMemberWithRoles{
			Email:  cloudflare.F("user@example.com"),
			Roles:  cloudflare.F([]string{"3536bcfad5faccb999b47003c79917fb"}),
			Status: cloudflare.F(accounts.MemberNewParamsBodyIAMCreateMemberWithRolesStatusAccepted),
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

func TestMemberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Members.Update(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberUpdateParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
			Body: accounts.MemberUpdateParamsBodyIAMUpdateMemberWithRoles{
				Roles: cloudflare.F([]shared.RoleParam{{
					ID: cloudflare.F("3536bcfad5faccb999b47003c79917fb"),
				}}),
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

func TestMemberListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Members.List(context.TODO(), accounts.MemberListParams{
		AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
		Direction: cloudflare.F(accounts.MemberListParamsDirectionAsc),
		Order:     cloudflare.F(accounts.MemberListParamsOrderUserFirstName),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(5.000000),
		Status:    cloudflare.F(accounts.MemberListParamsStatusAccepted),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemberDelete(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.Delete(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberDeleteParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
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

func TestMemberGet(t *testing.T) {
	t.Skip("HTTP 422 error from prism")
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
	_, err := client.Accounts.Members.Get(
		context.TODO(),
		"4536bcfad5faccb111b47003c79917fa",
		accounts.MemberGetParams{
			AccountID: cloudflare.F("eb78d65290b24279ba6f44721b3ea3c4"),
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
