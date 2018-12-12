package cloudflare_test

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	namespace = "xxxxxx96ee002e8xxxxxx665354c0449"
	org       = "xxxxxx10ee002e8xxxxxx665354c0410"
)

func ExampleAPI_CreateStorageNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	req := &cloudflare.StorageNamespaceRequest{Title: "test_namespace"}
	snr, err := api.CreateStorageNamespace(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(snr)
}

func ExampleAPI_ListStorageNamespaces() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	lsr, err := api.ListStorageNamespaces(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lsr)
}

func ExampleAPI_DeleteStorageNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	response, err := api.DeleteStorageNamespace(context.Background(), namespace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func ExampleAPI_UpdateStorageNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.UpdateStorageNamespace(context.Background(), namespace, &cloudflare.StorageNamespaceRequest{Title: "test_title"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_CreateStorageKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	payload := strings.NewReader("test payload")
	key := "test_key"

	resp, err := api.CreateStorageKV(context.Background(), namespace, key, payload)
	if err != nil {
		log.Fatal(err)
	}

	if payload.Len() > 0 {
		log.Fatalf("Reader was not drained, remaining bytes: %d", payload.Len())
	}

	fmt.Println(resp)
}

func ExampleAPI_ReadStorageKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.ReadStorageKV(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resp)
}

func ExampleAPI_DeleteStorageKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.DeleteStorageKV(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func ExampleAPI_ListStorageKeys() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingOrganization(org))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.ListStorageKeys(context.Background(), namespace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
