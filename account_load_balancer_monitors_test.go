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

const testAccountLoadBalancerMonitorIdentifier = "f1aba936b94213e5b8dca0c0dbf1f9cc"

func TestListAccountLoadBalancerMonitors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": [
			{
			  "id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			  "created_on": "2014-01-01T05:20:00.12345Z",
			  "modified_on": "2014-02-01T05:20:00.12345Z",
			  "type": "https",
			  "description": "Login page monitor",
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
			  "interval": 90,
			  "expected_body": "alive",
			  "expected_codes": "2xx",
			  "follow_redirects": true,
			  "allow_insecure": true,
			  "consecutive_up": 3,
			  "consecutive_down": 2,
			  "probe_zone": "example.com"
			}
		  ]
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := []AccountLoadBalancerMonitor{
		{
			ID:          testAccountLoadBalancerMonitorIdentifier,
			CreatedOn:   &createdOn,
			ModifiedOn:  &modifiedOn,
			Type:        "https",
			Description: "Login page monitor",
			Method:      http.MethodGet,
			Path:        "/health",
			Header: map[string][]string{
				"Host":     {"example.com"},
				"X-App-ID": {"abc123"},
			},
			Port:            8080,
			Timeout:         3,
			Retries:         0,
			Interval:        90,
			ExpectedBody:    "alive",
			ExpectedCodes:   "2xx",
			FollowRedirects: true,
			AllowInsecure:   true,
			ConsecutiveUp:   3,
			ConsecutiveDown: 2,
			ProbeZone:       "example.com",
		},
	}

	_, err := client.ListAccountLoadBalancerMonitors(context.Background(), ResourceIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	actual, err := client.ListAccountLoadBalancerMonitors(context.Background(), ResourceIdentifier(testAccountID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
			  "type": "https",
			  "description": "Login page monitor",
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
			  "interval": 90,
			  "expected_body": "alive",
			  "expected_codes": "2xx",
			  "follow_redirects": true,
			  "allow_insecure": true,
			  "consecutive_up": 3,
			  "consecutive_down": 2,
			  "probe_zone": "example.com"
			}`, string(b))
		}
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"type": "https",
			"description": "Login page monitor",
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
			"interval": 90,
			"expected_body": "alive",
			"expected_codes": "2xx",
			"follow_redirects": true,
			"allow_insecure": true,
			"consecutive_up": 3,
			"consecutive_down": 2,
			"probe_zone": "example.com"
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:          8080,
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "alive",
		ExpectedCodes: "2xx",

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}
	request := AccountLoadBalancerMonitor{
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "alive",
		ExpectedCodes: "2xx",
		Port:          8080,

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}

	_, err := client.CreateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), request)
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	actual, err := client.CreateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", testAccountID, testAccountLoadBalancerMonitorIdentifier), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"type": "https",
			"description": "Login page monitor",
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
			"interval": 90,
			"expected_body": "alive",
			"expected_codes": "2xx",
			"follow_redirects": true,
			"allow_insecure": true,
			"consecutive_up": 3,
			"consecutive_down": 2,
			"probe_zone": "example.com"
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:            8080,
		Timeout:         3,
		Retries:         0,
		Interval:        90,
		ExpectedBody:    "alive",
		ExpectedCodes:   "2xx",
		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}

	_, err := client.GetAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.GetAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerMonitorIdentifier, err)
	}

	_, err = client.GetAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), "invalid")
	assert.Error(t, err)

	result, err := client.GetAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerMonitorIdentifier)
	if assert.NoError(t, err) {
		assert.Equal(t, want, result)
	}
}

func TestModifyAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", testAccountID, testAccountLoadBalancerMonitorIdentifier), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
			  "id": "f1aba936b94213e5b8dca0c0dbf1f9cc",	
			  "type": "https",
			  "description": "Login page monitor",
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
			  "interval": 90,
			  "expected_body": "kicking",
			  "expected_codes": "2xx",
			  "follow_redirects": true,
			  "allow_insecure": true,
			  "consecutive_up": 3,
			  "consecutive_down": 2,
			  "probe_zone": "example.com"
			}`, string(b))
		}
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"type": "https",
			"description": "Login page monitor",
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
			"interval": 90,
			"expected_body": "kicking",
			"expected_codes": "2xx",
			"follow_redirects": true,
			"allow_insecure": true,
			"consecutive_up": 3,
			"consecutive_down": 2,
			"probe_zone": "example.com"
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:          8080,
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "kicking",
		ExpectedCodes: "2xx",

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}
	request := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:          8080,
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "kicking",
		ExpectedCodes: "2xx",

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}

	_, err := client.UpdateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), request)
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.UpdateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerMonitor{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerMonitorIdentifier, err)
	}

	_, err = client.UpdateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerMonitor{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.UpdateAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestPatchAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", testAccountID, testAccountLoadBalancerMonitorIdentifier), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if assert.NoError(t, err) {
			assert.JSONEq(t, `{
			  "id": "f1aba936b94213e5b8dca0c0dbf1f9cc",	
			  "type": "https",
			  "description": "Login page monitor",
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
			  "interval": 90,
			  "expected_body": "kicking",
			  "expected_codes": "2xx",
			  "follow_redirects": true,
			  "allow_insecure": true,
			  "consecutive_up": 3,
			  "consecutive_down": 2,
			  "probe_zone": "example.com"
			}`, string(b))
		}
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"id": "f1aba936b94213e5b8dca0c0dbf1f9cc",
			"created_on": "2014-01-01T05:20:00.12345Z",
			"modified_on": "2014-02-01T05:20:00.12345Z",
			"type": "https",
			"description": "Login page monitor",
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
			"interval": 90,
			"expected_body": "kicking",
			"expected_codes": "2xx",
			"follow_redirects": true,
			"allow_insecure": true,
			"consecutive_up": 3,
			"consecutive_down": 2,
			"probe_zone": "example.com"
		  }
		}`)
	})

	createdOn, _ := time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2014-02-01T05:20:00.12345Z")
	want := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		CreatedOn:   &createdOn,
		ModifiedOn:  &modifiedOn,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:          8080,
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "kicking",
		ExpectedCodes: "2xx",

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}

	request := AccountLoadBalancerMonitor{
		ID:          testAccountLoadBalancerMonitorIdentifier,
		Type:        "https",
		Description: "Login page monitor",
		Method:      http.MethodGet,
		Path:        "/health",
		Header: map[string][]string{
			"Host":     {"example.com"},
			"X-App-ID": {"abc123"},
		},
		Port:          8080,
		Timeout:       3,
		Retries:       0,
		Interval:      90,
		ExpectedBody:  "kicking",
		ExpectedCodes: "2xx",

		FollowRedirects: true,
		AllowInsecure:   true,
		ConsecutiveUp:   3,
		ConsecutiveDown: 2,
		ProbeZone:       "example.com",
	}

	_, err := client.PatchAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), request)
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.PatchAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerMonitor{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerMonitorIdentifier, err)
	}

	_, err = client.PatchAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), AccountLoadBalancerMonitor{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.PatchAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", testAccountID, testAccountLoadBalancerMonitorIdentifier), func(w http.ResponseWriter, r *http.Request) {
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

	err := client.DeleteAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	err = client.DeleteAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerMonitorIdentifier, err)
	}

	assert.NoError(t, client.DeleteAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), testAccountLoadBalancerMonitorIdentifier))
	assert.Error(t, client.DeleteAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), "invalid"))
}

func TestPreviewAccountLoadBalancerMonitor(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s/preview", testAccountID, testAccountLoadBalancerPoolID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
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

	request := PreviewAccountLoadBalancerMonitorParameters{
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
	want := PreviewAccountLoadBalancerMonitor{
		PreviewID: "f1aba936b94213e5b8dca0c0dbf1f9cc",
		Pools: map[string]string{
			"abwlnp5jbqn45ecgxd03erbgtxtqai0d": "WNAM Datacenter",
			"ve8h9lrcip5n5bbga9yqmdws28ay5d0l": "EEU Datacenter",
		},
	}
	_, err := client.PreviewAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(""), PreviewAccountLoadBalancerMonitorParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.PreviewAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), PreviewAccountLoadBalancerMonitorParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerMonitorIdentifier, err)
	}
	_, err = client.PreviewAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), PreviewAccountLoadBalancerMonitorParameters{ID: "invalid"})
	assert.Error(t, err)

	actual, err := client.PreviewAccountLoadBalancerMonitor(context.Background(), AccountIdentifier(testAccountID), request)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestPreviewAccountLoadBalancerMonitorResult(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/accounts/%s/load_balancers/preview/%s", testAccountID, "p1aba936b94213e5b8dca0c0dbf1f9cc"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"abwlnp5jbqn45ecgxd03erbgtxtqai0d": {
			  "healthy": true,
			  "origins": [
				{
				  "originone.example.com.": {
					"healthy": true,
					"rtt": "66ms",
					"failure_reason": "No failures",
					"response_code": 200
				  }
				}
			  ]
			}
		  }
		}`)
	})
	rttParsed, _ := time.ParseDuration("66ms")
	rtt := Duration{rttParsed}
	want := map[string]PreviewAccountLoadBalancerMonitorResult{
		"abwlnp5jbqn45ecgxd03erbgtxtqai0d": {
			Healthy: true,
			Origins: []map[string]PreviewAccountLoadBalancerMonitorResultOrigins{
				{
					"originone.example.com.": {
						Healthy:       true,
						RTT:           rtt,
						FailureReason: "No failures",
						ResponseCode:  200,
					},
				},
			},
		},
	}
	_, err := client.PreviewAccountLoadBalancerResult(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}
	_, err = client.PreviewAccountLoadBalancerResult(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountLoadBalancerPreviewID, err)
	}
	_, err = client.PreviewAccountLoadBalancerResult(context.Background(), AccountIdentifier(testAccountID), "invalid")
	assert.Error(t, err)

	actual, err := client.PreviewAccountLoadBalancerResult(context.Background(), AccountIdentifier(testAccountID), "p1aba936b94213e5b8dca0c0dbf1f9cc")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
