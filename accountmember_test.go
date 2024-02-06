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

func TestAccountMemberGet(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.AccountMembers.Get(
		context.TODO(),
		map[string]interface{}{},
		"4536bcfad5faccb111b47003c79917fa",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMemberUpdate(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.AccountMembers.Update(
		context.TODO(),
		map[string]interface{}{},
		"4536bcfad5faccb111b47003c79917fa",
		cloudflare.AccountMemberUpdateParams{
			Roles: cloudflare.F([]cloudflare.AccountMemberUpdateParamsRole{{
				ID: cloudflare.F("3536bcfad5faccb999b47003c79917fb"),
				Permissions: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissions{
					Analytics: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsAnalytics{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Billing: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsBilling{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					CachePurge: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsCachePurge{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNS: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNS{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNSRecords: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNSRecords{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Lb: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLb{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Logs: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLogs{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Organization: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsOrganization{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					SSL: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsSSL{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					WAF: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsWAF{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					ZoneSettings: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZoneSettings{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Zones: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZones{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
				}),
			}, {
				ID: cloudflare.F("3536bcfad5faccb999b47003c79917fb"),
				Permissions: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissions{
					Analytics: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsAnalytics{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Billing: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsBilling{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					CachePurge: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsCachePurge{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNS: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNS{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNSRecords: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNSRecords{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Lb: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLb{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Logs: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLogs{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Organization: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsOrganization{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					SSL: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsSSL{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					WAF: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsWAF{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					ZoneSettings: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZoneSettings{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Zones: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZones{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
				}),
			}, {
				ID: cloudflare.F("3536bcfad5faccb999b47003c79917fb"),
				Permissions: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissions{
					Analytics: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsAnalytics{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Billing: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsBilling{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					CachePurge: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsCachePurge{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNS: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNS{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					DNSRecords: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsDNSRecords{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Lb: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLb{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Logs: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsLogs{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Organization: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsOrganization{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					SSL: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsSSL{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					WAF: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsWAF{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					ZoneSettings: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZoneSettings{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
					Zones: cloudflare.F(cloudflare.AccountMemberUpdateParamsRolesPermissionsZones{
						Read:  cloudflare.F(true),
						Write: cloudflare.F(false),
					}),
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

func TestAccountMemberDelete(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.AccountMembers.Delete(
		context.TODO(),
		map[string]interface{}{},
		"4536bcfad5faccb111b47003c79917fa",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMemberAccountMembersAddMemberWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.AccountMembers.AccountMembersAddMember(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.AccountMemberAccountMembersAddMemberParams{
			Email:  cloudflare.F("user@example.com"),
			Roles:  cloudflare.F([]string{"3536bcfad5faccb999b47003c79917fb", "3536bcfad5faccb999b47003c79917fb", "3536bcfad5faccb999b47003c79917fb"}),
			Status: cloudflare.F(cloudflare.AccountMemberAccountMembersAddMemberParamsStatusAccepted),
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

func TestAccountMemberAccountMembersListMembersWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.AccountMembers.AccountMembersListMembers(
		context.TODO(),
		map[string]interface{}{},
		cloudflare.AccountMemberAccountMembersListMembersParams{
			Direction: cloudflare.F(cloudflare.AccountMemberAccountMembersListMembersParamsDirectionDesc),
			Order:     cloudflare.F(cloudflare.AccountMemberAccountMembersListMembersParamsOrderStatus),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
			Status:    cloudflare.F(cloudflare.AccountMemberAccountMembersListMembersParamsStatusAccepted),
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
