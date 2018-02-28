package cloudflare

import (
	"fmt"
	"net/http"
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
                },
                "ip_class": {
                  "badHost": 675,
                  "monitoringService": 7,
                  "noRecord": 4946167,
                  "searchEngine": 33098,
                  "tor": 10,
                  "whitelist": 109418
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
                "search_engine": {
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
                  },
                  "ip_class": {
                    "badHost": 675,
                    "monitoringService": 7,
                    "noRecord": 4946167,
                    "searchEngine": 33098,
                    "tor": 10,
                    "whitelist": 109418
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
                  "search_engine": {
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
			IPClass    map[string]int `json:"ip_class"`
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
			IPClass: map[string]int{
				"badHost":           675,
				"monitoringService": 7,
				"noRecord":          4946167,
				"searchEngine":      33098,
				"tor":               10,
				"whitelist":         109418,
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
			SearchEngines map[string]int `json:"search_engine"`
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
                    },
                    "ip_class": {
                      "badHost": 675,
                      "monitoringService": 7,
                      "noRecord": 4946167,
                      "searchEngine": 33098,
                      "tor": 10,
                      "whitelist": 109418
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
                    "search_engine": {
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
			IPClass    map[string]int `json:"ip_class"`
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
			IPClass: map[string]int{
				"badHost":           675,
				"monitoringService": 7,
				"noRecord":          4946167,
				"searchEngine":      33098,
				"tor":               10,
				"whitelist":         109418,
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
			SearchEngines map[string]int `json:"search_engine"`
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
	want := []ZoneAnalyticsData{
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

func TestZoneDNSAnalytics(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "2018-02-26T02:53:07Z", r.URL.Query().Get("since"))
		assert.Equal(t, "2018-02-26T02:54:07Z", r.URL.Query().Get("until"))
		assert.Equal(t, "queryCount,uncachedCount,staleCount,responseTimeAvg", r.URL.Query().Get("metrics"))
		assert.Equal(t, "queryName,responseCode,origin,tcp,ipVersion,coloName,queryType", r.URL.Query().Get("dimensions"))
		assert.Equal(t, "", r.URL.Query().Get("filters"))
		assert.Equal(t, "", r.URL.Query().Get("sort"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
        "result": {
          "rows": 52,
          "data": [
            {
              "dimensions": [
                "mydomain.com",
                "NOERROR",
                "::",
                "0",
                "4",
                "ATL",
                "AAAA"
              ],
              "metrics": [
                6,
                6,
                0,
                1.0
              ]
            },
            {
              "dimensions": [
                "mydomain.com",
                "NOERROR",
                "::",
                "0",
                "4",
                "LAX",
                "A"
              ],
              "metrics": [
                2,
                2,
                0,
                1.0
              ]
            }
          ],
          "data_lag": 60,
          "min": {
            "queryCount": 1,
            "staleCount": 0,
            "uncachedCount": 1,
            "responseTimeAvg": 1.0
          },
          "max": {
            "queryCount": 15,
            "staleCount": 0,
            "uncachedCount": 15,
            "responseTimeAvg": 1.0
          },
          "totals": {
            "queryCount": 147,
            "staleCount": 0,
            "uncachedCount": 147,
            "responseTimeAvg": 2.0
          },
          "query": {
            "dimensions": [
              "queryName",
              "responseCode",
              "origin",
              "tcp",
              "ipVersion",
              "coloName",
              "queryType"
            ],
            "metrics": [
              "queryCount",
              "uncachedCount",
              "staleCount",
              "responseTimeAvg"
            ],
            "since": "2018-02-26T02:53:07Z",
            "until": "2018-02-26T02:54:07Z",
            "limit": 10000
          }
        },
        "success": true,
        "errors": [],
        "messages": []
      }`)
	}

	mux.HandleFunc("/zones/foo/dns_analytics/report", handler)

	since, _ := time.Parse(time.RFC3339, "2018-02-26T02:53:07Z")
	until, _ := time.Parse(time.RFC3339, "2018-02-26T02:54:07Z")
	limit := 10000
	rows := []ZoneDNSAnalyticsRow{
		ZoneDNSAnalyticsRow{
			Dimensions: []string{"mydomain.com", "NOERROR", "::", "0", "4", "ATL", "AAAA"},
			Metrics:    []float64{6, 6, 0, 1.0},
		},
		ZoneDNSAnalyticsRow{
			Dimensions: []string{"mydomain.com", "NOERROR", "::", "0", "4", "LAX", "A"},
			Metrics:    []float64{2, 2, 0, 1.0},
		},
	}
	want := ZoneDNSAnalyticsData{
		Rows:    rows,
		DataLag: 60,
		Max: ZoneDNSAnalyticsDataContainer{
			QueryCount:      15,
			StaleCount:      0,
			UncachedCount:   15,
			ResponseTimeAvg: 1.0,
		},
		Min: ZoneDNSAnalyticsDataContainer{
			QueryCount:      1,
			StaleCount:      0,
			UncachedCount:   1,
			ResponseTimeAvg: 1.0,
		},
		Query: ZoneDNSAnalyticsOptions{
			Dimensions: []string{"queryName", "responseCode", "origin", "tcp", "ipVersion", "coloName", "queryType"},
			Metrics:    []string{"queryCount", "uncachedCount", "staleCount", "responseTimeAvg"},
			Since:      &since,
			Until:      &until,
			Limit:      &limit,
		},
		RowCount: 52,
		Totals: ZoneDNSAnalyticsDataContainer{
			QueryCount:      147,
			StaleCount:      0,
			UncachedCount:   147,
			ResponseTimeAvg: 2.0,
		},
	}

	d, err := client.ZoneDNSAnalytics("foo", ZoneDNSAnalyticsOptions{
		Since:      &since,
		Until:      &until,
		Metrics:    []string{"queryCount", "uncachedCount", "staleCount", "responseTimeAvg"},
		Dimensions: []string{"queryName", "responseCode", "origin", "tcp", "ipVersion", "coloName", "queryType"},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, d)
	}

	_, err = client.ZoneDNSAnalytics("bar", ZoneDNSAnalyticsOptions{})
	assert.Error(t, err)
}

func TestZoneDNSAnalyticsByTime(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "2018-02-28T01:00:58Z", r.URL.Query().Get("since"))
		assert.Equal(t, "2018-02-28T01:50:00Z", r.URL.Query().Get("until"))
		assert.Equal(t, "queryCount,uncachedCount,staleCount,responseTimeAvg", r.URL.Query().Get("metrics"))
		assert.Equal(t, "queryName,responseCode,origin,tcp,ipVersion,coloName,queryType", r.URL.Query().Get("dimensions"))
		assert.Equal(t, "dekaminute", r.URL.Query().Get("time_delta"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
      "result": {
        "rows": 1853,
        "data": [
          {
            "dimensions": [
              "www.mydomain.com",
              "NOERROR",
              "::",
              "0",
              "6",
              "DFW",
              "AAAA"
            ],
            "metrics": [
              [
                3,
                0,
                1,
                1,
                0,
                0
              ],
              [
                3,
                0,
                1,
                1,
                0,
                0
              ],
              [
                0,
                0,
                0,
                0,
                0,
                0
              ]
            ]
          },
          {
            "dimensions": [
              "www.mydomain.com",
              "NOERROR",
              "::",
              "0",
              "4",
              "AMS",
              "AAAA"
            ],
            "metrics": [
              [
                10,
                3,
                4,
                4,
                4,
                0
              ],
              [
                10,
                3,
                4,
                4,
                4,
                0
              ],
              [
                0,
                0,
                0,
                0,
                0,
                0
              ]
            ]
          }
        ],
        "data_lag": 60,
        "min": {
          "queryCount": 1,
          "staleCount": 0,
          "uncachedCount": 1
        },
        "max": {
          "queryCount": 299,
          "staleCount": 0,
          "uncachedCount": 299
        },
        "totals": {
          "queryCount": 15703,
          "staleCount": 0,
          "uncachedCount": 15703
        },
        "time_intervals": [
          [
            "2018-02-28T01:00:00Z",
            "2018-02-28T01:09:59Z"
          ],
          [
            "2018-02-28T01:10:00Z",
            "2018-02-28T01:19:59Z"
          ],
          [
            "2018-02-28T01:20:00Z",
            "2018-02-28T01:29:59Z"
          ],
          [
            "2018-02-28T01:30:00Z",
            "2018-02-28T01:39:59Z"
          ],
          [
            "2018-02-28T01:40:00Z",
            "2018-02-28T01:49:59Z"
          ],
          [
            "2018-02-28T01:50:00Z",
            "2018-02-28T01:50:00Z"
          ]
        ],
        "query": {
          "dimensions": [
            "queryName",
            "responseCode",
            "origin",
            "tcp",
            "ipVersion",
            "coloName",
            "queryType"
          ],
          "metrics": [
            "queryCount",
            "uncachedCount",
            "staleCount"
          ],
          "since": "2018-02-28T01:00:58Z",
          "until": "2018-02-28T01:50:00Z",
          "time_delta": "dekaminute",
          "limit": 10000
        }
      },
      "success": true,
      "errors": [],
      "messages": []
    }`)
	}

	mux.HandleFunc("/zones/foo/dns_analytics/report/bytime", handler)

	since, _ := time.Parse(time.RFC3339, "2018-02-28T01:00:58Z")
	until, _ := time.Parse(time.RFC3339, "2018-02-28T01:50:00Z")
	time1Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:00:00Z")
	time1End, _ := time.Parse(time.RFC3339, "2018-02-28T01:09:59Z")
	time1 := []time.Time{time1Start, time1End}
	time2Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:10:00Z")
	time2End, _ := time.Parse(time.RFC3339, "2018-02-28T01:19:59Z")
	time2 := []time.Time{time2Start, time2End}
	time3Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:20:00Z")
	time3End, _ := time.Parse(time.RFC3339, "2018-02-28T01:29:59Z")
	time3 := []time.Time{time3Start, time3End}
	time4Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:30:00Z")
	time4End, _ := time.Parse(time.RFC3339, "2018-02-28T01:39:59Z")
	time4 := []time.Time{time4Start, time4End}
	time5Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:40:00Z")
	time5End, _ := time.Parse(time.RFC3339, "2018-02-28T01:49:59Z")
	time5 := []time.Time{time5Start, time5End}
	time6Start, _ := time.Parse(time.RFC3339, "2018-02-28T01:50:00Z")
	time6End, _ := time.Parse(time.RFC3339, "2018-02-28T01:50:00Z")
	time6 := []time.Time{time6Start, time6End}
	limit := 10000
	timeDelta := "dekaminute"
	rows := []ZoneDNSAnalyticsByTimeRow{
		{
			Dimensions: []string{"www.mydomain.com", "NOERROR", "::", "0", "6", "DFW", "AAAA"},
			Metrics:    [][]float64{{3, 0, 1, 1, 0, 0}, {3, 0, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0}},
		},
		{
			Dimensions: []string{"www.mydomain.com", "NOERROR", "::", "0", "4", "AMS", "AAAA"},
			Metrics:    [][]float64{{10, 3, 4, 4, 4, 0}, {10, 3, 4, 4, 4, 0}, {0, 0, 0, 0, 0, 0}},
		},
	}
	want := ZoneDNSAnalyticsByTimeData{
		Rows:    rows,
		DataLag: 60,
		Max: ZoneDNSAnalyticsDataContainer{
			QueryCount:    299,
			StaleCount:    0,
			UncachedCount: 299,
		},
		Min: ZoneDNSAnalyticsDataContainer{
			QueryCount:    1,
			StaleCount:    0,
			UncachedCount: 1,
		},
		Query: ZoneDNSAnalyticsOptions{
			Dimensions: []string{"queryName", "responseCode", "origin", "tcp", "ipVersion", "coloName", "queryType"},
			Metrics:    []string{"queryCount", "uncachedCount", "staleCount"},
			Since:      &since,
			Until:      &until,
			Limit:      &limit,
			TimeDelta:  &timeDelta,
		},
		RowCount: 1853,
		Totals: ZoneDNSAnalyticsDataContainer{
			QueryCount:    15703,
			StaleCount:    0,
			UncachedCount: 15703,
		},
		TimeIntervals: [][]time.Time{time1, time2, time3, time4, time5, time6},
	}

	d, err := client.ZoneDNSAnalyticsByTime("foo", ZoneDNSAnalyticsOptions{
		Since:      &since,
		Until:      &until,
		Metrics:    []string{"queryCount", "uncachedCount", "staleCount", "responseTimeAvg"},
		Dimensions: []string{"queryName", "responseCode", "origin", "tcp", "ipVersion", "coloName", "queryType"},
		TimeDelta:  &timeDelta,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, d)
	}

	_, err = client.ZoneDNSAnalyticsByTime("bar", ZoneDNSAnalyticsOptions{})
	assert.Error(t, err)
}
