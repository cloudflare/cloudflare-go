package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const testAccountLoadBalancerPoolID = "17b5962d775c646f3f9725cbc7a53df4"

func TestListAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": [{
			"id": "17b5962d775c646f3f9725cbc7a53df4",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"description": "Primary data center - Provider XYZ",
			"name": "primary-dc-1",
			"enabled": false,
			"load_shedding": {
			  "default_percent": 0,
			  "default_policy": "random",
			  "session_percent": 0,
			  "session_policy": "hash"
			},
			"minimum_origins": 2,
			"monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"check_regions": [
			  "WEU",
			  "ENAM"
			],
			"origins": [
			  {
				"name": "app-server-1",
				"address": "0.0.0.0",
				"enabled": true,
				"weight": 0.56,
				"header": {
				  "Host": [
					"example.com"
				  ]
				}
			  }
			],
			"origin_steering": {
			  "policy": "random"
			},
			"notification_email": "someone@example.com,sometwo@example.com",
			"notification_filter": {
			  "origin": {
				"disable": false,
				"healthy": null
			  },
			  "pool": {
				"disable": false,
				"healthy": null
			  }
			}
		  }]
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	//longitude := float32(01.23)
	//latitude := float32(01.23)
	want := []AccountLoadBalancerPool{
		{
			ID:          testAccountLoadBalancerPoolID,
			CreatedOn:   &createdOn,
			ModifiedOn:  &modifiedOn,
			Description: "Primary data center - Provider XYZ",
			Name:        "primary-dc-1",
			Enabled:     false,
			LoadShedding: &AccountLoadBalancerPoolLoadShedding{
				DefaultPercent: 0,
				DefaultPolicy:  "random",
				SessionPercent: 0,
				SessionPolicy:  "hash",
			},
			MinimumOrigins: 2,
			Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
			CheckRegions:   []string{"WEU", "ENAM"},
			Origins: []AccountLoadBalancerOrigin{
				{
					Name:    "app-server-1",
					Address: "0.0.0.0",
					Enabled: true,
					Weight:  0.56,
					Header: map[string][]string{
						"Host": {"example.com"},
					},
				},
			},
			OriginSteering: &AccountLoadBalancerOriginSteering{
				Policy: "random",
			},
			NotificationEmail: "someone@example.com,sometwo@example.com",
			NotificationFilter: &AccountLoadBalancerNotificationFilter{
				Origin: &AccountLoadBalancerNotificationResourceType{
					Disable: false,
					Healthy: nil,
				},
				Pool: &AccountLoadBalancerNotificationResourceType{
					Disable: false,
					Healthy: nil,
				},
			},
			//Longitude: &longitude,
			//Latitude:  &latitude,

		},
	}

	_, err := client.ListAccountLoadBalancerPools(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	actual, err := client.ListAccountLoadBalancerPools(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
			  "description": "Primary data center - Provider XYZ",
			  "name": "primary-dc-1",
			  "enabled": false,
			  "load_shedding": {
				"default_percent": 0,
				"default_policy": "random",
				"session_percent": 0,
				"session_policy": "hash"
			  },
			  "minimum_origins": 2,
			  "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			  "check_regions": [
			  	"WEU",
			  	"ENAM"
				],
			  "origins": [
				{
				  "name": "app-server-1",
				  "address": "0.0.0.0",
				  "enabled": true,
				  "weight": 0.56,
				  "header": {
					"Host": [
					  "example.com"
					]
				  }
				}
			  ],
			  "origin_steering": {
				"policy": "random"
			  },
			  "notification_email": "someone@example.com,sometwo@example.com",
			  "notification_filter": {
				"origin": {
				  "disable": false,
				  "healthy": null
				},
				"pool": {
				  "disable": false,
				  "healthy": null
				}
			  }
			}`, string(b))
		}
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "17b5962d775c646f3f9725cbc7a53df4",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"description": "Primary data center - Provider XYZ",
			"name": "primary-dc-1",
			"enabled": false,
			"load_shedding": {
			  "default_percent": 0,
			  "default_policy": "random",
			  "session_percent": 0,
			  "session_policy": "hash"
			},
			"minimum_origins": 2,
			"monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"check_regions": [
			  "WEU",
			  "ENAM"
			],
			"origins": [
			  {
				"name": "app-server-1",
				"address": "0.0.0.0",
				"enabled": true,
				"weight": 0.56,
				"header": {
				  "Host": [
					"example.com"
				  ]
				}
			  }
			],
			"origin_steering": {
			  "policy": "random"
			},
			"notification_email": "someone@example.com,sometwo@example.com",
			"notification_filter": {
			  "origin": {
				"disable": false,
				"healthy": null
			  },
			  "pool": {
				"disable": false,
				"healthy": null
			  }
			}
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	//longitude := float32(01.23)
	//latitude := float32(01.23)
	want := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}
	request := AccountLoadBalancerPool{
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}

	_, err := client.CreateAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), AccountLoadBalancerPool{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	actual, err := client.CreateAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "17b5962d775c646f3f9725cbc7a53df4",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"description": "Primary data center - Provider XYZ",
			"name": "primary-dc-1",
			"enabled": false,
			"load_shedding": {
			  "default_percent": 0,
			  "default_policy": "random",
			  "session_percent": 0,
			  "session_policy": "hash"
			},
			"minimum_origins": 2,
			"monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"check_regions": [
			  "WEU",
			  "ENAM"
			],
			"origins": [
			  {
				"name": "app-server-1",
				"address": "0.0.0.0",
				"enabled": true,
				"weight": 0.56,
				"header": {
				  "Host": [
					"example.com"
				  ]
				}
			  }
			],
			"origin_steering": {
			  "policy": "random"
			},
			"notification_email": "someone@example.com,sometwo@example.com",
			"notification_filter": {
			  "origin": {
				"disable": false,
				"healthy": null
			  },
			  "pool": {
				"disable": false,
				"healthy": null
			  }
			}
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	//longitude := float32(01.23)
	//latitude := float32(01.23)
	want := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}

	_, err := client.GetAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.GetAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}

	_, err = client.GetAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), "invalid")
	assert.Error(t, err)

	actual, err := client.GetAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerPoolID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetAccountLoadBalancerPoolHealth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/health", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"pool_id": "17b5962d775c646f3f9725cbc7a53df4",
			"pop_health": {
			  "Amsterdam, NL": {
				"healthy": true,
				"origins": [
				  {
					"2001:DB8::5": {
					  "healthy": true,
					  "rtt": "12.1ms",
					  "failure_reason": "No failures",
					  "response_code": 401
					}
				  }
				]
			  }
			}
		  }
		}`)
	})
	rtt, _ := time.ParseDuration("12.1ms")
	want := AccountLoadBalancerPoolHealth{
		ID: testAccountLoadBalancerPoolID,
		PopHealth: map[string]AccountLoadBalancerPoolPopHealth{
			"Amsterdam, NL": {
				Healthy: true,
				Origins: []map[string]AccountLoadBalancerOriginHealth{
					{
						"2001:DB8::5": {
							Healthy:       true,
							RTT:           Duration{rtt},
							FailureReason: "No failures",
							ResponseCode:  401,
						},
					},
				},
			},
		},
	}

	_, err := client.GetAccountLoadBalancerPoolHeath(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.GetAccountLoadBalancerPoolHeath(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}

	_, err = client.GetAccountLoadBalancerPoolHeath(context.Background(), AccountIdentifier(testAccountID), "invalid")
	assert.Error(t, err)

	actual, err := client.GetAccountLoadBalancerPoolHeath(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerPoolID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "17b5962d775c646f3f9725cbc7a53df4",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"description": "Primary data center - Provider XYZ",
			"name": "primary-dc-1",
			"enabled": false,
			"load_shedding": {
			  "default_percent": 0,
			  "default_policy": "random",
			  "session_percent": 0,
			  "session_policy": "hash"
			},
			"minimum_origins": 2,
			"monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"check_regions": [
			  "WEU",
			  "ENAM"
			],
			"origins": [
			  {
				"name": "app-server-1",
				"address": "0.0.0.0",
				"enabled": true,
				"weight": 0.56,
				"header": {
				  "Host": [
					"example.com"
				  ]
				}
			  }
			],
			"origin_steering": {
			  "policy": "random"
			},
			"notification_email": "someone@example.com,sometwo@example.com",
			"notification_filter": {
			  "origin": {
				"disable": false,
				"healthy": null
			  },
			  "pool": {
				"disable": false,
				"healthy": null
			  }
			}
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	//longitude := float32(01.23)
	//latitude := float32(01.23)
	want := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}
	request := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}

	_, err := client.UpdateAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), AccountLoadBalancerPool{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.UpdateAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPool{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}
	_, err = client.UpdateAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPool{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.UpdateAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestPatchAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "17b5962d775c646f3f9725cbc7a53df4",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"description": "Primary data center - Provider XYZ",
			"name": "primary-dc-1",
			"enabled": false,
			"load_shedding": {
			  "default_percent": 0,
			  "default_policy": "random",
			  "session_percent": 0,
			  "session_policy": "hash"
			},
			"minimum_origins": 2,
			"monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"check_regions": [
			  "WEU",
			  "ENAM"
			],
			"origins": [
			  {
				"name": "app-server-1",
				"address": "0.0.0.0",
				"enabled": true,
				"weight": 0.56,
				"header": {
				  "Host": [
					"example.com"
				  ]
				}
			  }
			],
			"origin_steering": {
			  "policy": "random"
			},
			"notification_email": "someone@example.com,sometwo@example.com",
			"notification_filter": {
			  "origin": {
				"disable": false,
				"healthy": null
			  },
			  "pool": {
				"disable": false,
				"healthy": null
			  }
			}
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	//longitude := float32(01.23)
	//latitude := float32(01.23)
	want := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}
	request := AccountLoadBalancerPool{
		ID:          testAccountLoadBalancerPoolID,
		Description: "Primary data center - Provider XYZ",
		Name:        "primary-dc-1",
		Enabled:     false,
		LoadShedding: &AccountLoadBalancerPoolLoadShedding{
			DefaultPercent: 0,
			DefaultPolicy:  "random",
			SessionPercent: 0,
			SessionPolicy:  "hash",
		},
		MinimumOrigins: 2,
		Monitor:        "f1aba936b94213e5b8dca0c0dbf1f9cc",
		CheckRegions:   []string{"WEU", "ENAM"},
		Origins: []AccountLoadBalancerOrigin{
			{
				Name:    "app-server-1",
				Address: "0.0.0.0",
				Enabled: true,
				Weight:  0.56,
				Header: map[string][]string{
					"Host": {"example.com"},
				},
			},
		},
		OriginSteering: &AccountLoadBalancerOriginSteering{
			Policy: "random",
		},
		NotificationEmail: "someone@example.com,sometwo@example.com",
		NotificationFilter: &AccountLoadBalancerNotificationFilter{
			Origin: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
			Pool: &AccountLoadBalancerNotificationResourceType{
				Disable: false,
				Healthy: nil,
			},
		},
		//Longitude: &longitude,
		//Latitude:  &latitude,
	}

	_, err := client.PatchAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), AccountLoadBalancerPool{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.PatchAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPool{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}
	_, err = client.PatchAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPool{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.PatchAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

}

func TestDeleteAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": {
              "id": "f1aba936b94213e5b8dca0c0dbf1f9cc"
            }
        }`)
	})

	err := client.DeleteAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	err = client.DeleteAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}

	assert.NoError(t, client.DeleteAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerPoolID))
	assert.Error(t, client.DeleteAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), "invalid"))
}

func TestPreviewAccountLoadBalancerPool(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/preview", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
			  "type": "https",
			  "method": "GET",
			  "path": "/health",
			  "header": {
				"Host": [
				  "example.com"
				],
				"X-App-ID": [
				  "abc123"
				]
			  },
			  "port": 8080,
			  "timeout": 3,
			  "retries": 0,
			  "expected_body": "alive",
			  "expected_codes": "2xx",
			  "follow_redirects": true,
			  "allow_insecure": true,
			  "probe_zone": "example.com"
			}`, string(b))
		}
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"preview_id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"pools": {
			  "abwlnp5jbqn45ecgxd03erbgtxtqai0d": "WNAM Datacenter",
			  "ve8h9lrcip5n5bbga9yqmdws28ay5d0l": "EEU Datacenter"
			}
		}
  		}`)
	})

	request := AccountLoadBalancerPoolPreviewParameters{
		ID:     testAccountLoadBalancerPoolID,
		Type:   "https",
		Method: "GET",
		Path:   "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:            8080,
		Timeout:         3,
		Retries:         0,
		ExpectedBody:    "alive",
		ExpectedCodes:   "2xx",
		FollowRedirects: true,
		AllowInsecure:   true,
		ProbeZone:       "example.com",
	}
	want := AccountLoadBalancerPoolPreview{
		PreviewID: "f1aba936b94213e5b8dca0c0dbf1f9cc",
		Pools: map[string]string{
			"abwlnp5jbqn45ecgxd03erbgtxtqai0d": "WNAM Datacenter",
			"ve8h9lrcip5n5bbga9yqmdws28ay5d0l": "EEU Datacenter",
		},
	}
	_, err := client.PreviewAccountLoadBalancerPool(context.Background(), AccountIdentifier(""), AccountLoadBalancerPoolPreviewParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.PreviewAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPoolPreviewParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}
	_, err = client.PreviewAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerPoolPreviewParameters{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.PreviewAccountLoadBalancerPool(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListAccountLoadBalancerPoolReferences(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/references", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": [
			{
			  "resource_id": "699d98642c564d2e855e9661899b7252",
			  "resource_name": "www.example.com",
			  "resource_type": "load_balancer",
			  "reference_type": "referrer"
			},
			{
			  "resource_id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			  "resource_name": "Login page monitor",
			  "resource_type": "monitor",
			  "reference_type": "referral"
			}
		  ]
		}`)
	})

	want := []AccountLoadBalancerPoolReferences{
		{
			ResourceID:    "699d98642c564d2e855e9661899b7252",
			ResourceName:  "www.example.com",
			ResourceType:  "load_balancer",
			ReferenceType: "referrer",
		},
		{
			ResourceID:    "f1aba936b94213e5b8dca0c0dbf1f9cc",
			ResourceName:  "Login page monitor",
			ResourceType:  "monitor",
			ReferenceType: "referral",
		},
	}
	_, err := client.ListAccountLoadBalancerPoolReferences(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.ListAccountLoadBalancerPoolReferences(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPoolID, err)
	}
	_, err = client.ListAccountLoadBalancerPoolReferences(context.Background(), AccountIdentifier(testAccountID), "invalid")
	assert.Error(t, err)

	actual, err := client.ListAccountLoadBalancerPoolReferences(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerPoolID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
