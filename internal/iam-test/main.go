package main

import (
	"context"
	"encoding/json"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

func main() {
	// Test module from the IAM team. This will be used to test the zola and non-zola API functionality

	// Construct a new API object using a global API key or scoped API token (preferred)
	var options []cloudflare.Option
	options = append(options, cloudflare.BaseURL("https://api.staging.cloudflare.com/client/v4"))

	// Non ZoLA
	nonZolaApi, err := cloudflare.New(
		os.Getenv("CLOUDFLARE_NONZOLA_API_KEY"),
		os.Getenv("CLOUDFLARE_NONZOLA_API_EMAIL"),
		options...,
	)
	if err != nil {
		log.Fatalf("ERROR: Failed to construct new API object: %v\n", err)
	}
	ctx := context.Background()

	log.Printf("Testing API with non-ZoLA account\n")
	runApiTest(
		ctx,
		nonZolaApi,
		os.Getenv("CLOUDFLARE_NONZOLA_ACCOUNT_NAME"),
	)

	// ZoLA
	zolaApi, err := cloudflare.New(
		os.Getenv("CLOUDFLARE_ZOLA_API_KEY"),
		os.Getenv("CLOUDFLARE_ZOLA_API_EMAIL"),
		options...,
	)
	//zolaApi, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_ZOLA_API_TOKEN"))
	if err != nil {
		log.Fatalf("ERROR: Failed to construct new API object: %v\n", err)
	}
	ctx = context.Background()

	log.Printf("Testing API with ZoLA account\n")
	runApiTest(
		ctx,
		zolaApi,
		os.Getenv("CLOUDFLARE_ZOLA_ACCOUNT_NAME"),
	)
}

func runApiTest(ctx context.Context, api *cloudflare.API, accountName string) {
	// Accounts endpoints
	// List accounts
	accounts, _, err := api.Accounts(ctx, cloudflare.AccountsListParams{Name: accountName})
	if err != nil {
		log.Printf("ERROR: Failed to fetch list of accounts: %v\n", err)
	}

	// Get account details
	accountId := accounts[0].ID
	account, _, err := api.Account(ctx, accountId)
	if err != nil {
		log.Printf("ERROR: Failed to fetch account: %v\n", err)
	}
	debugPrint("current account", account.Name) // should match the one provided in env

	user, err := api.UserDetails(ctx)
	if err != nil {
		log.Printf("ERROR: Failed to fetch user details: %v\n", err)
	}
	debugPrint("current user", user.Email)

	// TODO: Update user with a role and a policy and verify expected response
	// Policies and roles expected compatibility:
	// - Zola/Policies - Yes
	// - Zola/Roles - Yes
	// - Nonzola/Policies - No
	// - Nonzola/Roles - Yes

	// Roles
	// TODO: Update user with roles and verify

	// List roles
	accountRoles, err := api.AccountRoles(ctx, accountId)
	if err != nil {
		log.Printf("ERROR: Failed to fetch list of account roles: %v\n", err)
	}
	debugPrint("number of account roles found", len(accountRoles))
	debugPrint("first role listed", accountRoles[0].Name)

	// Policies
	// TODO: Update user with policy and verify

	// List policies
	// note: a policy consists of a permission group and a resource group
	permissionGroups, _ := api.ListAPITokensPermissionGroups(ctx)
	debugPrint("number of api tokens permission groups", len(permissionGroups))
	debugPrint("first permission group listed", permissionGroups[0].Name)

}

func debugPrint(name string, object any) {
	s, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Printf("ERROR: Failed to parse json format for: %+v\n", name)
	}
	log.Printf("%+v: %+v\n", name, string(s))
}
