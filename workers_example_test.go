package cloudflare_test

import (
	"fmt"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"log"
)

var (
	workerScript = "addEventListener('fetch', event => {\n    event.passThroughOnException()\nevent.respondWith(handleRequest(event.request))\n})\n\nasync function handleRequest(request) {\n    return fetch(request)\n}"
)

func ExampleAPI_UploadWorker() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.UploadWorker(zoneID, workerScript)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_DownloadWorker() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.DownloadWorker(zoneID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_DeleteWorker() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}
	res, err := api.DeleteWorker(zoneID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_CreateWorkerRoute() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}
	route := cloudflare.WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
	res, err := api.CreateWorkerRoute(zoneID, route)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_UpdateWorkerRoute() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}
	//pull from existing list of routes to perform update on
	routesResponse, err := api.ListWorkerRoutes(zoneID)
	if err != nil {
		log.Fatal(err)
	}
	route := cloudflare.WorkerRoute{Pattern: "app2.example.com/*", Enabled: true}
	//update first route retrieved from the listWorkerRoutes call with details above
	res, err := api.UpdateWorkerRoute(zoneID, routesResponse.Routes[0].ID, route)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_ListWorkerRoutes() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}
	res, err := api.ListWorkerRoutes(zoneID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

func ExampleAPI_DeleteWorkerRoute() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}
	//pull from existing list of routes to perform delete on
	routesResponse, err := api.ListWorkerRoutes(zoneID)
	if err != nil {
		log.Fatal(err)
	}
	//delete first route retrieved from the listWorkerRoutes call
	res, err := api.DeleteWorkerRoute(zoneID, routesResponse.Routes[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res)
}
