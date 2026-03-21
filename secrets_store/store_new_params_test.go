// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6/internal/param"
)

func TestStoreNewParamsMarshalJSON(t *testing.T) {
	tests := []struct {
		name           string
		params         StoreNewParams
		expectedHasKey bool
	}{
		{
			name: "single body element marshals as object not array",
			params: StoreNewParams{
				AccountID: param.Field[string]{Present: true, Value: "account-123"},
				Body: []StoreNewParamsBody{{
					Name: param.Field[string]{Present: true, Value: "test-store"},
				}},
			},
			expectedHasKey: true,
		},
		{
			name: "empty body marshals empty object",
			params: StoreNewParams{
				AccountID: param.Field[string]{Present: true, Value: "account-123"},
				Body:      []StoreNewParamsBody{},
			},
			expectedHasKey: false, // empty body results in empty object
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.params.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON returned error: %v", err)
			}

			// The JSON should be an object, not an array
			var v interface{}
			if err := json.Unmarshal(data, &v); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			// Verify it's a map/object, not an array
			_, isMap := v.(map[string]interface{})
			if !isMap {
				t.Fatalf("expected JSON object, got: %s", string(data))
			}

			// Verify it has the "name" key (the body content)
			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				t.Fatalf("failed to unmarshal to map: %v", err)
			}

			if tt.expectedHasKey {
				if _, ok := result["name"]; !ok {
					t.Errorf("expected 'name' key in JSON output, got: %s", string(data))
				}
			}
		})
	}
}

func TestStoreNewParamsMarshalJSONNotArray(t *testing.T) {
	// This test specifically verifies the bug fix: MarshalJSON should not return an array
	params := StoreNewParams{
		AccountID: param.Field[string]{Present: true, Value: "account-123"},
		Body: []StoreNewParamsBody{{
			Name: param.Field[string]{Present: true, Value: "test-store"},
		}},
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON returned error: %v", err)
	}

	// First character should be '{', not '['
	if len(data) == 0 {
		t.Fatal("MarshalJSON returned empty data")
	}
	if data[0] != '{' {
		t.Errorf("expected JSON object starting with '{', got: %s", string(data))
	}

	// Should not be a JSON array
	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err == nil {
		t.Errorf("MarshalJSON returned JSON array instead of object: %s", string(data))
	}
}

func TestStoreNewParamsMarshalJSONMultipleElements(t *testing.T) {
	// When multiple body elements are provided, only the first should be marshaled
	params := StoreNewParams{
		AccountID: param.Field[string]{Present: true, Value: "account-123"},
		Body: []StoreNewParamsBody{
			{
				Name: param.Field[string]{Present: true, Value: "first-store"},
			},
			{
				Name: param.Field[string]{Present: true, Value: "second-store"},
			},
		},
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON returned error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	// Should only contain "first-store", not "second-store"
	name, ok := result["name"].(string)
	if !ok {
		t.Fatalf("expected 'name' to be a string, got: %T", result["name"])
	}
	if name != "first-store" {
		t.Errorf("expected name 'first-store', got: %s", name)
	}
}
