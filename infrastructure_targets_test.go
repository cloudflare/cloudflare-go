package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// randomly generated uuid
const testInfrastructureTargetId = "019205b5-97d7-7272-b00e-0ea05e61a124"

var (
	infrastrctureTargetCreatedOn, _  = time.Parse(time.RFC3339, "2024-08-25T05:00:22Z")
	infrastrctureTargetModifiedOn, _ = time.Parse(time.RFC3339, "2024-08-25T05:00:22Z")
	expectedInfrastructureTarget     = Target{
		Hostname: "infra-access-target",
		ID:       testInfrastructureTargetId,
		IP: IPInfo{
			IPV4: &IPDetails{
				IPAddr:           "187.26.29.249",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
			IPV6: &IPDetails{
				IPAddr:           "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
		},
		CreatedAt:  infrastrctureTargetCreatedOn,
		ModifiedAt: infrastrctureTargetModifiedOn,
	}
	expectedInfrastructureModified = Target{
		Hostname: "infra-access-target-modified",
		ID:       testInfrastructureTargetId,
		IP: IPInfo{
			IPV4: &IPDetails{
				IPAddr:           "250.26.29.250",
				VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
			},
		},
		CreatedAt:  infrastrctureTargetCreatedOn,
		ModifiedAt: infrastrctureTargetModifiedOn,
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
					"ip_addr": "187.26.29.249",
					"virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
				},
				"ipv6": {
					"ip_addr": "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
					"virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
				}
				},
				"modified_at": "2024-08-25T05:00:22Z"
			  }
			}`)
	})

	// Make sure missing account ID is thrown
	_, err := client.CreateInfrastructureTarget(context.Background(), AccountIdentifier(""), CreateInfrastructureTargetParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, err := client.CreateInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), CreateInfrastructureTargetParams{
		InfrastructureTargetParams: InfrastructureTargetParams{
			Hostname: "infra-access-target",
			IP: IPInfo{
				IPV4: &IPDetails{
					IPAddr:           "187.26.29.249",
					VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
				},
				IPV6: &IPDetails{
					IPAddr:           "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
					VirtualNetworkId: "c77b744e-acc8-428f-9257-6878c046ed55",
				},
			},
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureTarget, out, "create infrastructure_target structs not equal")
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
		  "ip_addr": "187.26.29.249",
		  "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
		},
		"ipv6": {
		  "ip_addr": "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
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

	_, _, err := client.ListInfrastructureTargets(context.Background(), AccountIdentifier(""), TargetListParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, results, err := client.ListInfrastructureTargets(context.Background(), AccountIdentifier(testAccountID), TargetListParams{})
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(out), "expected 1 challenge_widgets")
		assert.Equal(t, 20, results.PerPage, "expected 20 per page")
		assert.Equal(t, expectedInfrastructureTarget, out[0], "list challenge_widgets structs not equal")
	}
}

func TestInfrastructureTarget_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureTargetId, func(w http.ResponseWriter, r *http.Request) {
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
        "ip_addr": "187.26.29.249",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      },
      "ipv6": {
        "ip_addr": "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      }
    },
    "modified_at": "2024-08-25T05:00:22Z"
  }
}`)
	})

	_, err := client.GetInfrastructureTarget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.GetInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.GetInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), testInfrastructureTargetId)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedInfrastructureTarget, out, "get infrastructure_target not equal to expected")
	}
}

func TestInfrastructureTarget_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureTargetId, func(w http.ResponseWriter, r *http.Request) {
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
        "ip_addr": "250.26.29.250",
        "virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55"
      }
    },
    "modified_at": "2024-08-25T05:00:22Z"
  }
}`)
	})

	_, err := client.UpdateInfrastructureTarget(context.Background(), AccountIdentifier(""), UpdateInfrastructureTargetParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.UpdateInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureTargetParams{
		ID:           "",
		ModifyParams: InfrastructureTargetParams{},
	})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	out, err := client.UpdateInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), UpdateInfrastructureTargetParams{
		ID: testInfrastructureTargetId,
		ModifyParams: InfrastructureTargetParams{
			// Updates hostname and IPv4 address. Deletes IPv6 address.
			Hostname: "infra-access-target-modified",
			IP: IPInfo{
				IPV4: &IPDetails{
					IPAddr:           "250.26.29.250",
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

	mux.HandleFunc("/accounts/"+testAccountID+"/infrastructure/targets/"+testInfrastructureTargetId, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, ``)
	})

	// Make sure missing account ID is thrown
	err := client.DeleteInfrastructureTarget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.DeleteInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTargetId, err)
	}

	err = client.DeleteInfrastructureTarget(context.Background(), AccountIdentifier(testAccountID), testInfrastructureTargetId)
	assert.NoError(t, err)
}
