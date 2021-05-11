package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *API

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup(opts ...Option) {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// disable rate limits and retries in testing - prepended so any provided value overrides this
	opts = append([]Option{UsingRateLimit(100000), UsingRetryPolicy(0, 0, 0)}, opts...)

	// Cloudflare client configured to use test server
	client, _ = New("deadbeef", "cloudflare@example.org", opts...)
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestClient_Headers(t *testing.T) {
	// it should set default headers
	setup()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "cloudflare@example.org", r.Header.Get("X-Auth-Email"))
		assert.Equal(t, "deadbeef", r.Header.Get("X-Auth-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails(context.Background())
	teardown()

	// it should override appropriate default headers when custom headers given
	headers := make(http.Header)
	headers.Set("Content-Type", "application/xhtml+xml")
	headers.Add("X-Random", "a random header")
	setup(Headers(headers))
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "cloudflare@example.org", r.Header.Get("X-Auth-Email"))
		assert.Equal(t, "deadbeef", r.Header.Get("X-Auth-Key"))
		assert.Equal(t, "application/xhtml+xml", r.Header.Get("Content-Type"))
		assert.Equal(t, "a random header", r.Header.Get("X-Random"))
	})
	client.UserDetails(context.Background())
	teardown()

	// it should set X-Auth-User-Service-Key and omit X-Auth-Email and X-Auth-Key when client.authType is AuthUserService
	setup()
	client.SetAuthType(AuthUserService)
	client.APIUserServiceKey = "userservicekey"
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Empty(t, r.Header.Get("X-Auth-Email"))
		assert.Empty(t, r.Header.Get("X-Auth-Key"))
		assert.Empty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "userservicekey", r.Header.Get("X-Auth-User-Service-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails(context.Background())
	teardown()

	// it should set X-Auth-User-Service-Key and omit X-Auth-Email and X-Auth-Key when using NewWithUserServiceKey
	setup()
	client, err := NewWithUserServiceKey("userservicekey")
	assert.NoError(t, err)
	client.BaseURL = server.URL
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Empty(t, r.Header.Get("X-Auth-Email"))
		assert.Empty(t, r.Header.Get("X-Auth-Key"))
		assert.Empty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "userservicekey", r.Header.Get("X-Auth-User-Service-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails(context.Background())
	teardown()

	// it should set Authorization and omit others credential headers when using NewWithAPIToken
	setup()
	client, err = NewWithAPIToken("my-api-token")
	assert.NoError(t, err)
	client.BaseURL = server.URL
	mux.HandleFunc("/zones/123456", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Empty(t, r.Header.Get("X-Auth-Email"))
		assert.Empty(t, r.Header.Get("X-Auth-Key"))
		assert.Empty(t, r.Header.Get("X-Auth-User-Service-Key"))
		assert.Equal(t, "Bearer my-api-token", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails(context.Background())
	teardown()
}

func TestClient_RetryCanSucceedAfterErrors(t *testing.T) {
	setup(UsingRetryPolicy(2, 0, 1))
	defer teardown()

	requestsReceived := 0
	// could test any function, using ListLoadBalancerPools
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		// we are doing three *retries*
		if requestsReceived == 0 {
			// return error causing client to retry
			w.WriteHeader(500)
			fmt.Fprint(w, `{
				"success": false,
				"errors": [ "server created some error"],
				"messages": [],
				"result": []
			}`)
		} else if requestsReceived == 1 {
			// return error causing client to retry
			w.WriteHeader(429)
			fmt.Fprint(w, `{
				"success": false,
				"errors": [ "this is a rate limiting error"],
				"messages": [],
				"result": []
			}`)
		} else {
			// return success response
			fmt.Fprint(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": [
                {
                    "id": "17b5962d775c646f3f9725cbc7a53df4",
                    "created_on": "2014-01-01T05:20:00.12345Z",
                    "modified_on": "2014-02-01T05:20:00.12345Z",
                    "description": "Primary data center - Provider XYZ",
                    "name": "primary-dc-1",
                    "enabled": true,
                    "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc",
                    "origins": [
                      {
                        "name": "app-server-1",
                        "address": "0.0.0.0",
                        "enabled": true
                      }
                    ],
                    "notification_email": "someone@example.com"
                }
            ],
            "result_info": {
                "page": 1,
                "per_page": 20,
                "count": 1,
                "total_count": 2000
            }
        }`)
		}
		requestsReceived++

	}

	mux.HandleFunc("/user/load_balancers/pools", handler)

	_, err := client.ListLoadBalancerPools(context.Background())
	assert.NoError(t, err)
}

func TestClient_RetryReturnsPersistentErrorResponse(t *testing.T) {
	setup(UsingRetryPolicy(2, 0, 1))
	defer teardown()

	// could test any function, using ListLoadBalancerPools
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		// return error causing client to retry
		w.WriteHeader(500)
		fmt.Fprint(w, `{
			"success": false,
			"errors": [ "server created some error"],
			"messages": [],
			"result": []
		}`)

	}

	mux.HandleFunc("/user/load_balancers/pools", handler)

	_, err := client.ListLoadBalancerPools(context.Background())
	assert.Error(t, err)
}

func TestZoneIDByNameWithNonUniqueZonesWithoutOrgID(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "023e105f4ecef8ad9ca31a8372d0c353",
					"name": "example.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@example.com",
						"owner_type": "user"
					},
					"account": {
						"id": "01a7362d577a6c3019a474fd6f485823",
						"name": "Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				},
				{
					"id": "023e105f4ecef8ad9ca31a8372d0c353",
					"name": "example.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@example.com",
						"owner_type": "user"
					},
					"account": {
						"id": "01a7362d577a6c3019a474fd6f485823",
						"name": "Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000,
				"total_pages": 1
			}
		}
		`)
	}

	// `HandleFunc` doesn't handle query parameters so we just need to
	// handle the `/zones` endpoint instead.
	mux.HandleFunc("/zones", handler)

	_, err := client.ZoneIDByName("example.com")
	assert.EqualError(t, err, "ambiguous zone name used without an account ID")
}

func TestZoneIDByNameWithNonUniqueZonesWithOrgId(t *testing.T) {
	setup(UsingAccount("01a7362d577a6c3019a474fd6f485823"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "023e105f4ecef8ad9ca31a8372d0c353",
					"name": "example.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@example.com",
						"owner_type": "user"
					},
					"account": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Another Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				},
				{
					"id": "7c5dae5552338874e5053f2534d2767a",
					"name": "example.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@example.com",
						"owner_type": "user"
					},
					"account": {
						"id": "01a7362d577a6c3019a474fd6f485823",
						"name": "Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000,
				"total_pages": 1
			}
		}
		`)
	}

	// `HandleFunc` doesn't handle query parameters so we just need to
	// handle the `/zones` endpoint instead.
	mux.HandleFunc("/zones", handler)

	actual, err := client.ZoneIDByName("example.com")
	if assert.NoError(t, err) {
		assert.Equal(t, actual, "7c5dae5552338874e5053f2534d2767a")
	}
}

func TestZoneIDByNameWithIDN(t *testing.T) {
	setup(UsingAccount("01a7362d577a6c3019a474fd6f485823"))
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"id": "023e105f4ecef8ad9ca31a8372d0c353",
					"name": "example.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@example.com",
						"owner_type": "user"
					},
					"account": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Another Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				},
				{
					"id": "7c5dae5552338874e5053f2534d2767a",
					"name": "exämple.com",
					"development_mode": 7200,
					"original_name_servers": [
						"ns1.originaldnshost.com",
						"ns2.originaldnshost.com"
					],
					"original_registrar": "GoDaddy",
					"original_dnshost": "NameCheap",
					"created_on": "2014-01-01T05:20:00.12345Z",
					"modified_on": "2014-01-01T05:20:00.12345Z",
					"owner": {
						"id": "7c5dae5552338874e5053f2534d2767a",
						"email": "user@exämple.com",
						"owner_type": "user"
					},
					"account": {
						"id": "01a7362d577a6c3019a474fd6f485823",
						"name": "Demo Account"
					},
					"permissions": [
						"#zone:read",
						"#zone:edit"
					],
					"plan": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"plan_pending": {
						"id": "e592fd9519420ba7405e1307bff33214",
						"name": "Pro Plan",
						"price": 20,
						"currency": "USD",
						"frequency": "monthly",
						"legacy_id": "pro",
						"is_subscribed": true,
						"can_subscribe": true
					},
					"status": "active",
					"paused": false,
					"type": "full",
					"name_servers": [
						"tony.ns.cloudflare.com",
						"woz.ns.cloudflare.com"
					]
				}
			],
			"result_info": {
				"page": 1,
				"per_page": 20,
				"count": 1,
				"total_count": 2000,
				"total_pages": 1
			}
		}
		`)
	}

	// `HandleFunc` doesn't handle query parameters so we just need to
	// handle the `/zones` endpoint instead.
	mux.HandleFunc("/zones", handler)

	actual, err := client.ZoneIDByName("exämple.com")
	if assert.NoError(t, err) {
		assert.Equal(t, actual, "7c5dae5552338874e5053f2534d2767a")
	}

	actual, err = client.ZoneIDByName("xn--exmple-cua.com")
	if assert.NoError(t, err) {
		assert.Equal(t, actual, "7c5dae5552338874e5053f2534d2767a")
	}

	_, err = client.ZoneIDByName("xn--exmple-corruted.com")
	assert.Error(t, err)
}

func TestClient_ContextIsPassedToRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	httpClient := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, ctx, r.Context())

			rec := httptest.NewRecorder()
			rec.WriteHeader(http.StatusOK)
			return rec.Result(), nil
		}),
	}

	cfClient, _ := New("deadbeef", "cloudflare@example.org", HTTPClient(httpClient))

	cfClient.ListZonesContext(ctx)
}

func TestErrorFromResponseWithUnmarshalingError(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, `{ not valid json`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/access/apps", handler)

	_, err := client.CreateAccessApplication(context.Background(), "01a7362d577a6c3019a474fd6f485823", AccessApplication{
		Name:            "Admin Site",
		Domain:          "test.example.com/admin",
		SessionDuration: "24h",
	})

	assert.Regexp(t, "error unmarshalling the JSON response error body: ", err)
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (t RoundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return t(request)
}
