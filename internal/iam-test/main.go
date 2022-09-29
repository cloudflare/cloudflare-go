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
	// User's account memberships endpoints
	// TODO:
	// 	 list user memberships
	//	 get user membership details
	//	 update user memberships
	//	 delete user memberships

	// - Get user details
	user, err := api.UserDetails(ctx)
	if err != nil {
		log.Printf("ERROR: Failed to fetch user details: %v\n", err)
	}
	debugPrint("[user] get user details", user)

	// Accounts endpoints
	// - List accounts
	accounts, _, err := api.Accounts(ctx, cloudflare.AccountsListParams{})
	if err != nil {
		log.Printf("ERROR: Failed to fetch list of accounts: %v\n", err)
	}
	for i := 0; i < len(accounts); i++ {
		debugPrint("", accounts[i].Name)
	}

	// - Get account details
	accountId := accounts[0].ID
	account, _, err := api.Account(ctx, accountId)
	if err != nil {
		log.Printf("ERROR: Failed to fetch account: %v\n", err)
	}
	debugPrint("[account] get account details", account)

	// - TODO: Update account
	//updatedAccount, err := api.UpdateAccount(ctx, accountId, account)
	//if err != nil {
	//	log.Printf("ERROR: Failed to fetch account: %v\n", err)
	//}
	//debugPrint("[account] update account", updatedAccount)

	// Account memberships endpoints
	// - List all members
	accountMembers, _, err := api.AccountMembers(ctx, accountId, cloudflare.PaginationOptions{})
	if err != nil {
		log.Printf("ERROR: Failed to fetch account members: %v\n", err)
	}
	for i := 0; i < len(accountMembers); i++ {
		debugPrint("", accountMembers[i].User.Email)
	}

	// - TODO: Add account member
	//newMemberEmail := "jtdocf+cfgotest@gmail.com" // TODO: use testing email
	//var roles []string                            // TODO: need valid roles
	//newAccountMember, err := api.CreateAccountMember(ctx, accountId, newMemberEmail, roles)
	//if err != nil {
	//	log.Printf("ERROR: Failed to create account member: %v\n", err)
	//}
	//debugPrint("[account member] newAccountMember", newAccountMember)

	// - TODO: Update account member
	//updatedAccountMember, err := api.UpdateAccountMember(ctx, accountId, "someUserId", newAccountMember)
	//if err != nil {
	//	log.Printf("ERROR: Failed to update account member: %v\n", err)
	//}
	//debugPrint("[account member] updatedAccountMember", updatedAccountMember)

	// - TODO: Delete account member (and validate)
	//err = api.DeleteAccountMember(ctx, accountId, "someUserId")
	//if err != nil {
	//	log.Printf("ERROR: Failed to delete account member: %v\n", err)
	//}

	// Account roles
	// List all
	accountRoles, err := api.AccountRoles(ctx, accountId)
	if err != nil {
		log.Printf("ERROR: Failed to fetch list of account roles: %v\n", err)
	}
	debugPrint("[account roles] list roles", accountRoles)

	// Get details
	roleId := accountRoles[0].ID
	accountRole, err := api.AccountRole(ctx, accountId, roleId)
	if err != nil {
		log.Printf("ERROR: Failed to fetch list of account roles: %v\n", err)
	}
	debugPrint("[account roles] get role details", accountRole)

}

func debugPrint(name string, object any) {
	s, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Printf("ERROR: Failed to parse json format for: %+v\n", name)
	}
	log.Printf("%+v: %+v\n", name, string(s))
}
