package cloudflare_test

import (
	"context"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	namespace = "xxxxxx96ee002e8xxxxxx665354c0449"
	org       = "xxxxxx10ee002e8xxxxxx665354c0410"
)

func ExampleAPI_CreateWorkersKVNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	req := &cloudflare.WorkersKVNamespaceRequest{Title: "test_namespace2"}
	response, err := api.CreateWorkersKVNamespace(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func ExampleAPI_ListWorkersKVNamespaces() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	lsr, err := api.ListWorkersKVNamespaces(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lsr)
}

func ExampleAPI_DeleteWorkersKVNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	response, err := api.DeleteWorkersKVNamespace(context.Background(), namespace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func ExampleAPI_UpdateWorkersKVNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.UpdateWorkersKVNamespace(context.Background(), namespace, &cloudflare.WorkersKVNamespaceRequest{Title: "test_title"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_WriteWorkersKVPair() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	payload := "test payload"
	key := "/test_key"

	resp, err := api.WriteWorkersKVPair(context.Background(), namespace, &cloudflare.WorkersKVPair{Name: key, Value: payload})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_WriteWorkersKVPairBulk() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	pairs := []*cloudflare.WorkersKVPair{
		{Name: "key-1", Value: "value-1"},
		{Name: "key-2", Value: "value-2"}}

	resp, err := api.WriteWorkersKVPairBulk(context.Background(), namespace, pairs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_ReadWorkersKVPair() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.ReadWorkersKVPair(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func ExampleAPI_DeleteWorkersKVPair() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.DeleteWorkersKVPair(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func ExampleAPI_ListWorkersKVKeys() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.ListWorkersKVKeys(context.Background(), namespace, &cloudflare.WorkersKVKeyPagination{Count: 1000})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
