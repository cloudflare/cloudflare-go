package workers_test

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7/workers"
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
	t.Run("ScriptScriptAndVersionSettingGetResponsePlacement", func(t *testing.T) {
		var empty, smart, region workers.ScriptScriptAndVersionSettingGetResponsePlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers.ScriptScriptAndVersionSettingGetResponsePlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})

	t.Run("ScriptScriptAndVersionSettingEditResponsePlacement", func(t *testing.T) {
		var empty, smart, region workers.ScriptScriptAndVersionSettingEditResponsePlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers.ScriptScriptAndVersionSettingEditResponsePlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})

	t.Run("VersionPlacement", func(t *testing.T) {
		var empty, smart, region workers.VersionPlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers.VersionPlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})

	t.Run("WorkerPreviewsBaseConfigPlacement", func(t *testing.T) {
		var empty, smart, region workers.WorkerPreviewsBaseConfigPlacement
		decodePlacement(t, `{}`, &empty)
		decodePlacement(t, `{"mode":"smart"}`, &smart)
		decodePlacement(t, `{"region":"aws:us-east-1"}`, &region)
		if smart.Mode != workers.WorkerPreviewsBaseConfigPlacementModeModeSmart {
			t.Errorf("Mode = %q, want %q", smart.Mode, "smart")
		}
		if region.Region != "aws:us-east-1" {
			t.Errorf("Region = %q, want %q", region.Region, "aws:us-east-1")
		}
	})
}
