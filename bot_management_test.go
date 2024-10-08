package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBotManagement(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "2.0.0", r.Header.Get("Cloudflare-Version"))
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
	"success" : true,
	"errors": [],
	"messages": [],
	"result": {
		"enable_js": false,
		"fight_mode": true,
		"using_latest_model": true,
        "ai_bots_protection": "disabled"
	}
}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/bot_management", handler)

	want := BotManagement{
		EnableJS:         BoolPtr(false),
		FightMode:        BoolPtr(true),
		UsingLatestModel: BoolPtr(true),
		AIBotsProtection: StringPtr("disabled"),
	}

	actual, err := client.GetBotManagement(context.Background(), ZoneIdentifier(testZoneID))

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateBotManagement(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		assert.Equal(t, "2.0.0", r.Header.Get("Cloudflare-Version"))
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
	"success" : true,
	"errors": [],
	"messages": [],
	"result": {
		"enable_js": false,
		"fight_mode": true,
		"using_latest_model": true,
        "ai_bots_protection": "block"
	}
}
		`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/bot_management", handler)

	bmData := UpdateBotManagementParams{
		EnableJS:         BoolPtr(false),
		FightMode:        BoolPtr(true),
		AIBotsProtection: StringPtr("block"),
	}

	want := BotManagement{
		EnableJS:         BoolPtr(false),
		FightMode:        BoolPtr(true),
		UsingLatestModel: BoolPtr(true),
		AIBotsProtection: StringPtr("block"),
	}

	actual, err := client.UpdateBotManagement(context.Background(), ZoneIdentifier(testZoneID), bmData)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
