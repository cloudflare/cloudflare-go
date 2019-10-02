package cloudflare

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZoneAnalyticsDashboard(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "2015-01-01T12:23:00Z", r.URL.Query().Get("since"))
		assert.Equal(t, "2015-01-02T12:23:00Z", r.URL.Query().Get("until"))
		assert.Equal(t, "true", r.URL.Query().Get("continuous"))

		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-analytics-properties
		fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": {
            "totals": {
              "since": "2015-01-01T12:23:00Z",
              "until": "2015-01-02T12:23:00Z",
              "requests": {
                "all": 1234085328,
                "cached": 1234085328,
                "uncached": 13876154,
                "content_type": {
                  "css": 15343,
                  "html": 1234213,
                  "javascript": 318236,
                  "gif": 23178,
                  "jpeg": 1982048
                },
                "country": {
                  "US": 4181364,
                  "AG": 37298,
                  "GI": 293846
                },
                "ssl": {
                  "encrypted": 12978361,
                  "unencrypted": 781263
                },
                "http_status": {
                  "200": 13496983,
                  "301": 283,
                  "400": 187936,
                  "402": 1828,
                  "404": 1293
                }
              },
              "bandwidth": {
                "all": 213867451,
                "cached": 113205063,
                "uncached": 113205063,
                "content_type": {
                  "css": 237421,
                  "html": 1231290,
                  "javascript": 123245,
                  "gif": 1234242,
                  "jpeg": 784278
                },
                "country": {
                  "US": 123145433,
                  "AG": 2342483,
                  "GI": 984753
                },
                "ssl": {
                  "encrypted": 37592942,
                  "unencrypted": 237654192
                }
              },
              "threats": {
                "all": 23423873,
                "country": {
                  "US": 123,
                  "CN": 523423,
                  "AU": 91
                },
                "type": {
                  "user.ban.ip": 123,
                  "hot.ban.unknown": 5324,
                  "macro.chl.captchaErr": 1341,
                  "macro.chl.jschlErr": 5323
                }
              },
              "pageviews": {
                "all": 5724723,
                "search_engines": {
                  "googlebot": 35272,
                  "pingdom": 13435,
                  "bingbot": 5372,
                  "baidubot": 1345
                }
              },
              "uniques": {
                "all": 12343
              }
            },
            "timeseries": [
              {
                "since": "2015-01-01T12:23:00Z",
                "until": "2015-01-02T12:23:00Z",
                "requests": {
                  "all": 1234085328,
                  "cached": 1234085328,
                  "uncached": 13876154,
                  "content_type": {
                    "css": 15343,
                    "html": 1234213,
                    "javascript": 318236,
                    "gif": 23178,
                    "jpeg": 1982048
                  },
                  "country": {
                    "US": 4181364,
                    "AG": 37298,
                    "GI": 293846
                  },
                  "ssl": {
                    "encrypted": 12978361,
                    "unencrypted": 781263
                  },
                  "http_status": {
                    "200": 13496983,
                    "301": 283,
                    "400": 187936,
                    "402": 1828,
                    "404": 1293
                  }
                },
                "bandwidth": {
                  "all": 213867451,
                  "cached": 113205063,
                  "uncached": 113205063,
                  "content_type": {
                    "css": 237421,
                    "html": 1231290,
                    "javascript": 123245,
                    "gif": 1234242,
                    "jpeg": 784278
                  },
                  "country": {
                    "US": 123145433,
                    "AG": 2342483,
                    "GI": 984753
                  },
                  "ssl": {
                    "encrypted": 37592942,
                    "unencrypted": 237654192
                  }
                },
                "threats": {
                  "all": 23423873,
                  "country": {
                    "US": 123,
                    "CN": 523423,
                    "AU": 91
                  },
                  "type": {
                    "user.ban.ip": 123,
                    "hot.ban.unknown": 5324,
                    "macro.chl.captchaErr": 1341,
                    "macro.chl.jschlErr": 5323
                  }
                },
                "pageviews": {
                  "all": 5724723,
                  "search_engines": {
                    "googlebot": 35272,
                    "pingdom": 13435,
                    "bingbot": 5372,
                    "baidubot": 1345
                  }
                },
                "uniques": {
                  "all": 12343
                }
              }
            ]
          },
          "query": {
            "since": "2015-01-01T12:23:00Z",
            "until": "2015-01-02T12:23:00Z",
            "time_delta": 60
          }
        }`)
	}

	mux.HandleFunc("/zones/foo/analytics/dashboard", handler)

	since, _ := time.Parse(time.RFC3339, "2015-01-01T12:23:00Z")
	until, _ := time.Parse(time.RFC3339, "2015-01-02T12:23:00Z")
	data := ZoneAnalytics{
		Since: since,
		Until: until,
		Requests: struct {
			All         int            `json:"all"`
			Cached      int            `json:"cached"`
			Uncached    int            `json:"uncached"`
			ContentType map[string]int `json:"content_type"`
			Country     map[string]int `json:"country"`
			SSL         struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			} `json:"ssl"`
			HTTPStatus map[string]int `json:"http_status"`
		}{
			All:      1234085328,
			Cached:   1234085328,
			Uncached: 13876154,
			ContentType: map[string]int{
				"css":        15343,
				"html":       1234213,
				"javascript": 318236,
				"gif":        23178,
				"jpeg":       1982048,
			},
			Country: map[string]int{
				"US": 4181364,
				"AG": 37298,
				"GI": 293846,
			},
			SSL: struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			}{
				Encrypted:   12978361,
				Unencrypted: 781263,
			},
			HTTPStatus: map[string]int{
				"200": 13496983,
				"301": 283,
				"400": 187936,
				"402": 1828,
				"404": 1293,
			},
		},
		Bandwidth: struct {
			All         int            `json:"all"`
			Cached      int            `json:"cached"`
			Uncached    int            `json:"uncached"`
			ContentType map[string]int `json:"content_type"`
			Country     map[string]int `json:"country"`
			SSL         struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			} `json:"ssl"`
		}{
			All:      213867451,
			Cached:   113205063,
			Uncached: 113205063,
			ContentType: map[string]int{
				"css":        237421,
				"html":       1231290,
				"javascript": 123245,
				"gif":        1234242,
				"jpeg":       784278,
			},
			Country: map[string]int{
				"US": 123145433,
				"AG": 2342483,
				"GI": 984753,
			},
			SSL: struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			}{
				Encrypted:   37592942,
				Unencrypted: 237654192,
			},
		},
		Threats: struct {
			All     int            `json:"all"`
			Country map[string]int `json:"country"`
			Type    map[string]int `json:"type"`
		}{
			All: 23423873,
			Country: map[string]int{
				"US": 123,
				"CN": 523423,
				"AU": 91,
			},
			Type: map[string]int{
				"user.ban.ip":          123,
				"hot.ban.unknown":      5324,
				"macro.chl.captchaErr": 1341,
				"macro.chl.jschlErr":   5323,
			},
		},
		Pageviews: struct {
			All           int            `json:"all"`
			SearchEngines map[string]int `json:"search_engines"`
		}{
			All: 5724723,
			SearchEngines: map[string]int{
				"googlebot": 35272,
				"pingdom":   13435,
				"bingbot":   5372,
				"baidubot":  1345,
			},
		},
		Uniques: struct {
			All int `json:"all"`
		}{
			All: 12343,
		},
	}
	want := ZoneAnalyticsData{
		Totals:     data,
		Timeseries: []ZoneAnalytics{data},
	}

	continuous := true
	d, err := client.ZoneAnalyticsDashboard("foo", ZoneAnalyticsOptions{
		Since:      &since,
		Until:      &until,
		Continuous: &continuous,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, d)
	}

	_, err = client.ZoneAnalyticsDashboard("bar", ZoneAnalyticsOptions{})
	assert.Error(t, err)
}

func TestZoneAnalyticsByColocation(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "2015-01-01T12:23:00Z", r.URL.Query().Get("since"))
		assert.Equal(t, "2015-01-02T12:23:00Z", r.URL.Query().Get("until"))
		assert.Equal(t, "true", r.URL.Query().Get("continuous"))

		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-analytics-analytics-by-co-locations
		fmt.Fprintf(w, `{
          "success": true,
          "errors": [],
          "messages": [],
          "result": [
            {
              "colo_id": "SFO",
              "timeseries": [
                {
                  "since": "2015-01-01T12:23:00Z",
                  "until": "2015-01-02T12:23:00Z",
                  "requests": {
                    "all": 1234085328,
                    "cached": 1234085328,
                    "uncached": 13876154,
                    "content_type": {
                      "css": 15343,
                      "html": 1234213,
                      "javascript": 318236,
                      "gif": 23178,
                      "jpeg": 1982048
                    },
                    "country": {
                      "US": 4181364,
                      "AG": 37298,
                      "GI": 293846
                    },
                    "ssl": {
                      "encrypted": 12978361,
                      "unencrypted": 781263
                    },
                    "http_status": {
                      "200": 13496983,
                      "301": 283,
                      "400": 187936,
                      "402": 1828,
                      "404": 1293
                    }
                  },
                  "bandwidth": {
                    "all": 213867451,
                    "cached": 113205063,
                    "uncached": 113205063,
                    "content_type": {
                      "css": 237421,
                      "html": 1231290,
                      "javascript": 123245,
                      "gif": 1234242,
                      "jpeg": 784278
                    },
                    "country": {
                      "US": 123145433,
                      "AG": 2342483,
                      "GI": 984753
                    },
                    "ssl": {
                      "encrypted": 37592942,
                      "unencrypted": 237654192
                    }
                  },
                  "threats": {
                    "all": 23423873,
                    "country": {
                      "US": 123,
                      "CN": 523423,
                      "AU": 91
                    },
                    "type": {
                      "user.ban.ip": 123,
                      "hot.ban.unknown": 5324,
                      "macro.chl.captchaErr": 1341,
                      "macro.chl.jschlErr": 5323
                    }
                  },
                  "pageviews": {
                    "all": 5724723,
                    "search_engines": {
                      "googlebot": 35272,
                      "pingdom": 13435,
                      "bingbot": 5372,
                      "baidubot": 1345
                    }
                  },
                  "uniques": {
                    "all": 12343
                  }
                }
              ]
            }
          ],
          "query": {
            "since": "2015-01-01T12:23:00Z",
            "until": "2015-01-02T12:23:00Z",
            "time_delta": 60
          }
        }`)
	}

	mux.HandleFunc("/zones/foo/analytics/colos", handler)

	since, _ := time.Parse(time.RFC3339, "2015-01-01T12:23:00Z")
	until, _ := time.Parse(time.RFC3339, "2015-01-02T12:23:00Z")
	data := ZoneAnalytics{
		Since: since,
		Until: until,
		Requests: struct {
			All         int            `json:"all"`
			Cached      int            `json:"cached"`
			Uncached    int            `json:"uncached"`
			ContentType map[string]int `json:"content_type"`
			Country     map[string]int `json:"country"`
			SSL         struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			} `json:"ssl"`
			HTTPStatus map[string]int `json:"http_status"`
		}{
			All:      1234085328,
			Cached:   1234085328,
			Uncached: 13876154,
			ContentType: map[string]int{
				"css":        15343,
				"html":       1234213,
				"javascript": 318236,
				"gif":        23178,
				"jpeg":       1982048,
			},
			Country: map[string]int{
				"US": 4181364,
				"AG": 37298,
				"GI": 293846,
			},
			SSL: struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			}{
				Encrypted:   12978361,
				Unencrypted: 781263,
			},
			HTTPStatus: map[string]int{
				"200": 13496983,
				"301": 283,
				"400": 187936,
				"402": 1828,
				"404": 1293,
			},
		},
		Bandwidth: struct {
			All         int            `json:"all"`
			Cached      int            `json:"cached"`
			Uncached    int            `json:"uncached"`
			ContentType map[string]int `json:"content_type"`
			Country     map[string]int `json:"country"`
			SSL         struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			} `json:"ssl"`
		}{
			All:      213867451,
			Cached:   113205063,
			Uncached: 113205063,
			ContentType: map[string]int{
				"css":        237421,
				"html":       1231290,
				"javascript": 123245,
				"gif":        1234242,
				"jpeg":       784278,
			},
			Country: map[string]int{
				"US": 123145433,
				"AG": 2342483,
				"GI": 984753,
			},
			SSL: struct {
				Encrypted   int `json:"encrypted"`
				Unencrypted int `json:"unencrypted"`
			}{
				Encrypted:   37592942,
				Unencrypted: 237654192,
			},
		},
		Threats: struct {
			All     int            `json:"all"`
			Country map[string]int `json:"country"`
			Type    map[string]int `json:"type"`
		}{
			All: 23423873,
			Country: map[string]int{
				"US": 123,
				"CN": 523423,
				"AU": 91,
			},
			Type: map[string]int{
				"user.ban.ip":          123,
				"hot.ban.unknown":      5324,
				"macro.chl.captchaErr": 1341,
				"macro.chl.jschlErr":   5323,
			},
		},
		Pageviews: struct {
			All           int            `json:"all"`
			SearchEngines map[string]int `json:"search_engines"`
		}{
			All: 5724723,
			SearchEngines: map[string]int{
				"googlebot": 35272,
				"pingdom":   13435,
				"bingbot":   5372,
				"baidubot":  1345,
			},
		},
		Uniques: struct {
			All int `json:"all"`
		}{
			All: 12343,
		},
	}
	want := []ZoneAnalyticsColocation{
		{
			ColocationID: "SFO",
			Timeseries:   []ZoneAnalytics{data},
		},
	}

	continuous := true
	d, err := client.ZoneAnalyticsByColocation("foo", ZoneAnalyticsOptions{
		Since:      &since,
		Until:      &until,
		Continuous: &continuous,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, d)
	}

	_, err = client.ZoneAnalyticsDashboard("bar", ZoneAnalyticsOptions{})
	assert.Error(t, err)
}

func TestWithPagination(t *testing.T) {
	opt := reqOption{
		params: url.Values{},
	}
	popts := PaginationOptions{
		Page:    45,
		PerPage: 500,
	}
	of := WithPagination(popts)
	of(&opt)

	tests := []struct {
		name     string
		expected string
	}{
		{"page", "45"},
		{"per_page", "500"},
	}

	for _, tt := range tests {
		if got := opt.params.Get(tt.name); got != tt.expected {
			t.Errorf("expected param %s to be %s, got %s", tt.name, tt.expected, got)
		}
	}
}

func TestZoneFilter(t *testing.T) {
	opt := reqOption{
		params: url.Values{},
	}
	of := WithZoneFilter("example.org")
	of(&opt)

	if got := opt.params.Get("name"); got != "example.org" {
		t.Errorf("expected param %s to be %s, got %s", "name", "example.org", got)
	}
}

var createdAndModifiedOn, _ = time.Parse(time.RFC3339, "2014-01-01T05:20:00.12345Z")
var expectedFullZoneSetup = Zone{
	ID:      "023e105f4ecef8ad9ca31a8372d0c353",
	Name:    "example.com",
	DevMode: 7200,
	OriginalNS: []string{
		"ns1.originaldnshost.com",
		"ns2.originaldnshost.com",
	},
	OriginalRegistrar: "GoDaddy",
	OriginalDNSHost:   "NameCheap",
	CreatedOn:         createdAndModifiedOn,
	ModifiedOn:        createdAndModifiedOn,
	Owner: Owner{
		ID:        "7c5dae5552338874e5053f2534d2767a",
		Email:     "user@example.com",
		OwnerType: "user",
	},
	Account: Account{
		ID:   "01a7362d577a6c3019a474fd6f485823",
		Name: "Demo Account",
	},
	Permissions: []string{"#zone:read", "#zone:edit"},
	Plan: ZonePlan{
		ZonePlanCommon: ZonePlanCommon{
			ID:        "e592fd9519420ba7405e1307bff33214",
			Name:      "Pro Plan",
			Price:     20,
			Currency:  "USD",
			Frequency: "monthly",
		},
		LegacyID:     "pro",
		IsSubscribed: true,
		CanSubscribe: true,
	},
	PlanPending: ZonePlan{
		ZonePlanCommon: ZonePlanCommon{
			ID:        "e592fd9519420ba7405e1307bff33214",
			Name:      "Pro Plan",
			Price:     20,
			Currency:  "USD",
			Frequency: "monthly",
		},
		LegacyID:     "pro",
		IsSubscribed: true,
		CanSubscribe: true,
	},
	Status:      "active",
	Paused:      false,
	Type:        "full",
	NameServers: []string{"tony.ns.cloudflare.com", "woz.ns.cloudflare.com"},
}
var expectedPartialZoneSetup = Zone{
	ID:      "023e105f4ecef8ad9ca31a8372d0c353",
	Name:    "example.com",
	DevMode: 7200,
	OriginalNS: []string{
		"ns1.originaldnshost.com",
		"ns2.originaldnshost.com",
	},
	OriginalRegistrar: "GoDaddy",
	OriginalDNSHost:   "NameCheap",
	CreatedOn:         createdAndModifiedOn,
	ModifiedOn:        createdAndModifiedOn,
	Owner: Owner{
		ID:        "7c5dae5552338874e5053f2534d2767a",
		Email:     "user@example.com",
		OwnerType: "user",
	},
	Account: Account{
		ID:   "01a7362d577a6c3019a474fd6f485823",
		Name: "Demo Account",
	},
	Permissions: []string{"#zone:read", "#zone:edit"},
	Plan: ZonePlan{
		ZonePlanCommon: ZonePlanCommon{
			ID:        "e592fd9519420ba7405e1307bff33214",
			Name:      "Pro Plan",
			Price:     20,
			Currency:  "USD",
			Frequency: "monthly",
		},
		LegacyID:     "pro",
		IsSubscribed: true,
		CanSubscribe: true,
	},
	PlanPending: ZonePlan{
		ZonePlanCommon: ZonePlanCommon{
			ID:        "e592fd9519420ba7405e1307bff33214",
			Name:      "Pro Plan",
			Price:     20,
			Currency:  "USD",
			Frequency: "monthly",
		},
		LegacyID:     "pro",
		IsSubscribed: true,
		CanSubscribe: true,
	},
	Status:      "active",
	Paused:      false,
	Type:        "partial",
	NameServers: []string{"tony.ns.cloudflare.com", "woz.ns.cloudflare.com"},
}

func TestCreateZoneFullSetup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
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
				"activated_on": "2014-01-02T00:01:00.12345Z",
				"owner": {
					"id": "7c5dae5552338874e5053f2534d2767a",
					"email": "user@example.com",
					"type": "user"
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
		}
		`)
	}

	mux.HandleFunc("/zones", handler)

	actual, err := client.CreateZone(
		"example.com",
		false,
		Account{ID: "01a7362d577a6c3019a474fd6f485823"},
		"full",
	)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedFullZoneSetup, actual)
	}
}

