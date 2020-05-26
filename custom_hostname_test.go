package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomHostname_DeleteCustomHostname(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `
{
  "id": "bar"
}`)
	})

	err := client.DeleteCustomHostname("foo", "bar")

	assert.NoError(t, err)
}

func TestCustomHostname_CreateCustomHostname(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
    "hostname": "app.example.com",
    "ssl": {
      "status": "pending_validation",
      "method": "cname",
      "type": "dv",
      "cname_target": "dcv.digicert.com",
      "cname": "810b7d5f01154524b961ba0cd578acc2.app.example.com",
      "custom_certificate": "-----BEGIN CERTIFICATE-----\\nMIIFJDCCBAygAwIBAgIQD0ifmj/Yi5NP/2gdUySbfzANBgkqhkiG9w0BAQsFADBN\\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMScwJQYDVQQDEx5E...SzSHfXp5lnu/3V08I72q1QNzOCgY1XeL4GKVcj4or6cT6tX6oJH7ePPmfrBfqI/O\\nOeH8gMJ+FuwtXYEPa4hBf38M5eU5xWG7\\n-----END CERTIFICATE-----\\n",
      "custom_key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAwQHoetcl9+5ikGzV6cMzWtWPJHqXT3wpbEkRU9Yz7lgvddmG\ndtcGbg/1CGZu0jJGkMoppoUo4c3dts3iwqRYmBikUP77wwY2QGmDZw2FvkJCJlKn\nabIRuGvBKwzESIXgKk2016aTP6/dAjEHyo6SeoK8lkIySUvK0fyOVlsiEsCmOpid\ntnKX/a+50GjB79CJH4ER2lLVZnhePFR/zUOyPxZQQ4naHf7yu/b5jhO0f8fwt+py\nFxIXjbEIdZliWRkRMtzrHOJIhrmJ2A1J7iOrirbbwillwjjNVUWPf3IJ3M12S9pE\newooaeO2izNTERcG9HzAacbVRn2Y2SWIyT/18QIDAQABAoIBACbhTYXBZYKmYPCb\nHBR1IBlCQA2nLGf0qRuJNJZg5iEzXows/6tc8YymZkQE7nolapWsQ+upk2y5Xdp/\naxiuprIs9JzkYK8Ox0r+dlwCG1kSW+UAbX0bQ/qUqlsTvU6muVuMP8vZYHxJ3wmb\n+ufRBKztPTQ/rYWaYQcgC0RWI20HTFBMxlTAyNxYNWzX7RKFkGVVyB9RsAtmcc8g\n+j4OdosbfNoJPS0HeIfNpAznDfHKdxDk2Yc1tV6RHBrC1ynyLE9+TaflIAdo2MVv\nKLMLq51GqYKtgJFIlBRPQqKoyXdz3fGvXrTkf/WY9QNq0J1Vk5ERePZ54mN8iZB7\n9lwy/AkCgYEA6FXzosxswaJ2wQLeoYc7ceaweX/SwTvxHgXzRyJIIT0eJWgx13Wo\n/WA3Iziimsjf6qE+SI/8laxPp2A86VMaIt3Z3mJN/CqSVGw8LK2AQst+OwdPyDMu\niacE8lj/IFGC8mwNUAb9CzGU3JpU4PxxGFjS/eMtGeRXCWkK4NE+G08CgYEA1Kp9\nN2JrVlqUz+gAX+LPmE9OEMAS9WQSQsfCHGogIFDGGcNf7+uwBM7GAaSJIP01zcoe\nVAgWdzXCv3FLhsaZoJ6RyLOLay5phbu1iaTr4UNYm5WtYTzMzqh8l1+MFFDl9xDB\nvULuCIIrglM5MeS/qnSg1uMoH2oVPj9TVst/ir8CgYEAxrI7Ws9Zc4Bt70N1As+U\nlySjaEVZCMkqvHJ6TCuVZFfQoE0r0whdLdRLU2PsLFP+q7qaeZQqgBaNSKeVcDYR\n9B+nY/jOmQoPewPVsp/vQTCnE/R81spu0mp0YI6cIheT1Z9zAy322svcc43JaWB7\nmEbeqyLOP4Z4qSOcmghZBSECgYACvR9Xs0DGn+wCsW4vze/2ei77MD4OQvepPIFX\ndFZtlBy5ADcgE9z0cuVB6CiL8DbdK5kwY9pGNr8HUCI03iHkW6Zs+0L0YmihfEVe\nPG19PSzK9CaDdhD9KFZSbLyVFmWfxOt50H7YRTTiPMgjyFpfi5j2q348yVT0tEQS\nfhRqaQKBgAcWPokmJ7EbYQGeMbS7HC8eWO/RyamlnSffdCdSc7ue3zdVJxpAkQ8W\nqu80pEIF6raIQfAf8MXiiZ7auFOSnHQTXUbhCpvDLKi0Mwq3G8Pl07l+2s6dQG6T\nlv6XTQaMyf6n1yjzL+fzDrH3qXMxHMO/b13EePXpDMpY7HQpoLDi\n-----END RSA PRIVATE KEY-----\n",
      "settings": {
        "http2": "on"
      }
    }
  }
}`)
	})

	response, err := client.CreateCustomHostname("foo", CustomHostname{Hostname: "app.example.com", SSL: CustomHostnameSSL{Method: "cname", Type: "dv"}})

	want := &CustomHostnameResponse{
		Result: CustomHostname{
			ID:       "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
			Hostname: "app.example.com",
			SSL: CustomHostnameSSL{
				Type:        "dv",
				Method:      "cname",
				Status:      "pending_validation",
				CnameTarget: "dcv.digicert.com",
				CnameName:   "810b7d5f01154524b961ba0cd578acc2.app.example.com",
				CustomCertificate: "-----BEGIN CERTIFICATE-----\\nMIIFJDCCBAygAwIBAgIQD0ifmj/Yi5NP/2gdUySbfzANBgkqhkiG9w0BAQsFADBN\\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMScwJQYDVQQDEx5E...SzSHfXp5lnu/3V08I72q1QNzOCgY1XeL4GKVcj4or6cT6tX6oJH7ePPmfrBfqI/O\\nOeH8gMJ+FuwtXYEPa4hBf38M5eU5xWG7\\n-----END CERTIFICATE-----\\n",
				CustomKey: "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAwQHoetcl9+5ikGzV6cMzWtWPJHqXT3wpbEkRU9Yz7lgvddmG\ndtcGbg/1CGZu0jJGkMoppoUo4c3dts3iwqRYmBikUP77wwY2QGmDZw2FvkJCJlKn\nabIRuGvBKwzESIXgKk2016aTP6/dAjEHyo6SeoK8lkIySUvK0fyOVlsiEsCmOpid\ntnKX/a+50GjB79CJH4ER2lLVZnhePFR/zUOyPxZQQ4naHf7yu/b5jhO0f8fwt+py\nFxIXjbEIdZliWRkRMtzrHOJIhrmJ2A1J7iOrirbbwillwjjNVUWPf3IJ3M12S9pE\newooaeO2izNTERcG9HzAacbVRn2Y2SWIyT/18QIDAQABAoIBACbhTYXBZYKmYPCb\nHBR1IBlCQA2nLGf0qRuJNJZg5iEzXows/6tc8YymZkQE7nolapWsQ+upk2y5Xdp/\naxiuprIs9JzkYK8Ox0r+dlwCG1kSW+UAbX0bQ/qUqlsTvU6muVuMP8vZYHxJ3wmb\n+ufRBKztPTQ/rYWaYQcgC0RWI20HTFBMxlTAyNxYNWzX7RKFkGVVyB9RsAtmcc8g\n+j4OdosbfNoJPS0HeIfNpAznDfHKdxDk2Yc1tV6RHBrC1ynyLE9+TaflIAdo2MVv\nKLMLq51GqYKtgJFIlBRPQqKoyXdz3fGvXrTkf/WY9QNq0J1Vk5ERePZ54mN8iZB7\n9lwy/AkCgYEA6FXzosxswaJ2wQLeoYc7ceaweX/SwTvxHgXzRyJIIT0eJWgx13Wo\n/WA3Iziimsjf6qE+SI/8laxPp2A86VMaIt3Z3mJN/CqSVGw8LK2AQst+OwdPyDMu\niacE8lj/IFGC8mwNUAb9CzGU3JpU4PxxGFjS/eMtGeRXCWkK4NE+G08CgYEA1Kp9\nN2JrVlqUz+gAX+LPmE9OEMAS9WQSQsfCHGogIFDGGcNf7+uwBM7GAaSJIP01zcoe\nVAgWdzXCv3FLhsaZoJ6RyLOLay5phbu1iaTr4UNYm5WtYTzMzqh8l1+MFFDl9xDB\nvULuCIIrglM5MeS/qnSg1uMoH2oVPj9TVst/ir8CgYEAxrI7Ws9Zc4Bt70N1As+U\nlySjaEVZCMkqvHJ6TCuVZFfQoE0r0whdLdRLU2PsLFP+q7qaeZQqgBaNSKeVcDYR\n9B+nY/jOmQoPewPVsp/vQTCnE/R81spu0mp0YI6cIheT1Z9zAy322svcc43JaWB7\nmEbeqyLOP4Z4qSOcmghZBSECgYACvR9Xs0DGn+wCsW4vze/2ei77MD4OQvepPIFX\ndFZtlBy5ADcgE9z0cuVB6CiL8DbdK5kwY9pGNr8HUCI03iHkW6Zs+0L0YmihfEVe\nPG19PSzK9CaDdhD9KFZSbLyVFmWfxOt50H7YRTTiPMgjyFpfi5j2q348yVT0tEQS\nfhRqaQKBgAcWPokmJ7EbYQGeMbS7HC8eWO/RyamlnSffdCdSc7ue3zdVJxpAkQ8W\nqu80pEIF6raIQfAf8MXiiZ7auFOSnHQTXUbhCpvDLKi0Mwq3G8Pl07l+2s6dQG6T\nlv6XTQaMyf6n1yjzL+fzDrH3qXMxHMO/b13EePXpDMpY7HQpoLDi\n-----END RSA PRIVATE KEY-----\n",
				Settings: CustomHostnameSSLSettings{
					HTTP2: "on",
				},
			},
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}

func TestCustomHostname_CreateCustomHostname_CustomOrigin(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
	"success": true,
	"errors": [],
	"messages": [],
	"result": {
		"id": "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
		"hostname": "app.example.com",
		"custom_origin_server": "example.app.com",
		"ssl": {
			"status": "pending_validation",
			"method": "cname",
			"type": "dv",
			"cname_target": "dcv.digicert.com",
			"cname": "810b7d5f01154524b961ba0cd578acc2.app.example.com",
			"settings": {
			"http2": "on"
			}
		}
  	}
}`)
	})

	response, err := client.CreateCustomHostname("foo", CustomHostname{Hostname: "app.example.com", CustomOriginServer: "example.app.com", SSL: CustomHostnameSSL{Method: "cname", Type: "dv"}})

	want := &CustomHostnameResponse{
		Result: CustomHostname{
			ID:                 "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
			Hostname:           "app.example.com",
			CustomOriginServer: "example.app.com",
			SSL: CustomHostnameSSL{
				Type:        "dv",
				Method:      "cname",
				Status:      "pending_validation",
				CnameTarget: "dcv.digicert.com",
				CnameName:   "810b7d5f01154524b961ba0cd578acc2.app.example.com",
				Settings: CustomHostnameSSLSettings{
					HTTP2: "on",
				},
			},
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}

func TestCustomHostname_CustomHostnames(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
	"success": true,
	"result": [
		{
			"id": "custom_host_1",
			"hostname": "custom.host.one",
			"ssl": {
				"type": "dv",
				"method": "cname",
				"status": "pending_validation",
				"cname_target": "dcv.digicert.com",
				"cname": "810b7d5f01154524b961ba0cd578acc2.app.example.com"
			},
			"custom_metadata": {
				"a_random_field": "random field value"
			},
			"status": "pending",
			"verification_errors": [
				"None of the A or AAAA records are owned by this account and the pre-generated ownership verification token was not found."
    		],
			"ownership_verification": {
				"type": "txt",
      			"name": "_cf-custom-hostname.app.example.com",
      			"value": "5cc07c04-ea62-4a5a-95f0-419334a875a4"
    		}
		}
	],
	"result_info": {
		"page": 1,
		"per_page": 20,
		"count": 5,
		"total_count": 5
	}
}`)
	})

	customHostnames, _, err := client.CustomHostnames("foo", 1, CustomHostname{})

	want := []CustomHostname{
		{
			ID:       "custom_host_1",
			Hostname: "custom.host.one",
			SSL: CustomHostnameSSL{
				Type:        "dv",
				Method:      "cname",
				Status:      "pending_validation",
				CnameTarget: "dcv.digicert.com",
				CnameName:   "810b7d5f01154524b961ba0cd578acc2.app.example.com",
			},
			CustomMetadata: CustomMetadata{"a_random_field": "random field value"},
			Status:         PENDING,
			VerificationErrors: []string{"None of the A or AAAA records are owned " +
				"by this account and the pre-generated ownership verification token was not found."},
			OwnershipVerification: CustomHostnameOwnershipVerification{
				Type:  "txt",
				Name:  "_cf-custom-hostname.app.example.com",
				Value: "5cc07c04-ea62-4a5a-95f0-419334a875a4",
			},
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, customHostnames)
	}
}

func TestCustomHostname_CustomHostname(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/bar", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
"success": true,
"result": {
    "id": "bar",
    "hostname": "foo.bar.com",
    "ssl": {
      "type": "dv",
      "method": "http",
      "status": "active",
      "settings": {
        "ciphers": ["ECDHE-RSA-AES128-GCM-SHA256","AES128-SHA"],
        "http2": "on",
        "min_tls_version": "1.2"
      }
    },
    "custom_metadata": {
      "origin": "a.custom.origin"
    },
	"status": "pending",
	"verification_errors": [
		"None of the A or AAAA records are owned by this account and the pre-generated ownership verification token was not found."
    ],
	"ownership_verification": {
		"type": "txt",
     	"name": "_cf-custom-hostname.app.example.com",
    	"value": "5cc07c04-ea62-4a5a-95f0-419334a875a4"
    }
  }
}`)
	})

	customHostname, err := client.CustomHostname("foo", "bar")

	want := CustomHostname{
		ID:       "bar",
		Hostname: "foo.bar.com",
		SSL: CustomHostnameSSL{
			Status: "active",
			Method: "http",
			Type:   "dv",
			Settings: CustomHostnameSSLSettings{
				HTTP2:         "on",
				MinTLSVersion: "1.2",
				Ciphers:       []string{"ECDHE-RSA-AES128-GCM-SHA256", "AES128-SHA"},
			},
		},
		CustomMetadata: CustomMetadata{"origin": "a.custom.origin"},
		Status:         PENDING,
		VerificationErrors: []string{"None of the A or AAAA records are owned " +
			"by this account and the pre-generated ownership verification token was not found."},
		OwnershipVerification: CustomHostnameOwnershipVerification{
			Type:  "txt",
			Name:  "_cf-custom-hostname.app.example.com",
			Value: "5cc07c04-ea62-4a5a-95f0-419334a875a4",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, customHostname)
	}
}

