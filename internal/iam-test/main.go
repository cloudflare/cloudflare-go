package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

func main() {
	// Test module from the IAM team. This will be used to test the zola and non-zola API functionality

	// Construct a new API object using a global API key or scoped API token (preferred)
	var options []cloudflare.Option
	options = append(options, cloudflare.BaseURL("https://api.staging.cloudflare.com/client/v4"))
	zolaApi, err := cloudflare.New(
		os.Getenv("CLOUDFLARE_ZOLA_API_KEY"),
		os.Getenv("CLOUDFLARE_ZOLA_API_EMAIL"),
		options...,
	)
	//zolaApi, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_ZOLA_API_TOKEN"))
	if err != nil {
		log.Fatalf("ERROR: Failed to construct new API object: %v\n", err)
	}
	ctx := context.Background()

	// Testing variables
	accountName := os.Getenv("CLOUDFLARE_ZOLA_ACCOUNT_NAME")
	newMemberEmail := os.Getenv("CLOUDFLARE_ZOLA_INVITEE")

	/*
		TODO: currently using a staging zola account, need nonzola

		TODO: List of endpoints rom API docs:
		- [ ] user account memberships
			- [ ] list memberships
			- [ ] get membership details
			- [ ] update memberships
			- [ ] delete memberships
		- [ ] accounts
			- [ ] list accounts
			- [ ] get account details
			- [ ] update account
		- [ ] account members
			- [ ] list
			- [ ] add
			- [ ] update
			- [ ] remove
		- [ ] account roles
			- [ ] list
			- [ ] get details
	*/

	// User account memberships endpoints
	fmt.Printf("\n\n---\n" +
		"User account memberships" +
		"\n---\n\n")

	user, err := zolaApi.UserDetails(ctx)
	if err != nil {
		log.Printf("ERROR: Failed to fetch user details: %v\n", err)
	}
	debugPrint("user", user)

	updatedUser, err := zolaApi.UpdateUser(ctx, &user)
	if err != nil {
		log.Printf("ERROR: Failed to fetch user details: %v\n", err)
	}
	debugPrint("updatedUser", updatedUser)

	// Accounts endpoints
	fmt.Printf("\n\n---\n" +
		"Accounts" +
		"\n---\n\n")

	accounts, resultInfo, err := zolaApi.Accounts(ctx, cloudflare.AccountsListParams{Name: accountName})
	if err != nil {
		log.Printf("ERROR: Failed to fetch account: %v\n", err)
	}
	accountId := accounts[0].ID

	debugPrint("account", accounts)
	debugPrint("resultInfo", resultInfo)
	debugPrint("accountId", accountId)

	// Account memberships endpoints
	fmt.Printf("\n\n---\n" +
		"Account memberships" +
		"\n---\n\n")

	newAccountMember, err := zolaApi.CreateAccountMember(ctx, accountId, newMemberEmail, []string{})
	if err != nil {
		log.Printf("ERROR: Failed to create account member: %v\n", err)
	}
	debugPrint("newAccountMember", newAccountMember)

	accountMembers, resultInfo, err := zolaApi.AccountMembers(ctx, accountId, cloudflare.PaginationOptions{})
	if err != nil {
		log.Printf("ERROR: Failed to fetch account members: %v\n", err)
	}
	for i := 0; i < len(accountMembers); i++ {
		debugPrint("", accountMembers[i].User.Email)
	}

}

func debugPrint(name string, object any) {
	s, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Printf("ERROR: Failed to parse json format for: %+v\n", name)
	}
	log.Printf("%+v: %+v\n", name, string(s))
}
