package workers_for_platforms_test

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7/workers_for_platforms"
)

func decodePlacement(t *testing.T, data string, v any) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("decoding %s panicked: %v", data, r)
		}
	}()
	if err := json.Unmarshal([]byte(data), v); err != nil {
		t.Fatalf("decoding %s: %v", data, err)
	}
}

func TestPlacementModeUnmarshal(t *testing.T) {
	t.Run("DispatchNamespaceScriptSettingGetResponsePlacement", func(t *testing.T) {
		var empty, smart, region workers_for_platforms.DispatchNamespaceScriptSettingGetResponsePlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers_for_platforms.DispatchNamespaceScriptSettingGetResponsePlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})

	t.Run("DispatchNamespaceScriptSettingEditResponsePlacement", func(t *testing.T) {
		var empty, smart, region workers_for_platforms.DispatchNamespaceScriptSettingEditResponsePlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers_for_platforms.DispatchNamespaceScriptSettingEditResponsePlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})
}