func TestCustomHostname_UpdateCustomHostnameSSL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/0d89c70d-ad9f-4843-b99f-6cc0252067e9", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
	"success": true,
	"errors": [],
	"messages": [],
	"result": {
		"id": "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
		"hostname": "app.example.com",
		"custom_origin_server": "example.app.com",
		"ssl": {
			"status": "pending_validation",
			"method": "cname",
			"type": "dv",
			"cname_target": "dcv.digicert.com",
			"cname": "810b7d5f01154524b961ba0cd578acc2.app.example.com",
			"settings": {
			"http2": "off",
			"tls_1_3": "on"
			}
		}
  	}
}`)
	})

	response, err := client.UpdateCustomHostnameSSL("foo", "0d89c70d-ad9f-4843-b99f-6cc0252067e9", CustomHostnameSSL{Method: "cname", Type: "dv", Settings: CustomHostnameSSLSettings{HTTP2: "off", TLS13: "on"}})

	want := &CustomHostnameResponse{
		Result: CustomHostname{
			ID:                 "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
			Hostname:           "app.example.com",
			CustomOriginServer: "example.app.com",
			SSL: CustomHostnameSSL{
				Type:        "dv",
				Method:      "cname",
				Status:      "pending_validation",
				CnameTarget: "dcv.digicert.com",
				CnameName:   "810b7d5f01154524b961ba0cd578acc2.app.example.com",
				Settings: CustomHostnameSSLSettings{
					HTTP2: "off",
					TLS13: "on",
				},
			},
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}

func TestCustomHostname_UpdateCustomHostname(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/0d89c70d-ad9f-4843-b99f-6cc0252067e9", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
	"success": true,
	"errors": [],
	"messages": [],
	"result": {
		"id": "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
		"hostname": "app.example.com",
		"custom_origin_server": "example.app.com",
		"ssl": {
			"status": "pending_validation",
			"method": "cname",
			"type": "dv",
			"cname_target": "dcv.digicert.com",
			"cname": "810b7d5f01154524b961ba0cd578acc2.app.example.com",
			"settings": {
			"http2": "off",
			"tls_1_3": "on"
			}
		}
  	}
}`)
	})

	response, err := client.UpdateCustomHostname("foo", "0d89c70d-ad9f-4843-b99f-6cc0252067e9", CustomHostname{Hostname: "app.example.com", CustomOriginServer: "example.app.com", SSL: CustomHostnameSSL{Method: "cname", Type: "dv"}})

	want := &CustomHostnameResponse{
		Result: CustomHostname{
			ID:                 "0d89c70d-ad9f-4843-b99f-6cc0252067e9",
			Hostname:           "app.example.com",
			CustomOriginServer: "example.app.com",
			SSL: CustomHostnameSSL{
				Type:        "dv",
				Method:      "cname",
				Status:      "pending_validation",
				CnameTarget: "dcv.digicert.com",
				CnameName:   "810b7d5f01154524b961ba0cd578acc2.app.example.com",
				Settings: CustomHostnameSSLSettings{
					HTTP2: "off",
					TLS13: "on",
				},
			},
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}

func TestCustomHostname_CustomHostnameFallbackOrigin(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/fallback_origin", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "origin": "fallback.example.com",
    "status": "pending_deployment",
    "errors": [
      "DNS records are not setup correctly. Origin should be a proxied A/AAAA/CNAME dns record"
    ],
    "created_at": "2019-10-28T18:11:23.37411Z",
    "updated_at": "2020-03-16T18:11:23.531995Z"
  }
}`)
	})

	customHostnameFallbackOrigin, err := client.CustomHostnameFallbackOrigin("foo")

	want := CustomHostnameFallbackOrigin{
		Origin: "fallback.example.com",
		Status: "pending_deployment",
		Errors: []string {"DNS records are not setup correctly. Origin should be a proxied A/AAAA/CNAME dns record", },
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, customHostnameFallbackOrigin)
	}
}

func TestCustomHostname_UpdateCustomHostnameFallbackOrigin(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/custom_hostnames/fallback_origin", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "origin": "fallback.example.com",
    "status": "pending_deployment",
    "errors": [
      "DNS records are not setup correctly. Origin should be a proxied A/AAAA/CNAME dns record"
    ],
    "created_at": "2019-10-28T18:11:23.37411Z",
    "updated_at": "2020-03-16T18:11:23.531995Z"
  }
}`)
	})

	response, err := client.UpdateCustomHostnameFallbackOrigin("foo", CustomHostnameFallbackOrigin{Origin: "fallback.example.com",})

	want := &CustomHostnameFallbackOriginResponse{
		Result: CustomHostnameFallbackOrigin{
			Origin: "fallback.example.com",
			Status: "pending_deployment",
			Errors: []string {"DNS records are not setup correctly. Origin should be a proxied A/AAAA/CNAME dns record", },
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}
