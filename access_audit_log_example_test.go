package cloudflare_test

import (
	"encoding/json"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func ExampleAPI_AccessAuditLogs() {
	api, err := cloudflare.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	filterOpts := cloudflare.AccessAuditLogFilterOptions{}
	results, _ := api.AccessAuditLogs("someaccountid", filterOpts)

	for _, record := range results {
		b, _ := json.Marshal(record)
		fmt.Println(string(b))
	}
}
