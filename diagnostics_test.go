package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagnosticsPerformTraceroute(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `
		{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "target": "1.1.1.1",
      "colos": [
        {
          "error": "",
          "colo": {
            "name": "den01",
            "city": "Denver, CO, US"
          },
          "traceroute_time_ms": 969,
          "target_summary": {
            "asn": "",
            "ip": "1.1.1.1",
            "name": "1.1.1.1",
            "packet_count": 3,
            "mean_latency_ms": 0.021,
            "std_dev_latency_ms": 0.011269427669584647,
            "min_latency_ms": 0.014,
            "max_latency_ms": 0.034
          },
          "hops": [
            {
              "packets_ttl": 1,
              "packets_sent": 3,
              "packets_lost": 0,
              "nodes": [
                {
                  "asn": "AS13335",
                  "ip": "1.1.1.1",
                  "name": "one.one.one.one",
                  "packet_count": 3,
                  "mean_latency_ms": 0.021,
                  "std_dev_latency_ms": 0.011269427669584647,
                  "min_latency_ms": 0.014,
                  "max_latency_ms": 0.034
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
		`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/diagnostics/traceroute", handler)

	want := []DiagnosticsTracerouteResponseResult{{
		Target: "1.1.1.1",
		Colos: []DiagnosticsTracerouteResponseColos{{
			Error: "",
			Colo: DiagnosticsTracerouteResponseColo{
				Name: "den01", City: "Denver, CO, US",
			},
			TracerouteTimeMs: 969,
			TargetSummary: DiagnosticsTracerouteResponseNodes{
				Asn:             "",
				IP:              "1.1.1.1",
				Name:            "1.1.1.1",
				PacketCount:     3,
				MeanLatencyMs:   0.021,
				StdDevLatencyMs: 0.011269427669584647,
				MinLatencyMs:    0.014,
				MaxLatencyMs:    0.034,
			},
			Hops: []DiagnosticsTracerouteResponseHops{{
				PacketsTTL:  1,
				PacketsSent: 3,
				PacketsLost: 0,
				Nodes: []DiagnosticsTracerouteResponseNodes{{
					Asn:             "AS13335",
					IP:              "1.1.1.1",
					Name:            "one.one.one.one",
					PacketCount:     3,
					MeanLatencyMs:   0.021,
					StdDevLatencyMs: 0.011269427669584647,
					MinLatencyMs:    0.014,
					MaxLatencyMs:    0.034,
				}},
			}},
		}},
	},
	}

	opts := DiagnosticsTracerouteConfigurationOptions{PacketsPerTTL: 1, PacketType: "imcp", MaxTTL: 1, WaitTime: 1}
	trace, err := client.PerformTraceroute("01a7362d577a6c3019a474fd6f485823", []string{"1.1.1.1"}, []string{"den01"}, opts)

	if assert.NoError(t, err) {
		assert.Equal(t, want, trace)
	}
}
