package magic_transit_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7/magic_transit"
)

var connectorSpecEventKinds = []string{
	"Init", "Leave", "StartAttestation", "FinishAttestationSuccess",
	"FinishAttestationFailure", "StartRotateCryptKey", "FinishRotateCryptKeySuccess",
	"FinishRotateCryptKeyFailure", "StartRotatePki", "FinishRotatePkiSuccess",
	"FinishRotatePkiFailure", "StartUpgrade", "FinishUpgradeSuccess",
	"FinishUpgradeFailure", "BlessSlotSuccess", "BlessSlotPending", "BlessSlotFailure",
	"Reconcile", "ConfigureCloudflaredTunnel", "RekeyInstallBoth", "RekeyStart",
	"RekeyRestart", "RekeyAdvance", "RekeyComplete", "RekeyReset",
	"HaTransition", "HaError", "HaInit", "HaLeave",
}

func connectorKindUsesOpenFallback(kind string) bool {
	switch kind {
	case "BlessSlotSuccess", "BlessSlotPending", "BlessSlotFailure", "RekeyRestart",
		"HaTransition", "HaError", "HaInit", "HaLeave", "FutureEvent":
		return true
	default:
		return false
	}
}

func TestConnectorEventSpecKinds(t *testing.T) {
	for _, endpoint := range []string{"get", "latest"} {
		for _, kind := range connectorSpecEventKinds {
			t.Run(endpoint+"/"+kind, func(t *testing.T) {
				event := decodeConnectorRecordedEvent(t, endpoint, fmt.Sprintf(`{"k":%q}`, kind))
				if event.kind != kind || !event.known {
					t.Errorf("kind=%q known=%v; want %q known=true", event.kind, event.known, kind)
				}
				if connectorKindUsesOpenFallback(kind) {
					if event.union != nil {
						t.Errorf("expected open fallback for %q; got %T", kind, event.union)
					}
				} else if event.union == nil {
					t.Errorf("missing typed variant for %q", kind)
				} else if !strings.HasSuffix(reflect.TypeOf(event.union).String(), "E"+kind) {
					t.Errorf("wrong variant %T for %q", event.union, kind)
				}
			})
		}
	}
}

func TestConnectorEventPayloadPreservation(t *testing.T) {
	for _, endpoint := range []string{"get", "latest"} {
		// Exercise typed controls, every missing spec variant, and a future kind.
		for _, kind := range []string{
			"RekeyStart", "StartUpgrade", "BlessSlotSuccess", "BlessSlotPending",
			"BlessSlotFailure", "RekeyRestart", "HaTransition", "HaError", "HaInit",
			"HaLeave", "FutureEvent",
		} {
			t.Run(endpoint+"/"+kind, func(t *testing.T) {
				// These optional payload fields are permitted by the open-object contract.
				payload := fmt.Sprintf(`{"k":%q,"tunnel_id":"test-tunnel","url":"https://example.com/upgrade","details":{"attempt":2,"ok":true}}`, kind)
				event := decodeConnectorRecordedEvent(t, endpoint, payload)
				if event.kind != kind {
					t.Errorf("kind=%q; want %q", event.kind, kind)
				}
				if event.tunnelID != "test-tunnel" || event.url != "https://example.com/upgrade" {
					t.Errorf("payload lost: tunnel_id=%q url=%q", event.tunnelID, event.url)
				}
				wantDetails := map[string]interface{}{"attempt": float64(2), "ok": true}
				if !reflect.DeepEqual(event.extra["details"], wantDetails) {
					t.Errorf("details=%#v; want %#v", event.extra["details"], wantDetails)
				}
				var gotWire, wantWire interface{}
				if err := json.Unmarshal([]byte(event.raw), &gotWire); err != nil {
					t.Fatalf("invalid raw JSON %q: %v", event.raw, err)
				}
				if err := json.Unmarshal([]byte(payload), &wantWire); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(gotWire, wantWire) {
					t.Errorf("raw event changed: got %s; want %s", event.raw, payload)
				}
				if connectorKindUsesOpenFallback(kind) && event.union != nil {
					t.Errorf("fallback selected unrelated variant %T", event.union)
				}
				if kind == "FutureEvent" && event.known {
					t.Error("future kind incorrectly marked known")
				}
			})
		}
	}
}

func TestConnectorEventDecodeResetsState(t *testing.T) {
	first := []byte(`{"k":"RekeyStart","tunnel_id":"old","details":{"old":true}}`)
	second := []byte(`{"k":"HaInit"}`)
	t.Run("get", func(t *testing.T) {
		var event magic_transit.ConnectorEventGetResponseE
		for _, payload := range [][]byte{first, second} {
			if err := json.Unmarshal(payload, &event); err != nil {
				t.Fatal(err)
			}
		}
		if event.AsUnion() != nil || event.TunnelID != "" || len(event.ExtraFields) != 0 {
			t.Errorf("stale state after fallback decode: %+v", event)
		}
	})
	t.Run("latest", func(t *testing.T) {
		var event magic_transit.ConnectorEventLatestListResponseItemsE
		for _, payload := range [][]byte{first, second} {
			if err := json.Unmarshal(payload, &event); err != nil {
				t.Fatal(err)
			}
		}
		if event.AsUnion() != nil || event.TunnelID != "" || len(event.ExtraFields) != 0 {
			t.Errorf("stale state after fallback decode: %+v", event)
		}
	})
}

type connectorEventSnapshot struct {
	kind, tunnelID, url, raw string
	known                    bool
	extra                    map[string]interface{}
	union                    interface{}
}

func decodeConnectorRecordedEvent(t *testing.T, endpoint, eventJSON string) connectorEventSnapshot {
	t.Helper()
	payload := []byte(fmt.Sprintf(`{"e":%s,"n":1,"t":1700000000,"v":"test"}`, eventJSON))
	if endpoint == "get" {
		var response magic_transit.ConnectorEventGetResponse
		if err := json.Unmarshal(payload, &response); err != nil {
			t.Fatal(err)
		}
		e := response.E
		return connectorEventSnapshot{string(e.K), e.TunnelID, e.URL, e.JSON.RawJSON(), e.K.IsKnown(), e.ExtraFields, e.AsUnion()}
	}
	var response magic_transit.ConnectorEventLatestListResponseItem
	if err := json.Unmarshal(payload, &response); err != nil {
		t.Fatal(err)
	}
	e := response.E
	return connectorEventSnapshot{string(e.K), e.TunnelID, e.URL, e.JSON.RawJSON(), e.K.IsKnown(), e.ExtraFields, e.AsUnion()}
}