func TestCreateZonePartialSetup(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
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
				"activated_on": "2014-01-02T00:01:00.12345Z",
				"owner": {
					"id": "7c5dae5552338874e5053f2534d2767a",
					"email": "user@example.com",
					"type": "user"
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
				"type": "partial",
				"name_servers": [
					"tony.ns.cloudflare.com",
					"woz.ns.cloudflare.com"
				]
			}
		}
		`)
	}

	mux.HandleFunc("/zones", handler)

	actual, err := client.CreateZone(
		"example.com",
		false,
		Account{ID: "01a7362d577a6c3019a474fd6f485823"},
		"partial",
	)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedPartialZoneSetup, actual)
	}
}

func TestFallbackOrigin_FallbackOrigin(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/fallback_origin", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
"success": true,
"errors": [],
"messages": [],
"result": {
    "id": "fallback_origin",
    "value": "app.example.com",
    "editable": true
  }
}`)
	})

	fallbackOrigin, err := client.FallbackOrigin("foo")

	want := FallbackOrigin{
		ID:    "fallback_origin",
		Value: "app.example.com",
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, fallbackOrigin)
	}
}

func TestFallbackOrigin_UpdateFallbackOrigin(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/foo/fallback_origin", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "fallback_origin",
    "value": "app.example.com",
		"editable": true
  }
}`)
	})

	response, err := client.UpdateFallbackOrigin("foo", FallbackOrigin{Value: "app.example.com"})

	want := &FallbackOriginResponse{
		Result: FallbackOrigin{
			ID:    "fallback_origin",
			Value: "app.example.com",
		},
		Response: Response{Success: true, Errors: []ResponseInfo{}, Messages: []ResponseInfo{}},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, response)
	}
}

func Test_normalizeZoneName(t *testing.T) {
	tests := []struct {
		name     string
		zone     string
		expected string
	}{
		{
			name:     "unicode stays unicode",
			zone:     "ünì¢øðe.tld",
			expected: "ünì¢øðe.tld",
		}, {
			name:     "valid punycode is normalized to unicode",
			zone:     "xn--ne-7ca90ava1cya.tld",
			expected: "ünì¢øðe.tld",
		}, {
			name:     "valid punycode in second label",
			zone:     "example.xn--j6w193g",
			expected: "example.香港",
		}, {
			name:     "invalid punycode is returned without change",
			zone:     "xn-invalid.xn-invalid-tld",
			expected: "xn-invalid.xn-invalid-tld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := normalizeZoneName(tt.zone)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestZonePartialHasVerificationKey(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		// JSON data from: https://api.cloudflare.com/#zone-zone-details (plus an undocumented field verification_key from curl to API)
		fmt.Fprintf(w, `{
  "result": {
    "id": "foo",
    "name": "bar",
    "status": "active",
    "paused": false,
    "type": "partial",
    "development_mode": 0,
    "verification_key": "foo-bar",
    "original_name_servers": ["a","b","c","d"],
    "original_registrar": null,
    "original_dnshost": null,
    "modified_on": "2019-09-04T15:11:43.409805Z",
    "created_on": "2018-12-06T14:33:38.410126Z",
    "activated_on": "2018-12-06T14:34:39.274528Z",
    "meta": {
      "step": 4,
      "wildcard_proxiable": true,
      "custom_certificate_quota": 1,
      "page_rule_quota": 100,
      "phishing_detected": false,
      "multiple_railguns_allowed": false
    },
    "owner": {
      "id": "bbbbbbbbbbbbbbbbbbbbbbbb",
      "type": "organization",
      "name": "OrgName"
    },
    "account": {
      "id": "aaaaaaaaaaaaaaaaaaaaaaaa",
      "name": "AccountName"
    },
    "permissions": [
      "#access:edit",
      "#access:read",
      "#analytics:read",
      "#app:edit",
      "#auditlogs:read",
      "#billing:read",
      "#cache_purge:edit",
      "#dns_records:edit",
      "#dns_records:read",
      "#lb:edit",
      "#lb:read",
      "#legal:read",
      "#logs:edit",
      "#logs:read",
      "#member:read",
      "#organization:edit",
      "#organization:read",
      "#ssl:edit",
      "#ssl:read",
      "#stream:edit",
      "#stream:read",
      "#subscription:edit",
      "#subscription:read",
      "#waf:edit",
      "#waf:read",
      "#webhooks:edit",
      "#webhooks:read",
      "#worker:edit",
      "#worker:read",
      "#zone:edit",
      "#zone:read",
      "#zone_settings:edit",
      "#zone_settings:read"
    ],
    "plan": {
      "id": "94f3b7b768b0458b56d2cac4fe5ec0f9",
      "name": "Enterprise Website",
      "price": 0,
      "currency": "USD",
      "frequency": "monthly",
      "is_subscribed": true,
      "can_subscribe": true,
      "legacy_id": "enterprise",
      "legacy_discount": false,
      "externally_managed": true
    }
  },
  "success": true,
  "errors": [],
  "messages": []
}`)
	}

	mux.HandleFunc("/zones/foo", handler)

	z, err := client.ZoneDetails("foo")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, z.VerificationKey)
		assert.Equal(t, z.VerificationKey,"foo-bar")
	}
}
