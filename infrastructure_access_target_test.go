package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testInfrastructureAccessTargetId = "019205b5-97d7-7272-b00e-0ea05e61a124"

var (
	infrastructureAccessTargetCreatedOn, _  = time.Parse(time.RFC3339, "2024-08-25T05:00:22Z")
	infrastructureAccessTargetModifiedOn, _ = time.Parse(time.RFC3339, "2024-08-25T05:00:22Z")
	expectedInfrastructureAccessTarget      = InfrastructureAccessTarget{
		Hostname: "infra-access-target",
		ID:       testInfrastructureAccessTargetId,
		IP: InfrastructureAccessTargetIPInfo{
			IPV4: &InfrastructureAccessTargetIPDetails{
				IPAddr:           "187.26.29.249",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
			IPV6: &InfrastructureAccessTargetIPDetails{
				IPAddr:           "2001:0db8:0000:0000:0000:0000:0000:1000",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
		},
		CreatedAt:  infrastructureAccessTargetCreatedOn.String(),
		ModifiedAt: infrastructureAccessTargetModifiedOn.String(),
	}
	expectedInfrastructureModified = InfrastructureAccessTarget{
		Hostname: "infra-access-target-modified",
		ID:       testInfrastructureAccessTargetId,
		IP: InfrastructureAccessTargetIPInfo{
			IPV4: &InfrastructureAccessTargetIPDetails{
				IPAddr:           "250.26.29.250",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
		},
		CreatedAt:  infrastructureAccessTargetCreatedOn.String(),
		ModifiedAt: infrastructureAccessTargetModifiedOn.String(),
	}
)

func TestInfrastructureTarget_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
			{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": {
			 	"created_at": "2024-08-25T05:00:22Z",
				"hostname": "infra-access-target",
				"id": "019205b5-97d7-7272-b00e-0ea05e61a124",
				"ip": {
				"ipv4": {
					"ip_addr": "198.51.100.1",
					"virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
				},
				"ipv6": {
					"ip_addr": "2001:0db8:0000:0000:0000:0000:0000:1000",
					"virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
				}
				},
				"modified_at": "2024-08-25T05:00:22Z"
			  }
			}`)
	})

	// Make sure missing account ID is thrown
	_, err := client.CreateInfrastructureAccessTarget(context.Background(), AccountIdentifier(""), CreateInfrastructureAccessTargetParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, err := client.CreateInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), CreateInfrastructureAccessTargetParams{
		InfrastructureAccessTargetParams: InfrastructureAccessTargetParams{
			Hostname: "infra-access-target",
			IP: InfrastructureAccessTargetIPInfo{
				IPV4: &InfrastructureAccessTargetIPDetails{
					IPAddr:           "198.51.100.1",
					VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
				},
				IPV6: &InfrastructureAccessTargetIPDetails{
					IPAddr:           "2001:0db8:0000:0000:0000:0000:0000:1000",
					VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
				},
			},
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureAccessTarget, out, "create infrastructure_target structs not equal")
	}
}

func TestInfrastructureTarget_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
	  "created_at": "2024-08-25T05:00:22Z",
	  "hostname": "infra-access-target",
	  "id": "019205b5-97d7-7272-b00e-0ea05e61a124",
	  "ip": {
	  	"ipv4": {
		  "ip_addr": "198.51.100.1",
		  "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
		},
		"ipv6": {
		  "ip_addr": "2001:0db8:0000:0000:0000:0000:0000:1000",
		  "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
		}
	  },
	  "modified_at": "2024-08-25T05:00:22Z"
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 20,
    "count": 1,
    "total_count": 1
  }
}`)
	})

	_, _, err := client.ListInfrastructureAccessTargets(context.Background(), AccountIdentifier(""), InfrastructureAccessTargetListParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, results, err := client.ListInfrastructureAccessTargets(context.Background(), AccountIdentifier(testAccountID), InfrastructureAccessTargetListParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(out), "expected 1 challenge_widgets")
		assert.Equal(t, 20, results.PerPage, "expected 20 per page")
		assert.Equal(t, expectedInfrastructureAccessTarget, out[0], "list challenge_widgets structs not equal")
	}
}

func TestInfrastructureTarget_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureAccessTargetId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
	"created_at": "2024-08-25T05:00:22Z",
    "hostname": "infra-access-target",
    "id": "019205b5-97d7-7272-b00e-0ea05e61a124",
    "ip": {
      "ipv4": {
        "ip_addr": "198.51.100.1",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      },
      "ipv6": {
        "ip_addr": "2001:0db8:0000:0000:0000:0000:0000:1000",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      }
    },
    "modified_at": "2024-08-25T05:00:22Z"
  }
}`)
	})

	_, err := client.GetInfrastructureAccessTarget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.GetInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.GetInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), testInfrastructureAccessTargetId)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureAccessTarget, out, "get infrastructure_target not equal to expected")
	}
}

func TestInfrastructureTarget_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureAccessTargetId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
	"created_at": "2024-08-25T05:00:22Z",
    "hostname": "infra-access-target-modified",
    "id": "019205b5-97d7-7272-b00e-0ea05e61a124",
    "ip": {
      "ipv4": {
        "ip_addr": "198.51.100.2",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      }
    },
    "modified_at": "2024-08-25T05:00:22Z"
  }
}`)
	})

	_, err := client.UpdateInfrastructureAccessTarget(context.Background(), AccountIdentifier(""), UpdateInfrastructureAccessTargetParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.UpdateInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureAccessTargetParams{
		ID:           "",
		ModifyParams: InfrastructureAccessTargetParams{},
	})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.UpdateInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureAccessTargetParams{
		ID: testInfrastructureAccessTargetId,
		ModifyParams: InfrastructureAccessTargetParams{
			// Updates hostname and IPv4 address. Deletes IPv6 address.
			Hostname: "infra-access-target-modified",
			IP: InfrastructureAccessTargetIPInfo{
				IPV4: &InfrastructureAccessTargetIPDetails{
					IPAddr:           "198.51.100.2",
					VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
				},
			},
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureModified, out, "update challenge_widgets structs not equal")
	}
}

func TestInfrastructureTarget_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureAccessTargetId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, ``)
	})

	// Make sure missing account ID is thrown
	err := client.DeleteInfrastructureAccessTarget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.DeleteInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	err = client.DeleteInfrastructureAccessTarget(context.Background(), AccountIdentifier(testAccountID), testInfrastructureAccessTargetId)
	assert.NoError(t, err)
}
