package cloudflare_test

import (
	"context"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	namespace = "xxxxxx96ee002e8xxxxxx665354c0449"
	accountID = "xxxxxx10ee002e8xxxxxx665354c0410"
)

func ExampleAPI_CreateWorkersKVNamespace() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
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
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
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
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
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
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.UpdateWorkersKVNamespace(context.Background(), namespace, &cloudflare.WorkersKVNamespaceRequest{Title: "test_title"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_WriteWorkersKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	payload := []byte("test payload")
	key := "test_key"

	resp, err := api.WriteWorkersKV(context.Background(), namespace, key, payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}

func ExampleAPI_WriteWorkersKVBulk() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	payload := cloudflare.WorkersKVBulkWriteRequest{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}

	resp, err := api.WriteWorkersKVBulk(context.Background(), namespace, payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_ReadWorkersKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.ReadWorkersKV(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resp)
}

func ExampleAPI_DeleteWorkersKV() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.DeleteWorkersKV(context.Background(), namespace, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func ExampleAPI_ListWorkersKVs() {
	api, err := cloudflare.New(apiKey, user, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.ListWorkersKVs(context.Background(), namespace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
